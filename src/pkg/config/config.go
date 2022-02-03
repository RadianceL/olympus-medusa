package config

import (
	"fmt"
)

type Config struct {
	Web        Web
	DataSource DataSource
	MySQL      MySQL
	SQLite     SQLite
}

type Web struct {
	LogLevel     string
	Domain       string
	StaticPath   string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DataSource struct {
	Debug              bool
	DBType             string
	MaxLifetime        int
	MaxOpenConnections int
	MaxIdleConnections int
	TablePrefix        string
	DSN                string
}

type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

func (mysql MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.DBName, mysql.Parameters)
}

type SQLite struct {
	Path string
}

func (sqlite SQLite) DSN() string {
	return sqlite.Path
}
