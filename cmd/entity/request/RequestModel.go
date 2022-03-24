package request

type ApplicationAddRequest struct {
	// 应用名称
	ApplicationName string `json:"application_name,omitempty"`
	// 应用类型 WEB & APPLICATION
	ApplicationType string `json:"application_type,omitempty"`
	// 应用管理员
	ApplicationAdministrators int32 `json:"application_administrators,omitempty"`
	// 应用路径 默认应用路径
	ApplicationPath string `json:"application_path,omitempty"`
	// 包含的语言范围
	ApplicationLanguage []string `json:"application_language,omitempty"`
	// 应用环境
	ApplicationEnvironment string `json:"application_environment,omitempty"`
}

type DocumentAddRequest struct {
	Application string     `json:"application,omitempty"`
	Type        string     `json:"type,omitempty"`
	Path        string     `json:"path,omitempty"`
	Key         string     `json:"key,omitempty"`
	Document    []Document `json:"document,omitempty"`
}

type Document struct {
	CountryCode   string `json:"countryCode,omitempty"`
	LanguageValue string `json:"languageValue,omitempty"`
}
