package AppConfig

import (
	"github.com/spf13/viper"
	"medusa-globalization-copywriting-system/src/pkg/config"
)

// LoadConfig 加载配置
func LoadConfig(configPath string) (c *config.Config, err error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &config.Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.LogLevel = v.GetString("web.log_level")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.SQLite.Path = v.GetString("sqlite3.path")
	c.MySQL.Host = v.GetString("mysql.host")
	c.MySQL.Port = v.GetInt("mysql.port")
	c.MySQL.User = v.GetString("mysql.user")
	c.MySQL.Password = v.GetString("mysql.password")
	c.MySQL.DBName = v.GetString("mysql.db_name")
	c.MySQL.Parameters = v.GetString("mysql.parameters")
	c.DataSource.Debug = v.GetBool("gorm.debug")
	c.DataSource.DBType = v.GetString("gorm.db_type")
	c.DataSource.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.DataSource.MaxOpenConnections = v.GetInt("gorm.max_open_conns")
	c.DataSource.MaxIdleConnections = v.GetInt("gorm.max_idle_conns")
	c.DataSource.TablePrefix = v.GetString("gorm.table_prefix")
	return
}
