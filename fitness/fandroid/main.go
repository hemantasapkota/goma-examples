package fandroid

import (
	goma "github.com/hemantasapkota/goma"
	app "goma-examples/fitness/app"
)

func Init(dbPath string) error {
	return app.Init(dbPath)
}

func AddNewRecord(timestamp string, description string, value string, unit string) error {
	_, err := app.AddNewRecord(timestamp, description, value, unit)
	return err
}

func UpdateRecord(id string, newTimestamp string, description string, calories string) error {
	return app.UpdateRecord(id, newTimestamp, description, calories)
}

func GetRecords() ([]byte, error) {
	return goma.Marshal(app.GetRecords())
}

func DeleteRecord(id string) error {
	return app.DeleteRecord(id)
}

func TotalCaloriesByGroup(groupName string) string {
	return app.TotalCaloriesByGroup(groupName)
}

func GetUnits() ([]byte, error) {
	return goma.Marshal(app.GetUnits())
}
