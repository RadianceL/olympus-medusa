package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	apiPrefix := "/document/api"
	ginGroup := app.Group(apiPrefix)
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/logout")
	//ginGroup.Use(middleware.Validate(
	//	middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	//))

	// 权限验证
	var notCheckPermissionUrlArr []string
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	RegisterRouterSys(ginGroup)
}
