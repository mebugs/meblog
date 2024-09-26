package validate

func returnTrans(name string, loc string) string {
	switch name {
	case "letterUnder":
		return letterUnderTrans(loc)
	}
	return ""
}

// letterUnderTrans 字母下划线翻译注册
func letterUnderTrans(loc string) string {
	switch loc {
	case "en":
		return "{0} only support letters and underscores"
	default:
		return "{0}仅支持字母和下划线"
	}
}
