// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package datasource

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
	"medusa-globalization-copywriting-system/cmd/config"
	"medusa-globalization-copywriting-system/tools/logger"
	"os"
	"time"
)

// SQLTx is an in-progress database transaction.
type SQLTx struct {
	Tx *sql.Tx
}

// Mysql is a Connection of mysql.
type Mysql struct {
	Base
}

func (db *Mysql) Close() []error {
	//TODO implement me
	panic("implement me")
}

func (db *Mysql) GetDB(key string) *sql.DB {
	//TODO implement me
	panic("implement me")
}

func (db *Mysql) GetConfig(name string) config.DataSource {
	//TODO implement me
	panic("implement me")
}

func (db *Mysql) CreateDB(name string, beans ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

// GetMysqlDB return the global mysql connection.
func GetMysqlDB() *Mysql {
	return &Mysql{
		Base: Base{
			DbList: make(map[string]*gorm.DB),
		},
	}
}

func (db *Mysql) InitDB(config config.DataSource) Connection {
	db.Configs = config
	db.Once.Do(func() {
		var gdb *gorm.DB
		var err error
		config.DSN = config.MySQL.DSN()
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

		db.DbList["default"] = gdb

		if err := gdb.DB().Ping(); err != nil {
			panic(err)
		}
	})
	return db
}

// Name implements the method Connection.Name.
func (db *Mysql) Name() string {
	return "mysql"
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *Mysql) GetDelimiter() string {
	return "`"
}

// GetDelimiter2 implements the method Connection.GetDelimiter2.
func (db *Mysql) GetDelimiter2() string {
	return "`"
}

// GetDelimiters implements the method Connection.GetDelimiters.
func (db *Mysql) GetDelimiters() []string {
	return []string{"`", "`"}
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *Mysql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *Mysql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

// Query implements the method Connection.Query.
func (db *Mysql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

// Exec implements the method Connection.Exec.
func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

// QueryWithTx is query method within the transaction.
func (db *Mysql) QueryWithTx(tx *sql.Tx, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQueryWithTx(tx, query, args...)
}

// ExecWithTx is exec method within the transaction.
func (db *Mysql) ExecWithTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return CommonExecWithTx(tx, query, args...)
}

func (db *Mysql) QueryWith(tx *sql.Tx, conn, query string, args ...interface{}) ([]map[string]interface{}, error) {
	if tx != nil {
		return db.QueryWithTx(tx, query, args...)
	}
	return db.QueryWithConnection(conn, query, args...)
}

func (db *Mysql) ExecWith(tx *sql.Tx, conn, query string, args ...interface{}) (sql.Result, error) {
	if tx != nil {
		return db.ExecWithTx(tx, query, args...)
	}
	return db.ExecWithConnection(conn, query, args...)
}

// BeginTxWithReadUncommitted starts a transaction with level LevelReadUncommitted.
func (db *Mysql) BeginTxWithReadUncommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommitted starts a transaction with level LevelReadCommitted.
func (db *Mysql) BeginTxWithReadCommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableRead starts a transaction with level LevelRepeatableRead.
func (db *Mysql) BeginTxWithRepeatableRead() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelRepeatableRead)
}

// BeginTx starts a transaction with level LevelDefault.
func (db *Mysql) BeginTx() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelDefault)
}

// BeginTxWithLevel starts a transaction with given transaction isolation level.
func (db *Mysql) BeginTxWithLevel(level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], level)
}

// BeginTxWithReadUncommittedAndConnection starts a transaction with level LevelReadUncommitted and connection.
func (db *Mysql) BeginTxWithReadUncommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommittedAndConnection starts a transaction with level LevelReadCommitted and connection.
func (db *Mysql) BeginTxWithReadCommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableReadAndConnection starts a transaction with level LevelRepeatableRead and connection.
func (db *Mysql) BeginTxWithRepeatableReadAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelRepeatableRead)
}

// BeginTxAndConnection starts a transaction with level LevelDefault and connection.
func (db *Mysql) BeginTxAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelDefault)
}

// BeginTxWithLevelAndConnection starts a transaction with given transaction isolation level and connection.
func (db *Mysql) BeginTxWithLevelAndConnection(conn string, level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], level)
}
