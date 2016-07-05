package app

import (
	goma "github.com/hemantasapkota/goma"
	gomadb "github.com/hemantasapkota/goma/gomadb"
	ldb "github.com/hemantasapkota/goma/gomadb/leveldb"
)

func Init(dbPath string) error {
	goma.NewLogger()
	db, err := ldb.InitDB(dbPath)
	if err != nil {
		return err
	}
	gomadb.SetLevelDB(db)
	return nil
}
