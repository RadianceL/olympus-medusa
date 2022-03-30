package data

type TableApplication struct {
	Id int64 `json:"id,omitempty"`
	// 应用名称
	ApplicationName string `json:"applicationName,omitempty"`
	// 应用类型 WEB & APPLICATION
	ApplicationType string `json:"applicationType,omitempty"`
	// 应用管理员
	ApplicationAdministrators int32 `json:"applicationAdministrators,omitempty"`
	// 应用路径 默认应用路径
	ApplicationPath string `json:"applicationPath,omitempty"`
	// 包含的语言范围
	ApplicationLanguage []string `json:"applicationLanguage,omitempty"`
	// 应用环境
	ApplicationEnvironment string `json:"applicationEnvironment,omitempty"`
}

type TableApplicationNamespace struct {
	NamespaceId            int    `json:"namespaceId,omitempty"`
	NamespaceCode          string `json:"namespaceCode,omitempty"`
	NamespaceName          string `json:"namespaceName,omitempty"`
	NamespacePath          string `json:"namespacePath,omitempty"`
	NamespaceParentId      int    `json:"namespaceParentId,omitempty"`
	NamespaceApplicationId int    `json:"namespaceApplicationId,omitempty"`
	CreateUserId           int    `json:"createUserId,omitempty"`
}

type TableGlobalDocument struct {
	Id            int                           `json:"id,omitempty"`
	ApplicationId int                           `json:"applicationId,omitempty"`
	NamespaceId   int                           `json:"namespaceId,omitempty"`
	DocumentCode  string                        `json:"documentCode,omitempty"`
	Documents     []TableGlobalDocumentLanguage `json:"documents,omitempty"`
}

type TableGlobalDocumentLanguage struct {
	CountryIso    string `json:"countryIso,omitempty"`
	DocumentValue string `json:"documentValue,omitempty"`
}
