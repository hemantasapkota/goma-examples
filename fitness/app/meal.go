package app

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/hemantasapkota/goma"
)

type MealList []MealItem

type MealItem struct {
	Timestamp   string  `json:"timestamp"`
	Description string  `json:"description"`
	Calories    float64 `json:"calories"`
	Macros      struct {
		Protein float64 `json:"protein"`
		Fat     float64 `json:"fat"`
		Carbs   float64 `carbs:"carbs"`
	} `json:"macros"`

	// View releated stuffs
	GroupName string `json:"group"`
}

// Sort interface

func (m MealList) Len() int {
	return len(m)
}

func (m MealList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MealList) Less(i, j int) bool {
	return m[i].Timestamp > m[j].Timestamp
}

// End sort

func (m MealList) PrettyPrint() {
	for _, item := range m {
		group := item.GroupName
		if len(group) == 0 {
			group = "\t"
		}

		goma.Log(fmt.Sprintf("%s %s %f", group, item.Description, item.Calories))
	}
}

type Meal struct {
	*goma.Object
	Children MealList `json:"children"`

	// View related stuffs
	Groups map[string][]MealItem
}

func (m Meal) Key() string {
	return "fitnessApp.meal"
}

func (m Meal) SumCalories(items []MealItem) int {
	cals := 0.0
	for _, item := range items {
		cals += item.Calories
	}
	return int(cals)
}

func (m Meal) Filter(groupName string) []MealItem {
	count := len(m.Children)
	emptyList := make([]MealItem, 0)

	if count == 0 || groupName == "" {
		return emptyList
	}

	if m.Groups == nil {
		m.Groups = make(map[string][]MealItem)
	}

	cachedGroup, ok := m.Groups[groupName]
	if ok {
		return cachedGroup
	}

	// Start the filtering
	// Find indices of the group
	firstIndex := -1
	lastIndex := -1

	for i := 0; i < count; i++ {
		if m.Children[i].GroupName == groupName {
			firstIndex = i
		}
		if firstIndex >= 0 && m.Children[i].GroupName != "" && m.Children[i].GroupName != groupName {
			lastIndex = i
			break
		}
	}

	if firstIndex == -1 {
		return emptyList
	}

	if lastIndex == -1 {
		// lastIndex should be =count, not count - 1
		lastIndex = count
	}

	m.Groups[groupName] = m.Children[firstIndex:lastIndex]

	return m.Groups[groupName]
}

func (m Meal) GenerateGroupNames() {
	if len(m.Children) == 0 {
		return
	}

	makeGroup := func(o *MealItem) string {
		return goma.ParseDatetime(o.Timestamp).Format("Jan _2")
	}

	m.Children[0].GroupName = makeGroup(&m.Children[0])
	oldGroup := m.Children[0].GroupName

	for i := 1; i < len(m.Children); i++ {
		m.Children[i].GroupName = ""

		group := makeGroup(&m.Children[i])
		if oldGroup != group {
			m.Children[i].GroupName = group
		}
		oldGroup = group
	}
}

func EmptyMealContainer() *Meal {
	return &Meal{Children: make([]MealItem, 0)}
}

func EmptyWeightLogContainer() *WeightLog {
	return &WeightLog{Children: make([]WeightLogItem, 0)}
}

func AddNewMeal(timestamp string, description string, calories string, unit string) (*MealItem, error) {
	if strings.TrimSpace(description) == "" {
		return nil, errors.New("Please add a description of the meal.")
	}

	if strings.TrimSpace(calories) == "" {
		return nil, errors.New("Calories cannot be empty.")
	}

	cals, err := NewFormula(calories).Evaluate()
	if err != nil {
		return nil, err
	}

	if unit == "KJoules" {
		cals = cals * KJ_TO_KCAL_FACTOR
	}

	meal := goma.GetAppCache().Get(EmptyMealContainer()).(*Meal)

	item := MealItem{
		Timestamp:   timestamp,
		Description: description,
		Calories:    math.Ceil(cals),
	}

	meal.Children = append(meal.Children, item)

	err = meal.Save(meal)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateMeal(id string, description string, calories string) error {
	cals, err := NewFormula(calories).Evaluate()
	if err != nil {
		return err
	}

	meal := goma.GetAppCache().Get(EmptyMealContainer()).(*Meal)

	for i := 0; i < len(meal.Children); i++ {
		if meal.Children[i].Timestamp == id {
			mi := meal.Children[i]
			mi.Description = description
			mi.Calories = cals
			meal.Children[i] = mi
			break
		}
	}

	return meal.Save(meal)
}

func DeleteMeal(id string) error {
	meal := goma.GetAppCache().Get(EmptyMealContainer()).(*Meal)
	for i := 0; i < len(meal.Children); i++ {
		if meal.Children[i].Timestamp == id {
			meal.Children = append(meal.Children[:i], meal.Children[i+1:]...)
		}
	}

	return meal.Save(meal)
}

func GetMeals() MealList {
	var meal *Meal
	meal = goma.GetAppCache().Get(EmptyMealContainer()).(*Meal)
	if len(meal.Children) == 0 {
		meal = &Meal{}
		meal.Restore(meal)
		goma.GetAppCache().Put(meal)
	}

	sort.Sort(meal.Children)

	meal.GenerateGroupNames()

	return meal.Children
}

func TotalCaloriesByGroup(groupName string) string {
	meal := goma.GetAppCache().Get(EmptyMealContainer()).(*Meal)
	sum := meal.SumCalories(meal.Filter(groupName))
	if sum == 0 {
		return ""
	}
	return fmt.Sprintf("%d", sum)
}
