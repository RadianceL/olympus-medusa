package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	Entity "medusa-globalization-copywriting-system/cmd/entity"
	Response "medusa-globalization-copywriting-system/cmd/entity/response"
	"medusa-globalization-copywriting-system/tools/logger"
)

type RestHandler struct{}

var DB *gorm.DB

func (result RestHandler) CreateGlobalizationCopyWriting(context *gin.Context) {
	json := make(map[string]string)
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		return
	}
	logger.Info("aaa{}", json)
	email := &Entity.CopywritingAddRequest{}
	DB.Save(email)
	Response.ResSuccessMsg(context)
}
