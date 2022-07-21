package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite driver.
	"net/http"
	"olympus-medusa/cmd/config"
	"olympus-medusa/cmd/datasource"
	"olympus-medusa/cmd/handler/model"
	"olympus-medusa/cmd/middleware"
	"olympus-medusa/cmd/router"
	"olympus-medusa/tools/convert"
	"olympus-medusa/tools/logger"
	"time"
)

var (
	tomlFile = flag.String("config", "./configs/application.yaml", "config file")
)

// Run 运行
func main() {
	// 加载配置
	flag.Parse()
	loadConfig, err := config.LoadConfig(*tomlFile)
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
