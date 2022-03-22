package model

import (
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
	Entity "medusa-globalization-copywriting-system/cmd/entity/request"
)

// ApplicationModel is application model structure.
type ApplicationModel struct {
	Base

	Id        int64
	Title     string
	ParentId  int64
	Icon      string
	Uri       string
	Header    string
	CreatedAt string
	UpdatedAt string
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
func (applicationModel ApplicationModel) AddApplication(applicationAddRequest Entity.ApplicationAddRequest) (int64, error) {
	return applicationModel.Table("tb_application").
		Insert(dialect.H{
			"id":                         applicationModel.Id,
			"application_name":           applicationAddRequest.ApplicationName,
			"application_administrators": 1234,
			"application_type":           "111",
			"application_path":           "111",
			"must_contain_language":      "{}",
			"application_environment":    "111",
		})
}
