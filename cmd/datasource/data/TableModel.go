package data

type TbApplication struct {
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
	ApplicationLanguage []string
	// 应用环境
	ApplicationEnvironment string
}
