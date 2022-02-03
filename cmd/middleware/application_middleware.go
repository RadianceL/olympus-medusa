package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"medusa-globalization-copywriting-system/cmd/web/entity/response"
	"net/http"
	"sort"
	"strings"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(405, gin.H{"message": "方法不被允许"})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(500, gin.H{"message": "未找到请求路由的处理函数"})
	}
}

func Validate(skipperFunc SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := response.Model{}
		_ = c.Request.URL.Path
		b := skipperFunc(c)
		if b {
			c.Next()
			return
		}
		// remoteAddr := c.ClientIP()  // 也可以对客户端IP进行限制
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			c.Abort()
			//Content-Type 类型只支持 application/json
			c.JSON(http.StatusUnauthorized, resp)
			return
		}
		assessKey := c.DefaultQuery("assessKey", "")
		if assessKey == "" {
			c.Abort()
			// token缺失
			c.JSON(http.StatusUnauthorized, resp)
			return
		}

		method := strings.ToUpper(c.Request.Method)

		var urlParams, _ string
		if "POST" == method || "PUT" == method {
			body, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重设body
			_ = string(body)
		} else if "GET" == method || "DELETE" == method {
			queryParams := c.Request.URL.Query()
			allParams := make(map[string]string)
			for k, v := range queryParams {
				if k != "assessKey" && k != "expires" && k != "signature" {
					// 只用第一个的值
					allParams[k] = v[0]
				}
			}
			keys := getMapKeysSorted(allParams)
			for _, k := range keys {
				urlParams += k + allParams[k]
			}
		}
		c.Next()
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := context.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := context.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := JoinRouter(context.Request.Method, context.Request.URL.Path)
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}

// 签名算法如下
/*
Signature = HMAC-SHA1('SecretKey', UTF-8-Encoding-Of( StringToSign ) ) );
StringToSign = method + "\n" +
               URL + "\n" +
               Sort-UrlParams + "\n" +
               Content-MD5 + "\n" +  // md5(params)
               Expires + "\n" +
               assessKey;
*/
func genSignature(assessKey, secretKey, uri, method, urlParams, params, nowTS string) (string, error) {
	if params != "" {
		md5Ctx := md5.New()
		_, _ = io.WriteString(md5Ctx, params)
		params = fmt.Sprintf("%x", md5Ctx.Sum(nil))
	}

	strSign := method + "\n" + uri + "\n" + urlParams + "\n" + params + "\n" + nowTS + "\n" + assessKey
	sign := hmacSHA1Encrypt(strSign, secretKey)
	return sign, nil
}

// hmacSHA1Encrypt encrypt the encryptText use encryptKey
func hmacSHA1Encrypt(encryptText, encryptKey string) string {
	key := []byte(encryptKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(encryptText))
	var str = hex.EncodeToString(mac.Sum(nil))
	return str
}

func getMapKeysSorted(originMap map[string]string) []string {
	keys := make([]string, len(originMap))
	i := 0
	for k := range originMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
