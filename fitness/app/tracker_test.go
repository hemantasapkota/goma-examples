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

func TestRecordFeed(t *testing.T) {
	dbPath := "."
	Init(dbPath)

	// Start with an empty meal container
	tracker := &Tracker{}
	// overwrite any existing data
	tracker.Save(tracker)

	goma.GetAppCache().Put(tracker)

	t0 := time.Now()

	t1 := next(t0, d("-72h"))
	AddNewRecord(goma.TimestampFrom(t1), "200G Avocado", "320", "KCals")
	AddNewRecord(goma.TimestampFrom(t1), "100G Lollies", "400", "KCals")
	AddNewRecord(goma.TimestampFrom(t1), "200G Carrots", "82", "KCals")

	t2 := next(t0, d("-48h"))
	AddNewRecord(goma.TimestampFrom(t2), "50G Avocado", "110", "KCals")
	AddNewRecord(goma.TimestampFrom(t2), "20G Carrots", "82", "KCals")

	t3 := next(t0, d("-24h"))
	AddNewRecord(goma.TimestampFrom(t3), "50G Rice", "65", "KCals")
	AddNewRecord(goma.TimestampFrom(t3), "20G Coconut Oil", "172", "KCals")

	AddNewRecord(goma.TimestampFrom(t0), "500G Chicken Breat", "650", "KCals")
	AddNewRecord(goma.TimestampFrom(t0), "500G Rockmelon", "170", "KCals")

	GetRecords().PrettyPrint()
}
