package fandroid

import (
	goma "github.com/hemantasapkota/goma"
	app "goma-examples/fitness/app"
)

func Init(dbPath string) error {
	return app.Init(dbPath)
}

func AddNewMeal(description string, calories string, unit string) error {
	_, err := app.AddNewMeal(goma.Timestamp(), description, calories, unit)
	return err
}

func UpdateMeal(id string, description string, calories string) error {
	return app.UpdateMeal(id, description, calories)
}

func GetMeals() ([]byte, error) {
	return goma.Marshal(app.GetMeals())
}

func DeleteMeal(id string) error {
	return app.DeleteMeal(id)
}

func TotalCaloriesByGroup(groupName string) string {
	return app.TotalCaloriesByGroup(groupName)
}

func GetUnits() ([]byte, error) {
	return goma.Marshal(app.GetUnits())
}
