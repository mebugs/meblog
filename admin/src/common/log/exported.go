package log

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetStdLogger 返回Logger供其他包调用
func GetStdLogger() Logger {
	return std
}

// SetLogger 切换Logger实现
func SetLogger(l Logger) {
	std = l
	redirectSyslog()
}

// SetLevel 设置日志级别
func SetLevel(l string) {
	globalLevel = indexOfLevel(l)
}

// SetOutputPath 设置日志输出位置
func SetOutputPath(path string) {
	logPath = path

	// 以只写追加模式打开日志文件，不存在则创建
	fileP, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		ErrorF("open the log file error = %s", err)
		return
	}

	if fileP != nil {
		logFilePtr = fileP
		std.ChangeWriter(logFilePtr)
		InfoF("logFilePtr = %s", logFilePtr.Name())
	}
}

// SetAutoSplit 设置自动切分日志
func SetAutoSplit(rule string) {
	splitRule = rule
	cronStart(rule)
}

// SetAppRoot 设置应用根目录
func SetAppRoot(parent int) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	appRoot = filepath.Dir(file)

	// 取父级目录
	for parent > 0 {
		appRoot = filepath.Dir(appRoot)
		parent--
	}

	// 补上结尾/ 日志打印相对路径
	appRoot = strings.Replace(appRoot, "\\", "/", -1) + "/"
}

// ChangeWriter 改变输出位置，通过这个接口，可以实现日志文件按时间或按大小滚动
func ChangeWriter(w io.Writer) { std.ChangeWriter(w) }

// ChangeFormat 改变日志格式
func ChangeFormat(format string) {
	logFormat = format
	std.ChangeFormat(logFormat)
}

// IsLevelEnabled 判断某级别的日志是否会被输出
func IsLevelEnabled(level string) bool {
	return indexOfLevel(level) >= globalLevel
}

// 打印日志

func Trace(m ...any) {
	std.TPrintF(globalLevel, TraceLevel, "", "", m...)
}

func Debug(m ...any) {
	std.TPrintF(globalLevel, DebugLevel, "", "", m...)
}

func Info(m ...any) {
	std.TPrintF(globalLevel, InfoLevel, "", "", m...)
}

func Warn(m ...any) {
	std.TPrintF(globalLevel, WarnLevel, "", "", m...)
}

func Error(m ...any) {
	std.TPrintF(globalLevel, ErrorLevel, "", "", m...)
}

func Panic(m ...any) {
	std.TPrintF(globalLevel, PanicLevel, "", "", m...)
}

func Fatal(m ...any) {
	std.TPrintF(globalLevel, FatalLevel, "", "", m...)
}

func Print(m ...any) {
	std.TPrintF(globalLevel, PrintLevel, "", "", m...)
}

func Stack(m ...any) {
	std.TPrintF(globalLevel, StackLevel, "", "", m...)
}

// 按一定格式打印日志

func TraceF(format string, m ...any) {
	std.TPrintF(globalLevel, TraceLevel, "", format, m...)
}

func DebugF(format string, m ...any) {
	std.TPrintF(globalLevel, DebugLevel, "", format, m...)
}

func InfoF(format string, m ...any) {
	std.TPrintF(globalLevel, InfoLevel, "", format, m...)
}
func WarnF(format string, m ...any) {
	std.TPrintF(globalLevel, WarnLevel, "", format, m...)
}

func ErrorF(format string, m ...any) {
	std.TPrintF(globalLevel, ErrorLevel, "", format, m...)
}

func PanicF(format string, m ...any) {
	std.TPrintF(globalLevel, PanicLevel, "", format, m...)
}

func FatalF(format string, m ...any) {
	std.TPrintF(globalLevel, FatalLevel, "", format, m...)
}

func PrintF(format string, m ...any) {
	std.TPrintF(globalLevel, PrintLevel, "", format, m...)
}

func StackF(format string, m ...any) {
	std.TPrintF(globalLevel, StackLevel, "", format, m...)
}

// 打印日志时带上 tag

func TTrace(tag string, m ...any) {
	std.TPrintF(globalLevel, TraceLevel, tag, "", m...)
}

func TDebug(tag string, m ...any) {
	std.TPrintF(globalLevel, DebugLevel, tag, "", m...)
}

func TInfo(tag string, m ...any) {
	std.TPrintF(globalLevel, InfoLevel, tag, "", m...)
}

func TWarn(tag string, m ...any) {
	std.TPrintF(globalLevel, WarnLevel, tag, "", m...)
}

