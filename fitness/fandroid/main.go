package fandroid

import (
	goma "goma"
	app "goma/examples/fitness_app/app"
)

func Init(dbPath string) error {
	return app.Init(dbPath)
}

func AddNewMeal(description string, calories string) error {
	_, err := app.AddNewMeal(goma.Timestamp(), description, calories)
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
