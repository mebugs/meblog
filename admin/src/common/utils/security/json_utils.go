package security

import (
	"regexp"
	"strings"
)

// SafeJson 为keys的值进行安全脱敏
func SafeJson(jsonStr string, keys []string) string {
	if jsonStr == "" || len(keys) == 0 {
		return jsonStr
	}
	// var safeJson string
	// 处理匹配
	for _, key := range keys {
		exp := regexp.MustCompile(`\"` + key + `\"\s*:\s*"(.*?)\"`)
		matchList := exp.FindAllStringSubmatch(jsonStr, -1)
		for _, match := range matchList {
			// 未匹配
			if len(match) != 2 || match[1] == "" {
				continue
			}
			// 值脱敏
			desH := DesHash(match[1])
			// 脱敏值替换和Json替换
			jsonStr = strings.ReplaceAll(jsonStr, match[0], strings.ReplaceAll(match[0], match[1], desH))
		}
	}
	return jsonStr
}
