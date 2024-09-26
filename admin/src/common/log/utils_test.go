package log

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtactDateTimeFormat(t *testing.T) {
	format := `{"level": "info", "file": "path/main.go", "line":88, "log": "message"}`
	dateFmt, timeFmt := extactDateTimeFormat(format)
	if dateFmt != "" && timeFmt != "" {
		t.FailNow()
	}

	format = `{"datetime": "2006-01-02 15:04:05.999999999", "level": "info", "file": "path/main.go", "line":88, "log": "message"}`
	dateFmt, timeFmt = extactDateTimeFormat(format)
	if dateFmt != "2006-01-02 15:04:05.999999999" && timeFmt != "" {
		t.FailNow()
	}

	format = `{"date": "2006-01-02", "time": "15:04:05.999999999", "level": "info", "file": "path/main.go", "line":88, "log": "message"}`
	dateFmt, timeFmt = extactDateTimeFormat(format)
	if dateFmt != "2006-01-02" && timeFmt != "15:04:05.999999999" {
		t.FailNow()
	}

	// 测试 日期模式不能重复出现在 format 中，不能判定是模式还是固定字符串
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("must panic, but not")
				t.FailNow()
			}
		}()

		// 有两个 2006 ，会出错
		format = `{"date": "2006-01-02", "time": "15:04:05.999999999", "Tag": "2006" "level": "info", "file": "path/main.go", "line":88, "log": "message"}`
		dateFmt, timeFmt = extactDateTimeFormat(format)
	}()
}

func TestReplace(t *testing.T) {
	str := `/a/b/c`
	b := "/a/b"

	r := strings.Replace(str, b, "", 1)
	fmt.Println(r)
}

func TestParentDir(t *testing.T) {
	str := `/a/b/c`

	r := filepath.Dir(str)
	fmt.Println(r)
}
