package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	Models "medusa-globalization-copywriting-system/cmd/datasource"
	Entity "medusa-globalization-copywriting-system/cmd/web/entity"
	Response "medusa-globalization-copywriting-system/cmd/web/entity/response"
	"medusa-globalization-copywriting-system/tools/logger"
)

type RestHandler struct{}

func (result RestHandler) CreateGlobalizationCopyWriting(context *gin.Context) {
	json := make(map[string]string)
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		return
	}
	logger.Info("aaa{}", json)
	email := &Entity.Email{ID: 10, UserID: 1111, Email: json["111"], Subscribed: true}
	Models.DB.Save(email)
	Response.ResSuccessMsg(context)
}
