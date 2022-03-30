package request

type ApplicationRequest struct {
	// 应用名称
	ApplicationName string `json:"applicationName,omitempty"`
	// 应用类型 WEB & APPLICATION
	ApplicationType string `json:"applicationType,omitempty"`
	// 应用管理员
	ApplicationAdministrators int `json:"applicationAdministrators,omitempty"`
	// 应用路径 默认应用路径
	ApplicationPath string `json:"applicationPath,omitempty"`
	// 包含的语言范围
	ApplicationLanguage []string `json:"applicationLanguage,omitempty"`
	// 应用环境
	ApplicationEnvironment string `json:"applicationEnvironment,omitempty"`
}

type NamespaceRequest struct {
	NamespaceId            int    `json:"namespaceId,omitempty"`
	NamespaceCode          string `json:"namespaceCode,omitempty"`
	NamespaceName          string `json:"namespaceName,omitempty"`
	NamespacePath          string `json:"namespacePath,omitempty"`
	NamespaceParentId      int    `json:"namespaceParentId,omitempty"`
	NamespaceApplicationId int    `json:"namespaceApplicationId,omitempty"`
	CreateUserId           int    `json:"createUserId,omitempty"`
}

type GlobalDocumentRequest struct {
	ApplicationId int                      `json:"applicationId,omitempty"`
	NamespaceId   int                      `json:"namespaceId,omitempty"`
	Key           string                   `json:"key,omitempty"`
	Documents     []GlobalDocumentLanguage `json:"documents,omitempty"`
}

type GlobalDocumentLanguage struct {
	CountryCode   string `json:"countryCode,omitempty"`
	LanguageValue string `json:"languageValue,omitempty"`
}
