package language

import "errors"

var CN, _ = Initialize("CN", "中国", "中文")
var JP, _ = Initialize("JP", "日本", "日文")
var EN, _ = Initialize("EN", "英国", "英文")
var RU, _ = Initialize("RU", "俄罗斯", "俄文")

type Language struct {
	CountryCode     string
	CountryName     string
	CountryLanguage string
}

func FindLanguage(countryCode string) *Language {
	switch countryCode {
	case CN.CountryCode:
		return CN
	case JP.CountryCode:
		return JP
	case EN.CountryCode:
		return EN
	case RU.CountryCode:
		return RU
	default:
		return nil
	}
}

func Initialize(countryCode string, countryName string, countryLanguage string) (lan *Language, err error) {
	if countryCode == "" || len(countryName) != 2 {
		err = errors.New("国家二字编码不能为空 || 非2个字符")
	}
	return &Language{CountryCode: countryCode, CountryName: countryName, CountryLanguage: countryLanguage}, nil
}
