package datasource

// Create 方法执行
func Create(value interface{}) error {
	return DB.Create(value).Error
}
