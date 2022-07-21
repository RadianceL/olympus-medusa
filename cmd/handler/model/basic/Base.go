package basic

import (
	"database/sql"
	"olympus-medusa/cmd/datasource"
)

// BaseModel is basic model structure.
type BaseModel struct {
	TableName string

	Conn datasource.Connection
	Tx   *sql.Tx
}

func (b BaseModel) SetConn(con datasource.Connection) BaseModel {
	b.Conn = con
	return b
}

func (b BaseModel) Table(table string) *datasource.SQL {
	return datasource.Table(table).WithDriver(b.Conn)
}
