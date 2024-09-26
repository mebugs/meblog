package log

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	//c.AddFunc("0 0 0 * * ?", logSplit) // 每隔一天
	c.AddFunc("0/1 * * * * ?", func() {
		fmt.Println("cron Table")
	}) // 每隔3秒
	// c.AddFunc("0 0 0 1 * *", removeLogFile) // 每隔一个月执行
	c.Start()

	time.Sleep(1 * time.Hour)
}
