package controller

import (
	"github.com/gin-gonic/gin"
	models "medusa-globalization-copywriting-system/src/pkg"
	"medusa-globalization-copywriting-system/src/web/entity"
	"medusa-globalization-copywriting-system/src/web/entity/response"
)

type Result struct{}

func (Result) List(c *gin.Context) {
	email := entity.Email{ID: 10, UserID: 1111, Email: "931305033@qq.com", Subscribed: true}
	_ = models.Create(email)
	response.ResSuccessMsg(c)
}
