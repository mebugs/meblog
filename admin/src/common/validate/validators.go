package validate

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// 字符最大长度
func maxLength(lv validator.FieldLevel) bool {
	val := lv.Field().String()
	if val == "" {
		return true
	}
	i, _ := strconv.Atoi(lv.Param())
	l := len(val)
	return l <= i
}

// checkPhoneNum 校验手机号是否合法
func checkPhoneNum(phone string) (legal bool) {
	pattern := `^[1]([3-9])[0-9]{9}$`
	legal, _ = regexp.MatchString(pattern, phone)
	return
}

// letterUnder 校验字母下划线
func letterUnder(lv validator.FieldLevel) (check bool) {
	val := lv.Field().String()
	if val == "" {
		return true
	}
	pattern := `^[a-zA-Z_]*$`
	check, _ = regexp.MatchString(pattern, val)
	return
}
