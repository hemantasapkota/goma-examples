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

func UpdateRecord(prevTimestamp string, newTimestamp string, description string, value string) error {
	return app.UpdateRecord(prevTimestamp, newTimestamp, description, value)
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

func IndexOfUnit(value string) int {
	return app.IndexOfUnit(value)
}
