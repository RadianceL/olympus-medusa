package model

// ApplicationHandler 应用Model的数据库操作句柄
var ApplicationHandler ApplicationModel

// NamespaceHandler 命名空间Model的数据库操作句柄
var NamespaceHandler NamespaceModel

type RepositoryInitializationInterface interface {
	Initialization()
}

// RepositoryModelContainer RepositoryInitialization接口的实例
var RepositoryModelContainer = RepositoryInitialization{}

// repositoryInitializationList 初始化数据库仓库列表
var repositoryInitializationList []RepositoryInitializationInterface

type RepositoryInitialization struct {
}

// Register 注册数据库对象
func (r RepositoryInitialization) Register(repositoryInitializationInterface RepositoryInitializationInterface) {
	repositoryInitializationList = append(repositoryInitializationList, repositoryInitializationInterface)
}

// InitializationAll 初始化所有数据库对象
func (r RepositoryInitialization) InitializationAll() {
	for i := range repositoryInitializationList {
		repositoryInitializationList[i].Initialization()
	}
}
