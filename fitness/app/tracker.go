package app

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/hemantasapkota/goma"
)

type RecordList []Record

type Record struct {
	Timestamp   string  `json:"timestamp"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`

	// View releated stuffs
	GroupName string `json:"group"`
}

// Sort interface

func (m RecordList) Len() int {
	return len(m)
}

func (m RecordList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m RecordList) Less(i, j int) bool {
	return m[i].Timestamp > m[j].Timestamp
}

// End sort

func (m RecordList) PrettyPrint() {
	for _, item := range m {
		group := item.GroupName
		if len(group) == 0 {
			group = "\t"
		}

		goma.Log(fmt.Sprintf("%s %s %f", group, item.Description, item.Value))
	}
}

type Tracker struct {
	*goma.Object
	Children RecordList `json:"children"`

	// View related stuffs
	Groups map[string][]Record
}

func (m Tracker) Key() string {
	return "fitnessApp.tracker"
}

func (m Tracker) SumCalories(items []Record) int {
	cals := 0.0
	for _, item := range items {
		if item.Unit == "KCals" {
			cals += item.Value
		}
	}
	return int(cals)
}

func (m Tracker) Filter(groupName string) []Record {
	count := len(m.Children)
	emptyList := make([]Record, 0)

	if count == 0 || groupName == "" {
		return emptyList
	}

	if m.Groups == nil {
		m.Groups = make(map[string][]Record)
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

func (m Tracker) GenerateGroupNames() {
	if len(m.Children) == 0 {
		return
	}

	makeGroup := func(o *Record) string {
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

func EmptyTrackerContainer() *Tracker {
	return &Tracker{Children: make([]Record, 0)}
}

func AddNewRecord(timestamp string, description string, value string, unit string) (*Record, error) {
	if strings.TrimSpace(description) == "" {
		return nil, errors.New("Please add a description of the record.")
	}

	if strings.TrimSpace(value) == "" {
		return nil, errors.New("Value cannot be empty.")
	}

	val, err := NewFormula(value).Evaluate()
	if err != nil {
		return nil, err
	}

	if unit == "KJoules" {
		// convert to calories
		val = val * KJ_TO_KCAL_FACTOR
	}

	tracker := goma.GetAppCache().Get(EmptyTrackerContainer()).(*Tracker)

	item := Record{
		Timestamp:   timestamp,
		Description: description,
		Value:       val,
		Unit:        unit,
	}

	tracker.Children = append(tracker.Children, item)

	err = tracker.Save(tracker)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateRecord(prevTimestamp string, newtimestamp string, description string, value string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("Please add a description of the record.")
	}

	if strings.TrimSpace(value) == "" {
		return errors.New("Value cannot be empty.")
	}

	println(value)

	val, err := NewFormula(value).Evaluate()
	if err != nil {
		return err
	}

	tracker := goma.GetAppCache().Get(EmptyTrackerContainer()).(*Tracker)

	for i := 0; i < len(tracker.Children); i++ {
		if tracker.Children[i].Timestamp == prevTimestamp {
			mi := tracker.Children[i]
			mi.Timestamp = newtimestamp
			mi.Description = description
			mi.Value = val
			tracker.Children[i] = mi
			break
		}
	}

	return tracker.Save(tracker)
}

func DeleteRecord(id string) error {
	tracker := goma.GetAppCache().Get(EmptyTrackerContainer()).(*Tracker)
	for i := 0; i < len(tracker.Children); i++ {
		if tracker.Children[i].Timestamp == id {
			tracker.Children = append(tracker.Children[:i], tracker.Children[i+1:]...)
		}
	}

	return tracker.Save(tracker)
}

func GetRecords() RecordList {
	var tracker *Tracker
	tracker = goma.GetAppCache().Get(EmptyTrackerContainer()).(*Tracker)
	if len(tracker.Children) == 0 {
		tracker = &Tracker{}
		tracker.Restore(tracker)
		goma.GetAppCache().Put(tracker)
	}

	sort.Sort(tracker.Children)

	tracker.GenerateGroupNames()

	return tracker.Children
}

func TotalCaloriesByGroup(groupName string) string {
	tracker := goma.GetAppCache().Get(EmptyTrackerContainer()).(*Tracker)
	sum := tracker.SumCalories(tracker.Filter(groupName))
	if sum == 0 {
		return ""
	}
	return fmt.Sprintf("%d", sum)
}
