package model

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/data"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
	"medusa-globalization-copywriting-system/cmd/handler/model/basic"
	"medusa-globalization-copywriting-system/tools/logger"
)

const (
	// ApplicationModelTableName tb_application
	applicationModelTableName = "tb_application"
	id                        = "id"
	applicationName           = "application_name"
	applicationAdministrators = "application_administrators"
	applicationType           = "application_type"
	applicationPath           = "application_path"
	mustContainLanguage       = "must_contain_language"
	applicationEnvironment    = "application_environment"
)

// ApplicationModel is application model structure.
type ApplicationModel struct {
	basic.BaseModel
}

func init() {
	RepositoryModelContainer.Register(Application())
}

func (applicationModel ApplicationModel) Initialization() {
	ApplicationHandler = Application().SetConn(datasource.Conn)
	println("初始化ApplicationHandler")
}

// Application return a default application model.
func Application() ApplicationModel {
	return ApplicationModel{BaseModel: basic.BaseModel{TableName: applicationModelTableName}}
}

func (applicationModel ApplicationModel) SetConn(connection datasource.Connection) ApplicationModel {
	applicationModel.Conn = connection
	return applicationModel
}

// AddApplication add a role to the menu.
func (applicationModel ApplicationModel) AddApplication(applicationAddRequest *Entity.ApplicationRequest) (int64, error) {
	containLanguageList, err := json.Marshal(applicationAddRequest.ApplicationLanguage)
	if err != nil {
		logger.Panic(err)
	}
	if applicationAddRequest.ApplicationPath == "" {
		applicationAddRequest.ApplicationPath = "/" + applicationAddRequest.ApplicationName
	}
	return applicationModel.Table(applicationModelTableName).
		Insert(dialect.H{
			applicationName:           applicationAddRequest.ApplicationName,
			applicationAdministrators: applicationAddRequest.ApplicationAdministrators,
			applicationType:           applicationAddRequest.ApplicationType,
			applicationPath:           applicationAddRequest.ApplicationPath,
			mustContainLanguage:       string(containLanguageList),
			applicationEnvironment:    applicationAddRequest.ApplicationEnvironment,
		})
}

func (applicationModel ApplicationModel) SearchApplicationList(applicationAddRequest *Entity.ApplicationRequest) ([]data.TableApplication, error) {
	statement := applicationModel.Table(applicationModelTableName).Select("*")
	if applicationAddRequest.ApplicationName != "" {
		statement.Where(applicationEnvironment, "=", applicationAddRequest.ApplicationName)
	}
	resultData, err := statement.All()
	if err != nil {
		return []data.TableApplication{}, err
	}
	var result []data.TableApplication
	for _, value := range resultData {
		var outputResult data.TableApplication
		mapstructure.Decode(value, &outputResult)
		result = append(result, outputResult)
	}
	if result == nil {
		return []data.TableApplication{}, err
	}
	return result, nil
}

func (applicationModel ApplicationModel) SearchApplicationById(applicationId int) (data.TableApplication, error) {
	statement := applicationModel.Table(applicationModelTableName).Select("*")
	statement.Where(id, "=", applicationId)
	resultData, err := statement.All()
	if err != nil {
		return data.TableApplication{}, err
	}
	var outputResult data.TableApplication
	mapstructure.Decode(resultData, &outputResult)
	return outputResult, nil
}
