package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite driver.
	"medusa-globalization-copywriting-system/cmd/config"
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/handler/model"
	"medusa-globalization-copywriting-system/cmd/middleware"
	"medusa-globalization-copywriting-system/cmd/router"
	"medusa-globalization-copywriting-system/tools/convert"
	"medusa-globalization-copywriting-system/tools/logger"
	"net/http"
	"time"
)

const configPath = "./configs/application.yaml"

// Run 运行
func main() {
	// 加载配置
	loadConfig, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	// 初始化日志配置
	logger.InitLog("debug", "./build/data/log/log.log")
	// 初始化数据
	datasource.Conn = datasource.GetConnectionByDriver(loadConfig.DataSource.DBType).InitDB(loadConfig.DataSource)
	model.RepositoryModelContainer.InitializationAll()
	// 初始化web服务
	initWeb(loadConfig)
}

func initWeb(config *config.Config) {
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	app.NoRoute(middleware.NoRouteHandler())
	app.NoMethod(middleware.NoMethodHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	// 注册路由
	routers.RegisterRouter(app)

	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}
