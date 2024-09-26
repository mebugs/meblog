package log

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

// const uuid = "ME_BUGS_SITE_OL_STONE"

func TestLogLevel(t *testing.T) {
	setLevelIndex(InfoLevel)
	if IsLevelEnabled("debug") ||
		!IsLevelEnabled("info") ||
		!IsLevelEnabled("warn") {
		t.FailNow()
	}
	setLevelIndex(DebugLevel) // 恢复现场，避免影响其他单元测试
}

func TestChangeWriter(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 4096))
	ChangeWriter(buf)

	rand := time.Now().String()
	Info(rand)
	if !bytes.Contains(buf.Bytes(), ([]byte)(rand)) {
		t.FailNow()
	}
}

func TestChangeFormat(t *testing.T) {
	format := fmt.Sprintf(`<log><date>%s</date><time>%s</time><level>%s</level><file>%s</file><line>%d</line><msg>%s</msg><log>`,
		"2006-01-02", "15:04:05.000", LevelToken, FileToken, LineToken, MessageToken)
	ChangeFormat(format)

	buf := bytes.NewBuffer(make([]byte, 4096))
	ChangeWriter(buf)

	rand := time.Now().String()
	Debug(rand)
	if bytes.HasPrefix(buf.Bytes(), ([]byte)("<log><date>")) &&
		!bytes.HasSuffix(buf.Bytes(), ([]byte)("</msg><log>")) {
		t.FailNow()
	}
}

func TestPanicLog(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()
	Panic("test panic")
}

func TestNormalLog(t *testing.T) {
	setLevelIndex(AllLevel)

	Trace(AllLevel)
	Trace(TraceLevel)
	Debug(DebugLevel)
	Info(InfoLevel)
	Warn(WarnLevel)
	Error(ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		Panic(PanicLevel)
	}()
	// Fatal( FatalLevel)
	Print(PrintLevel)
	Stack(StackLevel)
}

func TestFormatLog(t *testing.T) {
	setLevelIndex(AllLevel)

	TraceF("%d %s", AllLevel, AllLevel)
	TraceF("%d %s", TraceLevel, TraceLevel)
	DebugF("%d %s", DebugLevel, DebugLevel)
	InfoF("%d %s", InfoLevel, InfoLevel)
	WarnF("%d %s", WarnLevel, WarnLevel)
	ErrorF("%d %s", ErrorLevel, ErrorLevel)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		PanicF("%d %s", PanicLevel, PanicLevel)
	}()
	// FatalF("%d %s", FatalLevel, FatalLevel)
	PrintF("%d %s", PrintLevel, PrintLevel)
	StackF("%d %s", StackLevel, StackLevel)
}

func TestGetLogger(t *testing.T) {
	std := GetStdLogger()
	std.Print("123")
	std.Print(321)
}

// 暂无意义的测试类
func TestRe(t *testing.T) {
	removeLogFile()
}
