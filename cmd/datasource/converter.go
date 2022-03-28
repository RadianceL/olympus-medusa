// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package datasource

import (
	"database/sql"
	"strings"
)

// SetColVarType set the column type.
func SetColVarType(colVar *[]interface{}, i int, typeName string) {
	dt := DT(typeName)
	switch {
	case Contains(dt, BoolTypeList):
		var s sql.NullBool
		(*colVar)[i] = &s
	case Contains(dt, IntTypeList):
		var s sql.NullInt64
		(*colVar)[i] = &s
	case Contains(dt, FloatTypeList):
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case Contains(dt, UintTypeList):
		var s []uint8
		(*colVar)[i] = &s
	case Contains(dt, StringTypeList):
		var s sql.NullString
		(*colVar)[i] = &s
	default:
		var s interface{}
		(*colVar)[i] = &s
	}
}

// SetResultValue set the result value.
func SetResultValue(result *map[string]interface{}, index string, colVar interface{}, typeName string) {
	dt := DT(typeName)
	index = camelString(index)
	switch {
	case Contains(dt, BoolTypeList):
		temp := *(colVar.(*sql.NullBool))
		if temp.Valid {
			(*result)[index] = temp.Bool
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, IntTypeList):
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, FloatTypeList):
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, UintTypeList):
		(*result)[index] = *(colVar.(*[]uint8))
	case Contains(dt, StringTypeList):
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	default:
		if colVar2, ok := colVar.(*interface{}); ok {
			if colVar, ok = (*colVar2).(int64); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).(string); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).(float64); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).([]uint8); ok {
				(*result)[index] = colVar
			} else {
				(*result)[index] = colVar
			}
		}
	}
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
