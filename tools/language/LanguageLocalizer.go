package language

import "errors"

var CN, _ = initialize("CH", "中文")
var JP, _ = initialize("JP", "日文")
var EN, _ = initialize("EN", "英文")

type Language struct {
	CountryCode string
	CountryName string
}

func findLanguage(countryCode string) *Language {
	switch countryCode {
	case CN.CountryCode:
		return CN
	case JP.CountryCode:
		return JP
	case EN.CountryCode:
		return EN
	default:
		return nil
	}
}

func initialize(countryCode string, countryName string) (lan *Language, err error) {
	if countryCode == "" || len(countryName) != 2 {
		err = errors.New("国家二字编码不能为空 || 非2个字符")
	}
	return &Language{CountryCode: countryCode, CountryName: countryName}, nil
}
