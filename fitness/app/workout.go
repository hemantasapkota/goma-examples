package app

import (
	goma "github.com/hemantasapkota/goma"
)

type Workout struct {
	*goma.Object

	Starttime string `json:"startTime"`
	Endtime   string `json:"endTime`
	Day       string `json:"day"`
}

func (w Workout) Key() string {
	return "fitnessApp.workout"
}

func StartNewWorkout() (*Workout, error) {
	w := Workout{
		Starttime: "12312",
		Endtime:   "12312",
		Day:       "today",
	}

	// Save out object to the db
	err := w.Save(w)
	if err != nil {
		return nil, err
	}

	return &w, nil
}
