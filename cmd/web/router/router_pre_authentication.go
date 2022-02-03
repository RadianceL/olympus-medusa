package routers

import (
	"medusa-globalization-copywriting-system/cmd/middleware"

	//"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	//首页
	app.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	apiPrefix := "/api"
	ginGroup := app.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/logout")
	ginGroup.Use(middleware.Validate(
		middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	))

	// 权限验证
	var notCheckPermissionUrlArr []string
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	RegisterRouterSys(ginGroup)
}
