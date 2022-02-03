package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type MysqlConnectionPool struct {
}

var instance *MysqlConnectionPool
var once sync.Once

var _ *gorm.DB
var err error

func GetInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDataPool() (iss bool) {
	_, err = gorm.Open("mysql", "user:password@tcp(192.168.1.4:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (m *MysqlConnectionPool) GetMysqlDB() (db *gorm.DB) {
	return db
}
