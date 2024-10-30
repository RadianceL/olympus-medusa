package model

import (
	"github.com/mitchellh/mapstructure"
	"olympus-medusa/cmd/datasource"
	"olympus-medusa/cmd/datasource/data"
	"olympus-medusa/cmd/datasource/dialect"
	Entity "olympus-medusa/cmd/entity/request"
	"olympus-medusa/cmd/handler/model/basic"
)

const (
	// NamespaceModelTableName tb_application_namespace
	namespaceModelTableName = "tb_application_namespace"
	namespaceCode           = "namespace_code"
	namespaceName           = "namespace_name"
	namespacePath           = "namespace_path"
	namespaceParentId       = "namespace_parent_id"
	namespaceApplicationId  = "application_id"
	createUser              = "create_user"
)

// NamespaceModel is application model structure.
type NamespaceModel struct {
	basic.BaseModel
}

func init() {
	RepositoryModelContainer.Register(Namespace())
}

func (namespaceModel NamespaceModel) Initialization() {
	NamespaceHandler = Namespace().SetConn(datasource.Conn)
	println("初始化NamespaceHandler")
}

// Namespace return a default menu model.
func Namespace() NamespaceModel {
	return NamespaceModel{BaseModel: basic.BaseModel{TableName: namespaceModelTableName}}
}

func (namespaceModel NamespaceModel) SetConn(connection datasource.Connection) NamespaceModel {
	namespaceModel.Conn = connection
	return namespaceModel
}

func (namespaceModel NamespaceModel) CreateApplicationNamespace(namespaceRequest *Entity.NamespaceRequest) (int64, error) {
	return namespaceModel.Table(namespaceModelTableName).
		Insert(dialect.H{
			namespaceCode:          namespaceRequest.NamespaceCode,
			namespaceName:          namespaceRequest.NamespaceName,
			namespacePath:          namespaceRequest.NamespacePath,
			namespaceParentId:      namespaceRequest.NamespaceParentId,
			namespaceApplicationId: namespaceRequest.NamespaceApplicationId,
			createUser:             namespaceRequest.CreateUserId,
		})
}

func (namespaceModel NamespaceModel) ListApplicationNamespace(namespaceRequest *Entity.NamespaceRequest) ([]data.TableApplicationNamespace, error) {
	statement := namespaceModel.Table(namespaceModelTableName).Select("*")
	if namespaceRequest.NamespaceCode != "" {
		statement.Where(namespaceCode, "=", namespaceRequest.NamespaceCode)
	}
	if namespaceRequest.NamespaceName != "" {
		statement.Where(namespaceName, "LIKE", "%"+namespaceRequest.NamespaceName+"%")
	}
	resultData, err := statement.All()
	if err != nil {
		return []data.TableApplicationNamespace{}, err
	}
	var result []data.TableApplicationNamespace
	for _, value := range resultData {
		var outputResult data.TableApplicationNamespace
		_ = mapstructure.Decode(value, &outputResult)
		result = append(result, outputResult)
	}
	if result == nil {
		return []data.TableApplicationNamespace{}, err
	}
	return result, nil
}
