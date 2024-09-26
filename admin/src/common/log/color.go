package log

import (
	"fmt"
)

const (
	// 颜色
	colorRed = iota + 31
	colorGreen
	colorYellow
	colorBlue
	colorMagenta // 紫红色
	colorCyan    // 蓝绿色

	dateColor    = colorBlue
	traceIDColor = colorCyan
)

// 加颜色
func boldColorStr(color int, str string) string {
	return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", color, str)
}

func colorStr(color int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}
