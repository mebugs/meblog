package log

import (
	syslog "log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	DEFAULT_SUFFIX = ".log"
	DEFAULT_DIR    = "../logs/"
	//Testing        = "test"
)

var (
	std         Logger
	globalLevel = DebugLevel // 默认 debug 级别

	logPath             = ""
	logFilePtr *os.File = nil
	splitRule           = ""
	appRoot             = ""
	logFormat           = DEFAULT_FORMAT_TRACEID
)

func init() {
	// 默认输出os.Std
	std = NewStandard(os.Stdout, logFormat)
	if logPath == "" {
		_, progName := filepath.Split(os.Args[0])
		logPath = DEFAULT_DIR + progName + DEFAULT_SUFFIX
	}

	redirectSyslog()
}

// 改变syslog包输出
func redirectSyslog() {
	syslog.SetOutput(std)
}

// 日志切分
func logSplit() {
	if !isFileExist(logPath) {
		return
	}

	// 新文件名时间部分
	// 日切设为前一天
	// 其他情况设为切割时间点 精确到秒防止文件被覆盖
	var splitTime string
	if splitRule == CronDaily {
		splitTime = time.Now().Add(-24 * time.Hour).Format("20060102")
	} else {
		splitTime = time.Now().Format("20060102150405")
	}

	newName := logPath + "." + splitTime
	err := os.Rename(logPath, newName) // 重命名当前log文件名
	if err != nil {
		ErrorF("move the log file error = %s", err)
		return
	}

	InfoF("logPath = %s", logPath)
	InfoF("time now: %s", time.Now())

	fileP, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		ErrorF("open the log file error = %s", err)
		return
	}

	if fileP != nil {
		oldFile := logFilePtr
		logFilePtr = fileP
		std.ChangeWriter(logFilePtr)
		InfoF("logFilePtr = %s", logFilePtr.Name())

		// 关掉之前的文件
		if oldFile != nil {
			err := oldFile.Close()
			if err != nil {
				ErrorF("close the log file %s fail, error Is :%s", oldFile.Name(), err)
			}
		}
	}
}

// 移除日志 没有使用
func removeLogFile() {
	sPath, err := os.Getwd()
	if err != nil {
		ErrorF("get the current file path error, the error Is :%s\n", err)
		return
	}

	logDir, logFileName := filepath.Split(logPath)
	sPath += "/" + logDir

	// 一天前
	d, err := time.ParseDuration("-24h")
	if err != nil {
		ErrorF("parse the time error, the error Is :%s", err)
		return
	}
	cur := time.Now()
	curDate := cur.Add(d * 60).Format("20060102")

	_ = filepath.Walk(sPath, func(path string, info os.FileInfo, err error) error {
		if (info != nil) && (!info.IsDir()) && (strings.Index(info.Name(), DEFAULT_SUFFIX) > 0) {
			nameArray := []byte(info.Name())
			var date []byte
			date = nameArray[len(logFileName)+1:]
			sArray := strings.Split(string(date), ".")
			if isMoreTwoMonth(sArray[0], curDate) { // 超过两个月删除
				err = os.Remove(sPath + info.Name())
				if err != nil {
					ErrorF("delete the file %s error, error is ", info.Name(), err)
					return err
				}
			}
		}
		return nil
	})
}

// true:超过两个月
func isMoreTwoMonth(origDate, curDate string) bool {
	if strings.Compare(origDate, curDate) <= 0 {
		return true
	}
	return false
}
