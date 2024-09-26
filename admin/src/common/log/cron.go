package log

import (
	"github.com/robfig/cron/v3"
)

const (
	CronSecond = "0/3 * * * * *"
	CronHour   = "0 0 * * * *"
	CronDaily  = "0 0 0 * * *"
	CronMonth  = "0 0 0 1 * *"
)

// 定时分割日志
func cronStart(rule string) {
	// CronV3在处理CronTab表达式默认处理分钟级别，如需要秒级需要增加cron.WithSeconds()
	c := cron.New(cron.WithSeconds())
	c.AddFunc(rule, logSplit) // 每隔一天
	c.Start()
}
