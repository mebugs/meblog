package security

import (
	"fmt"
	"strings"
)

// 脱敏规范 长度 保留前缀 保留后缀
var desMap = map[int][]int{
	4: {1, 0}, 5: {2, 0}, 6: {2, 1}, 7: {2, 1}, 8: {2, 2}, 9: {3, 2}, 10: {3, 3},
	11: {3, 4}, 12: {3, 4}, 13: {4, 4}, 14: {4, 4}, 15: {5, 4}, 16: {6, 4},
}

// DesHash 脱敏+Hash
func DesHash(str string) string {
	return fmt.Sprintf("%s(%s)", Desensitization(str), SHA256(str))
}

// Desensitization 字符脱敏
func Desensitization(str string) string {
	vL := len(str)
	rA := []rune(str)
	// 字符串中存在中文
	if vL != len(rA) {
		return desensitizationWithCn(rA)
	}
	var desStr string
	// 4位以内全*
	if vL < 4 {
		desStr = strings.Repeat("*", vL)
	} else {
		desC, ok := desMap[vL]
		if !ok {
			desC = []int{6, 4}
		}
		// 处理脱敏长度
		desVal := strings.Repeat("*", vL-desC[0]-desC[1])
		desStr = fmt.Sprintf("%s%s%s", str[:desC[0]], desVal, str[vL-desC[1]:vL])
	}
	return desStr
}

// 包含中文的脱敏处理
func desensitizationWithCn(r []rune) string {
	rL := len(r)
	var desStr string
	// 4位以内全*
	if rL < 4 {
		desStr = strings.Repeat("*", rL)
	} else {
		desC, ok := desMap[rL]
		if !ok {
			desC = []int{6, 4}
		}
		// 脱敏后缀下标
		endI := rL - desC[1]
		// 处理脱敏长度
		for i := 0; i < len(r); i++ {
			if i >= desC[0] && i < endI {
				r[i] = '*'
			}
		}
		desStr = string(r)
	}
	return desStr
}
