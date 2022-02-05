package routers

import (
	"github.com/gin-gonic/gin"
	"medusa-globalization-copywriting-system/cmd/web/handler"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := handler.RestHandler{}
	app.POST("/menu/list", menu.CreateGlobalizationCopyWriting)
}
