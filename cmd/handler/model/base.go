package model

import (
	"database/sql"
	"medusa-globalization-copywriting-system/cmd/datasource"
)

// Base is base model structure.
type Base struct {
	TableName string

	Conn datasource.Connection
	Tx   *sql.Tx
}

func (b Base) SetConn(con datasource.Connection) Base {
	b.Conn = con
	return b
}

func (b Base) Table(table string) *datasource.SQL {
	return datasource.Table(table).WithDriver(b.Conn)
}
