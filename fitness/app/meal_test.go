package app

import (
	goma "github.com/hemantasapkota/goma"
	"testing"
	"time"
)

func d(format string) time.Duration {
	d0, err := time.ParseDuration(format)
	if err != nil {
		d0, err = time.ParseDuration("24h")
	}
	return d0
}

func next(t0 time.Time, duration time.Duration) time.Time {
	return t0.Add(duration)
}

func TestMealFeed(t *testing.T) {
	dbPath := "."
	Init(dbPath)

	// Start with an empty meal container
	meal := &Meal{}
	// overwrite any existing data
	meal.Save(meal)

	goma.GetAppCache().Put(meal)

	t0 := time.Now()

	t1 := next(t0, d("-72h"))
	AddNewMeal(goma.TimestampFrom(t1), "200G Avocado", "320", "kcal")
	AddNewMeal(goma.TimestampFrom(t1), "100G Lollies", "400", "kcal")
	AddNewMeal(goma.TimestampFrom(t1), "200G Carrots", "82", "kcal")

	t2 := next(t0, d("-48h"))
	AddNewMeal(goma.TimestampFrom(t2), "50G Avocado", "110", "kcal")
	AddNewMeal(goma.TimestampFrom(t2), "20G Carrots", "82", "kcal")

	t3 := next(t0, d("-24h"))
	AddNewMeal(goma.TimestampFrom(t3), "50G Rice", "65", "kcal")
	AddNewMeal(goma.TimestampFrom(t3), "20G Coconut Oil", "172", "kcal")

	AddNewMeal(goma.TimestampFrom(t0), "500G Chicken Breat", "650", "kcal")
	AddNewMeal(goma.TimestampFrom(t0), "500G Rockmelon", "170", "kcal")

	GetMeals().PrettyPrint()
}
