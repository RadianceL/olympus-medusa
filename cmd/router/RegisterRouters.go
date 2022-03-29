package routers

import (
	"github.com/gin-gonic/gin"
	"medusa-globalization-copywriting-system/cmd/handler"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := handler.RestHandler{}
	// 创建应用
	app.POST("/application/create", menu.CreateApplication)
	// 查询应用列表
	app.POST("/application/list", menu.ListApplication)
	// 创建应用空间namespace
	app.POST("/application/namespace/create", menu.CreateGlobalizationCopyWritingNamespace)
	// 查询应用文案结构（返回该应用空间下的namespace -tree）
	app.POST("/application/namespace/list", menu.ListGlobalizationCopyWritingStruct)
	// 查询应用文案命名空间
	app.POST("/application/namespace/document/list", menu.ListGlobalizationCopyWritingNamespace)
	// 创建应用文案
	app.POST("/application/document/create", menu.CreateGlobalizationCopyWriting)
	// 更新应用文案
	app.POST("/application/document/update", menu.UpdateGlobalizationCopyWriting)
	// 审核应用文案
	app.POST("/application/document/commit", menu.CommitGlobalizationCopyWriting)
	// 查询应用文案历史记录
	app.POST("/application/document/query/history", menu.ListGlobalizationCopyWritingHistory)
}
