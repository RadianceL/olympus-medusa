package model

import (
	"medusa-globalization-copywriting-system/cmd/datasource"
	"medusa-globalization-copywriting-system/cmd/datasource/dialect"
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
	return ApplicationModel{Base: Base{TableName: "email"}}
}

func (t ApplicationModel) SetConn(con datasource.Connection) ApplicationModel {
	t.Conn = con
	return t
}

// AddApplication add a role to the menu.
func (t ApplicationModel) AddApplication(roleId string) (int64, error) {
	if roleId != "" {
		return t.Table("email").
			Insert(dialect.H{
				"id":         t.Id,
				"user_id":    "111",
				"email":      "111",
				"subscribed": "111",
			})
	}
	return 0, nil
}
