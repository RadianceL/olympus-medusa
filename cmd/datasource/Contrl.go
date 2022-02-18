package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"medusa-globalization-copywriting-system/cmd/config"
	"medusa-globalization-copywriting-system/tools/logger"
	"os"
	"time"
)

func ConnectionDatabase(config config.DataSource) {
	var gdb *gorm.DB
	var err error
	if config.DBType == "mysql" {
		config.DSN = config.MySQL.DSN()
	} else if config.DBType == "sqlite3" {
		config.DSN = config.SQLite.DSN()
	}
	logger.Info(config.DSN)
	gdb, err = gorm.Open(config.DBType, config.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if config.Debug {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(config.MaxIdleConnections)
	gdb.DB().SetMaxOpenConns(config.MaxOpenConnections)
	gdb.DB().SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	logger.Debug("数据库加载完成.......")
}
