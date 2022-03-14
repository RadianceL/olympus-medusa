package routers

import (
	"github.com/gin-gonic/gin"
	"medusa-globalization-copywriting-system/cmd/handler"
)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := handler.RestHandler{}
	// 创建应用
	app.POST("/application/create", menu.CreateApplication)
	// 创建应用文案
	app.POST("/application/document/create", menu.CreateGlobalizationCopyWriting)
	// 更新应用文案
	app.POST("/application/document/update", menu.UpdateGlobalizationCopyWriting)
	// 审核应用文案
	app.POST("/application/document/commit", menu.CommitGlobalizationCopyWriting)
	// 查询应用文案结构
	app.POST("/application/document/list/struct", menu.ListGlobalizationCopyWritingStruct)
	// 查询应用文案命名空间
	app.POST("/application/document/list/namespace", menu.ListGlobalizationCopyWritingNamespace)
	// 查询应用文案历史记录
	app.POST("/application/document/query/history", menu.ListGlobalizationCopyWritingHistory)
}
