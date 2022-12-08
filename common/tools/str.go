package tools

import (
	"regexp"
	"unicode"
)

func StringToStr(t string) *string {
	return &t
}

//IsChinese 是否存在中文
func IsChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

//MustWordOrNum 是否是英文和数字 true 合法
func MustWordOrNum(str string) bool {
	re, err := regexp.Compile(`^[a-zA-Z0-9]`)
	if err != nil {
		return false
	}
	return re.MatchString(str)
}
