package utils

import "strings"

// ArrayInclude 字符串是否位于数组中
func ArrayInclude(node string, nodes []string) bool {
	if len(nodes) == 0 {
		return false
	}
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

// ArrayStartWith 字符串是否以某开头
func ArrayStartWith(node string, nodes []string) bool {
	if len(nodes) == 0 {
		return false
	}
	for _, n := range nodes {
		if strings.HasPrefix(node, n) {
			return true
		}
	}
	return false
}
