package app

import (
	goma "goma"
	ldb "goma/gomadb/leveldb"
)

func Init(dbPath string) error {
	goma.NewLogger()
	_, err := ldb.InitDB(dbPath)
	return err
}
