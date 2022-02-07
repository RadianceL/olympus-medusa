package config

import (
	"github.com/spf13/viper"
)

// LoadConfig 加载配置
func LoadConfig(configPath string) (c *Config, resultError error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		resultError = err
		return
	}
	c = &Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.DataSource.SQLite.Path = v.GetString("sqlite3.path")
	c.DataSource.MySQL.Host = v.GetString("gorm.mysql.host")
	c.DataSource.MySQL.Port = v.GetInt("gorm.mysql.port")
	c.DataSource.MySQL.User = v.GetString("gorm.mysql.user")
	c.DataSource.MySQL.Password = v.GetString("gorm.mysql.password")
	c.DataSource.MySQL.DBName = v.GetString("gorm.mysql.db_name")
	c.DataSource.MySQL.Parameters = v.GetString("gorm.mysql.parameters")
	c.DataSource.Debug = v.GetBool("gorm.debug")
	c.DataSource.DBType = v.GetString("gorm.db_type")
	c.DataSource.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.DataSource.MaxOpenConnections = v.GetInt("gorm.max_open_conns")
	c.DataSource.MaxIdleConnections = v.GetInt("gorm.max_idle_conns")
	c.DataSource.TablePrefix = v.GetString("gorm.table_prefix")
	return
}
