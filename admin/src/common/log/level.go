package log

type levelIndex int

// 所有日志级别常量，级别越高，日志越重要，级别越低，日志越详细
const (
	AllLevel levelIndex = iota // 等同于 TraceLevel

	TraceLevel
	DebugLevel // 默认日志级别，方便开发
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel // panic 打印错误栈，但是可以 recover
	FatalLevel // fatal 表明严重错误，程序直接退出，慎用

	// 提供这个级别日志，方便在无论何种情况下，都打印必要信息，比如服务启动信息

	PrintLevel
	StackLevel // 打印堆栈信息
)

var (
	// Labels 每个级别对应的标签
	Labels = []string{"all", "trace", "debug", "info", "warn", "error", "panic", "fatal", "print", "stack"}

	colorLabels = make(map[levelIndex]string)
)

func init() {
	// 将colorLabel缓存进map
	for k := range Labels {
		lv := levelIndex(k)
		colorLabels[lv] = colorLevel(lv)
	}
}

func (v levelIndex) String() string {
	return Labels[v]
}

// ColorLevel 带颜色日志标签
func (v levelIndex) ColorLevel() string {
	if r, ok := colorLabels[v]; ok {
		return r
	} else {
		return ""
	}
}

func colorLevel(v levelIndex) string {
	color := 0

	switch v {
	case DebugLevel:
		color = colorCyan
	case InfoLevel:
		color = colorGreen
	case WarnLevel:
		color = colorYellow
	case ErrorLevel:
		color = colorRed
	}

	// 短标签补空格对齐
	str := fixLen(v.String())

	if color != 0 {
		str = boldColorStr(color, str)
	}

	return str
}

func fixLen(str string) string {
	l := len(str)
	for l < 5 {
		str += " "
		l++
	}

	return str
}

// 字符串转换成 levelIndex, "debug" => DebugLevel
func indexOfLevel(lv string) levelIndex {
	for i, v := range Labels {
		if lv == v {
			return levelIndex(i)
		}
	}

	ErrorF("level = %v not exist, use default level = %v", lv, globalLevel)
	return globalLevel
}

func setLevelIndex(i levelIndex) {
	globalLevel = i
}