func TError(tag string, m ...any) {
	std.TPrintF(globalLevel, ErrorLevel, tag, "", m...)
}

func TPanic(tag string, m ...any) {
	std.TPrintF(globalLevel, PanicLevel, tag, "", m...)
}

func TFatal(tag string, m ...any) {
	std.TPrintF(globalLevel, FatalLevel, tag, "", m...)
}

func TPrint(tag string, m ...any) {
	std.TPrintF(globalLevel, PrintLevel, tag, "", m...)
}

func TStack(tag string, m ...any) {
	std.TPrintF(globalLevel, StackLevel, tag, "", m...)
}

// 按一定格式打印日志，并在打印日志时带上 tag

func TTraceF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, TraceLevel, tag, format, m...)
}

func TDebugF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, DebugLevel, tag, format, m...)
}

func TInfoF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, InfoLevel, tag, format, m...)
}

func TWarnF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, WarnLevel, tag, format, m...)
}

func TErrorF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, ErrorLevel, tag, format, m...)
}

func TPanicF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, PanicLevel, tag, format, m...)
}

func TFatalF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, FatalLevel, tag, format, m...)
}

func TPrintF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, PrintLevel, tag, format, m...)
}

func TStackF(tag string, format string, m ...any) {
	std.TPrintF(globalLevel, StackLevel, tag, format, m...)
}

// 打印日志

func TraceT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, TraceLevel, "", traceID, "", m...)
}

func DebugT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, DebugLevel, "", traceID, "", m...)
}

func InfoT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, InfoLevel, "", traceID, "", m...)
}

func WarnT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, WarnLevel, "", traceID, "", m...)
}

func ErrorT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, ErrorLevel, "", traceID, "", m...)
}

func PanicT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, PanicLevel, "", traceID, "", m...)
}

func FatalT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, FatalLevel, "", traceID, "", m...)
}

func PrintT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, PrintLevel, "", traceID, "", m...)
}

func StackT(traceID string, m ...any) {
	std.TPrintTF(globalLevel, StackLevel, "", traceID, "", m...)
}

// 按一定格式打印日志

func TraceTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, TraceLevel, "", traceID, format, m...)
}

func DebugTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, DebugLevel, "", traceID, format, m...)
}

func InfoTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, InfoLevel, "", traceID, format, m...)
}
func WarnTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, WarnLevel, "", traceID, format, m...)
}

func ErrorTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, ErrorLevel, "", traceID, format, m...)
}

func PanicTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, PanicLevel, "", traceID, format, m...)
}

func FatalTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, FatalLevel, "", traceID, format, m...)
}

func PrintTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, PrintLevel, "", traceID, format, m...)
}

func StackTF(traceID, format string, m ...any) {
	std.TPrintTF(globalLevel, StackLevel, "", traceID, format, m...)
}

// 打印日志时带上 tag

func TTraceT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, TraceLevel, tag, traceID, "", m...)
}

func TDebugT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, DebugLevel, tag, traceID, "", m...)
}

func TInfoT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, InfoLevel, tag, traceID, "", m...)
}

func TWarnT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, WarnLevel, tag, traceID, "", m...)
}

func TErrorT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, ErrorLevel, tag, traceID, "", m...)
}

func TPanicT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, PanicLevel, tag, traceID, "", m...)
}

func TFatalT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, FatalLevel, tag, traceID, "", m...)
}

func TPrintT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, PrintLevel, tag, traceID, "", m...)
}

func TStackT(traceID, tag string, m ...any) {
	std.TPrintTF(globalLevel, StackLevel, tag, traceID, "", m...)
}

// 按一定格式打印日志，并在打印日志时带上 tag

func TTraceTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, TraceLevel, tag, traceID, format, m...)
}

func TDebugTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, DebugLevel, tag, traceID, format, m...)
}

func TInfoTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, InfoLevel, tag, traceID, format, m...)
}

func TWarnTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, WarnLevel, tag, traceID, format, m...)
}

func TErrorTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, ErrorLevel, tag, traceID, format, m...)
}

func TPanicTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, PanicLevel, tag, traceID, format, m...)
}

func TFatalTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, FatalLevel, tag, traceID, format, m...)
}

func TPrintTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, PrintLevel, tag, traceID, format, m...)
}

func TStackTF(traceID, tag string, format string, m ...any) {
	std.TPrintTF(globalLevel, StackLevel, tag, traceID, format, m...)
}
