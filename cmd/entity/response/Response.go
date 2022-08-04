package response

import (
	"net/http"
	"olympus-medusa/tools/request"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode = "0000" /* 成功的状态码 */
	FailCode    = "1001" /* 失败的状态码 */
)

type Model struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ModelBase struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := Model{Code: SuccessCode, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// ResSuccessMsg 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ModelBase{Code: SuccessCode, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFail 响应失败
func ResFail(context *gin.Context, msg string) {
	ret := ModelBase{Code: FailCode, Message: msg}
	ResJSON(context, http.StatusOK, &ret)
}

// ResFailCode 响应失败
func ResFailCode(context *gin.Context, msg string, code string) {
	ret := ModelBase{Code: code, Message: msg}
	ResJSON(context, http.StatusOK, &ret)
}

// ResJSON 响应JSON数据
func ResJSON(context *gin.Context, status int, v interface{}) {
	context.JSON(status, v)
	context.Abort()
}

// ResErrSrv 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ModelBase{Code: FailCode, Message: "服务端故障"}
	ResJSON(c, http.StatusOK, &ret)
}

// ResErrCli 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ModelBase{Code: FailCode, Message: err.Error()}
	ResJSON(c, http.StatusOK, &ret)
}

type PageData struct {
	Total uint64      `json:"total"`
	Items interface{} `json:"items"`
}

type Page struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    PageData `json:"data"`
}

// ResSuccessPage 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total uint64, list interface{}) {
	ret := Page{Code: SuccessCode, Message: "ok", Data: PageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

// GetPageIndex 获取页码
func GetPageIndex(c *gin.Context) uint64 {
	return request.GetQueryToUint64(c, "page", 1)
}

// GetPageLimit 获取每页记录数
func GetPageLimit(c *gin.Context) uint64 {
	limit := request.GetQueryToUint64(c, "limit", 20)
	if limit > 500 {
		limit = 20
	}
	return limit
}

// GetPageSort 获取排序信息
func GetPageSort(c *gin.Context) string {
	return request.GetQueryToStr(c, "sort")
}

// GetPageKey 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	return request.GetQueryToStr(c, "key")
}
