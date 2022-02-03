package routers

import (
	"github.com/gin-gonic/gin"
	"medusa-globalization-copywriting-system/cmd/web/controller"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := controller.Result{}
	app.GET("/menu/list", menu.List)
}
