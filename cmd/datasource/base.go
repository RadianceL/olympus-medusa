package datasource

import (
	"database/sql"
	"medusa-globalization-copywriting-system/cmd/config"
	"sync"
)

// Base is a common Connection.
type Base struct {
	DbList  map[string]*sql.DB
	Once    sync.Once
	Configs config.DataSource
}

// Close implements the method Connection.Close.
func (db *Base) Close() []error {
	errs := make([]error, 0)
	for _, d := range db.DbList {
		errs = append(errs, d.Close())
	}
	return errs
}

// GetDB implements the method Connection.GetDB.
func (db *Base) GetDB(key string) *sql.DB {
	return db.DbList[key]
}
