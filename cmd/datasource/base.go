package datasource

import (
	"github.com/jinzhu/gorm"
	"medusa-globalization-copywriting-system/cmd/config"
	"sync"
)

// Base is a common Connection.
type Base struct {
	DbList  map[string]*gorm.DB
	Once    sync.Once
	Configs config.DataSource
}

// GetDB implements the method Connection.GetDB.
func (db *Base) GetDB(key string) *gorm.DB {
	return db.DbList[key]
}
