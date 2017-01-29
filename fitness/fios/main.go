package fios

import (
	goma "github.com/hemantasapkota/goma"
	app "goma-examples/fitness/app"
)

func Init(dbPath string) error {
	return app.Init(dbPath)
}
