package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"olympus-medusa/tools/convert"
)

func GetQueryToStrE(c *gin.Context, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", errors.New("没有这个值传入")
	}
	return str, nil
}

func GetQueryToStr(c *gin.Context, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		return defaultValue
	}
	return str
}

func GetQueryToUintE(c *gin.Context, key string) (uint, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUintE(str)
}

func GetQueryToUint(c *gin.Context, key string, defaultValues ...uint) uint {
	var defaultValue uint
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUintE(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

func GetQueryToUint64E(c *gin.Context, key string) (uint64, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUint64E(str)
}

func GetQueryToUint64(c *gin.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUint64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}
