package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"medusa-globalization-copywriting-system/cmd/datasource"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
	Response "medusa-globalization-copywriting-system/cmd/entity/response"
	"medusa-globalization-copywriting-system/cmd/handler/model"
	"medusa-globalization-copywriting-system/tools/logger"
)

type RestHandler struct{}

// CreateApplication 创建多语言应用/**
func (result RestHandler) CreateApplication(context *gin.Context) {
	applicationAddRequest := &Entity.ApplicationAddRequest{}
	err := context.ShouldBindBodyWith(&applicationAddRequest, binding.JSON)
	if err != nil {
		return
	}
	applicationModel := model.Application().SetConn(datasource.Conn)
	_, err = applicationModel.AddApplication("123")
	if err != nil {
		Response.ResErrCli(context, err)
		return
	}

	Response.ResSuccessMsg(context)
}

// CreateGlobalizationCopyWriting 创建多语言文案/**
func (result RestHandler) CreateGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// UpdateGlobalizationCopyWriting 更新多语言文案/**
func (result RestHandler) UpdateGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// CommitGlobalizationCopyWriting 提交多语言文案更新/**
func (result RestHandler) CommitGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWritingNamespace 查询应用文案命名空间/**
func (result RestHandler) ListGlobalizationCopyWritingNamespace(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWritingStruct 创建多语言文案/**
func (result RestHandler) ListGlobalizationCopyWritingStruct(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWriting 创建多语言文案/**
func (result RestHandler) ListGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWritingHistory 创建多语言文案/**
func (result RestHandler) ListGlobalizationCopyWritingHistory(context *gin.Context) {
	json := &Entity.DocumentAddRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	if json.Path == "" {
		Response.ResFail(context, "")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}
