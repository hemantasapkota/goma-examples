package app

import (
	"goma"
)

type WorkoutScheme struct {
	*goma.Object

	Name string `json:"name"`
	Sets string `json:"sets"`
	Reps string `json:"reps"`
}

func (w WorkoutScheme) Key() string {
	return "fitnessApp.workoutScheme"
}

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
