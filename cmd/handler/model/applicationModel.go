package model

import (
	"encoding/json"
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
	"medusa-globalization-copywriting-system/tools/logger"
)

// ApplicationModel is application model structure.
type ApplicationModel struct {
	Base

	Id int64
	// 应用名称
	ApplicationName string
	// 应用类型 WEB & APPLICATION
	ApplicationType string
	// 应用管理员
	ApplicationAdministrators int32
	// 应用路径 默认应用路径
	ApplicationPath string
	// 包含的语言范围
	ApplicationLanguage string
	// 应用环境
	ApplicationEnvironment string
}

// Application return a default menu model.
func Application() ApplicationModel {
	return ApplicationModel{Base: Base{TableName: "tb_application"}}
}

func (applicationModel ApplicationModel) SetConn(connection datasource.Connection) ApplicationModel {
	applicationModel.Conn = connection
	return applicationModel
}

// AddApplication add a role to the menu.
func (applicationModel ApplicationModel) AddApplication(applicationAddRequest *Entity.ApplicationAddRequest) (int64, error) {
	containLanguageList, err := json.Marshal(applicationAddRequest.ApplicationLanguage)
	if err != nil {
		logger.Panic(err)
	}
	return applicationModel.Table("tb_application").
		Insert(dialect.H{
			"id":                         applicationModel.Id,
			"application_name":           applicationAddRequest.ApplicationName,
			"application_administrators": applicationAddRequest.ApplicationAdministrators,
			"application_type":           applicationAddRequest.ApplicationType,
			"application_path":           applicationAddRequest.ApplicationName,
			"must_contain_language":      string(containLanguageList),
			"application_environment":    applicationAddRequest.ApplicationEnvironment,
		})
}

func (applicationModel ApplicationModel) SearchApplicationList(applicationAddRequest *Entity.ApplicationAddRequest) (ApplicationModel, error) {
	return nil, nil
}
