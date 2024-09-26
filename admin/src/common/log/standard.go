package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type record struct {
	TraceID    string
	Date, Time string
	Tag        string
	Level      string
	File       string
	Line       int
	Divide     string
	Message    string
	Stack      []byte
}

// Standard 日志输出基本实现
type Standard struct {
	mu  sync.Mutex // ensures atomic writes; protects the following fields
	out io.Writer  // destination for output

	format    string // log format
	pattern   string // log template
	tpl       *template.Template
	prefixLen int
	dateFmt   string
	timeFmt   string
}

// NewStandard 返回标准实现
func NewStandard(out io.Writer, format string) *Standard {
	s := &Standard{out: out}
	s.ChangeFormat(format)
	return s
}

func (s *Standard) Write(p []byte) (n int, err error) {
	return s.out.Write(p)
}

// TPrintF 格式化打印日志
func (s *Standard) TPrintF(v, l levelIndex, tag string, format string, m ...any) {
	if v > l {
		return
	}

	if tag == "" {
		tag = DEFAULT_TAG
	}
	r := record{
		Level:  l.ColorLevel(),
		Tag:    tag,
		Divide: DEFAULT_DIVIDE,
	}
	if format == "" {
		r.Message = fmt.Sprint(m...)
	} else {
		r.Message = fmt.Sprintf(format, m...)
	}
	r.Message = strings.TrimRight(r.Message, "\n")

	if s.dateFmt != "" {
		now := time.Now() // get this early.
		r.Date = boldColorStr(dateColor, now.Format(s.dateFmt))
		if s.timeFmt != "" {
			r.Time = boldColorStr(dateColor, now.Format(s.timeFmt))
		}
	}

	_, file, line, ok := runtime.Caller(2) // 跳过两层调用
	if ok {
		r.File = file
		r.Line = line
	} else {
		// r.File = "???"
		_, r.File = filepath.Split(r.File)
	}

	// 测试和研发环境（Test）打印完整文件路径
	//if appRoot != "" && logEnv != Testing {
	r.File = strings.Replace(r.File, appRoot, "", 1)
	//}

	var buf []byte
	if l == StackLevel {
		buf = make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		buf = buf[:n]
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()

		if l == PanicLevel {
			panic(m)
		}

		if l == FatalLevel {
			os.Exit(1)
		}
	}()

	_ = s.tpl.Execute(s.out, r)
	_, _ = s.out.Write([]byte("\n"))

	if l == StackLevel {
		_, _ = s.out.Write(buf)
	}
}

// TPrintTF 格式化打印日志
func (s *Standard) TPrintTF(v, l levelIndex, tag string, traceID string, format string, m ...any) {
	if v > l {
		return
	}

	if tag == "" {
		tag = DEFAULT_TAG
	}
	r := record{
		Level:  l.ColorLevel(),
		Tag:    tag,
		Divide: DEFAULT_DIVIDE,
	}
	r.TraceID = boldColorStr(traceIDColor, traceID)

	if format == "" {
		r.Message = fmt.Sprint(m...)
	} else {
		r.Message = fmt.Sprintf(format, m...)
	}
	r.Message = strings.TrimRight(r.Message, "\n")

	if s.dateFmt != "" {
		now := time.Now() // get this early.
		r.Date = boldColorStr(dateColor, now.Format(s.dateFmt))
		if s.timeFmt != "" {
			r.Time = boldColorStr(dateColor, now.Format(s.timeFmt))
		}
	}

	_, file, line, ok := runtime.Caller(2) // 跳过两层调用
	if ok {
		r.File = file
		r.Line = line
	} else {
		// r.File = "???"
		_, r.File = filepath.Split(r.File)
	}
	// 测试和研发环境（Test）打印完整文件路径
	//if appRoot != "" && logEnv != Testing {
	r.File = strings.Replace(r.File, appRoot, "", 1)
	//}

	var buf []byte
	if l == StackLevel {
		buf = make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		buf = buf[:n]
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()

		if l == PanicLevel {
			panic(m)
		}

		if l == FatalLevel {
			os.Exit(1)
		}
	}()

	_ = s.tpl.Execute(s.out, r)
	_, _ = s.out.Write([]byte("\n"))

	if l == StackLevel {
		_, _ = s.out.Write(buf)
	}
}

// ChangeWriter 改变输出流
func (s *Standard) ChangeWriter(w io.Writer) {
	s.mu.Lock()
	s.out = w
	s.mu.Unlock()
}

// ChangeFormat 改变日志格式
func (s *Standard) ChangeFormat(format string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// println(format)
	s.format = format

	s.pattern = format

	s.pattern = strings.Replace(s.pattern, FileToken, "{{ .File }}", -1)
	s.pattern = strings.Replace(s.pattern, TagToken, "{{ .Tag }}", -1)
	s.pattern = strings.Replace(s.pattern, LevelToken, "{{ .Level }}", -1)
	s.pattern = strings.Replace(s.pattern, strconv.Itoa(LineToken), "{{ .Line }}", -1)
	s.pattern = strings.Replace(s.pattern, DivideToken, "{{ .Divide }}", -1)
	s.pattern = strings.Replace(s.pattern, MessageToken, "{{ .Message }}", -1)
	s.pattern = strings.Replace(s.pattern, TraceID, "{{ .TraceID }}", -1)

	// println(s.dateFmt, s.timeFmt)

	// 提取出日期和时间的格式化模式字符串
	s.dateFmt, s.timeFmt = extactDateTimeFormat(format)
	if s.dateFmt != "" {
		s.pattern = strings.Replace(s.pattern, s.dateFmt, "{{ .Date }}", -1)
	}
	if s.timeFmt != "" {
		s.pattern = strings.Replace(s.pattern, s.timeFmt, "{{ .Time }}", -1)
	}

	s.tpl = template.Must(template.New("record").Parse(s.pattern))
}
