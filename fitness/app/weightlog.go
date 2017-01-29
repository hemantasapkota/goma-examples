package app

import (
	"errors"
	"github.com/hemantasapkota/goma"
)

type WeightLogItem struct {
	Timestamp string  `json:"timestamp"`
	Weight    float64 `json:"weight"`
}

type WeightLog struct {
	*goma.Object

	Children []WeightLogItem `json:"children"`
}

func (w *WeightLog) Key() string {
	return "fitnessApp.weightLog"
}

func LogWeight(timestamp string, weight string) error {
	if timestamp == "" {
		return errors.New("Please select a date.")
	}

	if weight == "" {
		return errors.New("Weight cannot be empty.")
	}

	weightVal, err := NewFormula(weight).Evaluate()
	if err != nil {
		return err
	}

	weightLog := goma.GetAppCache().Get(EmptyWeightLogContainer()).(*WeightLog)

	weightLog.Children = append(weightLog.Children, WeightLogItem{timestamp, weightVal})

	err = weightLog.Save(weightLog)
	if err != nil {
		return err
	}

	return nil
}

func GetWeightLog() []WeightLogItem {
	var weightLog *WeightLog
	weightLog = goma.GetAppCache().Get(EmptyWeightLogContainer()).(*WeightLog)
	if len(weightLog.Children) == 0 {
		weightLog = &WeightLog{}
		weightLog.Restore(weightLog)
		goma.GetAppCache().Put(weightLog)
	}

	return weightLog.Children
}
