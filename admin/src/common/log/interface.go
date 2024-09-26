package log

import (
	"io"
)

// Logger 定义Logger接口
type Logger interface {
	io.Writer

	Printer

	Debug(v ...any)
	DebugF(format string, v ...any)

	Info(v ...any)
	InfoF(format string, v ...any)

	Warn(v ...any)
	WarnF(format string, v ...any)

	Error(v ...any)
	ErrorF(format string, v ...any)

	Fatal(v ...any)
	FatalF(format string, v ...any)

	Panic(v ...any)
	PanicF(format string, v ...any)

	Print(v ...any)
	PrintF(format string, v ...any)

	DebugT(traceID string, v ...any)
	DebugTF(traceID string, format string, v ...any)

	InfoT(traceID string, v ...any)
	InfoTF(traceID string, format string, v ...any)

	WarnT(traceID string, v ...any)
	WarnTF(traceID string, format string, v ...any)

	ErrorT(traceID string, v ...any)
	ErrorTF(traceID string, format string, v ...any)

	FatalT(traceID string, v ...any)
	FatalTF(traceID string, format string, v ...any)

	PanicT(traceID string, v ...any)
	PanicTF(traceID string, format string, v ...any)

	PrintT(traceID string, v ...any)
	PrintTF(traceID string, format string, v ...any)
}

type Printer interface {
	// TPrintF 所有方法最终归为这个方法，真正打印日志
	TPrintF(v, l levelIndex, tag string, format string, m ...any)

	// TPrintTF 所有方法最终归为这个方法，真正打印日志, 带traceID
	TPrintTF(v, l levelIndex, tag string, traceID string, format string, m ...any)

	// ChangeFormat 改变日志格式
	ChangeFormat(format string)

	// ChangeWriter 改变输出流
	ChangeWriter(w io.Writer)
}
