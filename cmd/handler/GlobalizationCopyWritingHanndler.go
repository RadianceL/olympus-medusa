package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	Entity "olympus-medusa/cmd/entity/request"
	Response "olympus-medusa/cmd/entity/response"
	"olympus-medusa/cmd/handler/model"
	"olympus-medusa/tools/logger"
)

type RestHandler struct{}

// CreateApplication 创建多语言应用/**
func (restHandler RestHandler) CreateApplication(context *gin.Context) {
	applicationAddRequest := &Entity.ApplicationRequest{}
	err := context.ShouldBindBodyWith(&applicationAddRequest, binding.JSON)
	if applicationAddRequest.ApplicationName == "" {
		Response.ResFail(context, "应用名称不能为空")
		return
	}
	if err != nil {
		Response.ResErrCli(context, err)
		return
	}
	_, err = model.ApplicationHandler.AddApplication(applicationAddRequest)
	if err != nil {
		Response.ResErrCli(context, err)
		return
	}
	Response.ResSuccessMsg(context)
}

// ListApplication 查询应用列表/**
func (restHandler RestHandler) ListApplication(context *gin.Context) {
	applicationAddRequest := &Entity.ApplicationRequest{}
	shouldBindBodyWithErr := context.ShouldBindBodyWith(&applicationAddRequest, binding.JSON)
	if shouldBindBodyWithErr != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	searchApplicationList, searchApplicationError := model.ApplicationHandler.SearchApplicationList(applicationAddRequest)
	if searchApplicationError != nil {
		logger.Error(searchApplicationError)
		Response.ResFail(context, "应用处理异常")
		return
	}
	Response.ResSuccess(context, searchApplicationList)
}

// CreateGlobalizationCopyWritingNamespace 创建应用空间namespace/**
func (restHandler RestHandler) CreateGlobalizationCopyWritingNamespace(context *gin.Context) {
	namespaceRequest := &Entity.NamespaceRequest{}
	shouldBindBodyWithErr := context.ShouldBindBodyWith(&namespaceRequest, binding.JSON)
	if shouldBindBodyWithErr != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	_, searchApplicationError := model.NamespaceHandler.CreateApplicationNamespace(namespaceRequest)
	if searchApplicationError != nil {
		logger.Error(searchApplicationError)
		Response.ResErrCli(context, searchApplicationError)
		return
	}
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWritingStruct 获取多语言文案结构/**
func (restHandler RestHandler) ListGlobalizationCopyWritingStruct(context *gin.Context) {
	namespaceRequest := &Entity.NamespaceRequest{}
	shouldBindBodyWithErr := context.ShouldBindBodyWith(&namespaceRequest, binding.JSON)
	if shouldBindBodyWithErr != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	searchApplicationList, searchApplicationError := model.NamespaceHandler.ListApplicationNamespace(namespaceRequest)
	if searchApplicationError != nil {
		logger.Error(searchApplicationError)
		Response.ResFail(context, "应用处理异常")
		return
	}
	Response.ResSuccess(context, searchApplicationList)
}

// ListGlobalizationCopyWritingNamespace 查询应用文案命名空间/**
func (restHandler RestHandler) ListGlobalizationCopyWritingNamespace(context *gin.Context) {
	globalDocumentRequest := &Entity.GlobalDocumentRequest{}
	shouldBindBodyWithErr := context.ShouldBindBodyWith(&globalDocumentRequest, binding.JSON)
	if shouldBindBodyWithErr != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	searchApplicationList, searchApplicationError :=
		model.DocumentHandler.SearchDocumentByNamespaceId(
			globalDocumentRequest.ApplicationId, globalDocumentRequest.NamespaceId)
	if searchApplicationError != nil {
		logger.Error(searchApplicationError)
		Response.ResFail(context, "应用处理异常")
		return
	}
	Response.ResSuccess(context, searchApplicationList)
}

// CreateGlobalizationCopyWriting 创建多语言文案/**
func (restHandler RestHandler) CreateGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	resultId, err := model.DocumentHandler.CreateDocument(json)
	if err != nil {
		Response.ResErrCli(context, err)
		return
	}
	println(resultId)
	Response.ResSuccessMsg(context)
}

// QueryGlobalizationCopyWritingDetail 创建多语言文案/**
func (restHandler RestHandler) QueryGlobalizationCopyWritingDetail(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
		return
	}
	result, err := model.DocumentHandler.SearchDocumentById(json.Id)
	if err != nil {
		Response.ResErrCli(context, err)
		return
	}
	Response.ResSuccess(context, result)
}

// UpdateGlobalizationCopyWriting 更新多语言文案/**
func (restHandler RestHandler) UpdateGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// CommitGlobalizationCopyWriting 提交多语言文案更新/**
func (restHandler RestHandler) CommitGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWriting 创建多语言文案/**
func (restHandler RestHandler) ListGlobalizationCopyWriting(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}

// ListGlobalizationCopyWritingHistory 创建多语言文案/**
func (restHandler RestHandler) ListGlobalizationCopyWritingHistory(context *gin.Context) {
	json := &Entity.GlobalDocumentRequest{}
	err := context.ShouldBindBodyWith(&json, binding.JSON)
	if err != nil {
		Response.ResFail(context, "json解析异常")
	}
	logger.Info("aaa{}", json)
	Response.ResSuccessMsg(context)
}
