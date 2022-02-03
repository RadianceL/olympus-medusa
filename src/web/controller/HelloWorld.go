package controller

import (
	"github.com/gin-gonic/gin"
	Models "medusa-globalization-copywriting-system/src/pkg"
	Entity "medusa-globalization-copywriting-system/src/web/entity"
	Response "medusa-globalization-copywriting-system/src/web/entity/response"
)

type Result struct{}

func (result Result) List(context *gin.Context) {
	email := Entity.Email{ID: 10, UserID: 1111, Email: "931305033@qq.com", Subscribed: true}
	_ = Models.Create(email)
	Response.ResSuccessMsg(context)
}
