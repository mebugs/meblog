# useage
日志文件默认输出到stdout，请使用重定向启动应用。
```sh
prog="stone"
logsFolder="../logs"
nohup ./$prog >> $logsFolder/$prog.log 2>&1 &
```

日志切割时，默认日志文件保存在prog/../logs/prog.log (父级目录中的logs文件夹)，如果找不到，则放弃切割。

## 配置示例
```sh
log.SetLevel(log.DebugLevel)
log.SetAutoSplit(log.CronDaily)
log.SetAppRoot(0)   
```

## 设置日志级别

log.SetLevel("debug")

设置日志级别，可选值，级别从低到高：

"all", "trace", "debug", "info", "warn", "error", "panic", "fatal", "print", "stack"

## 设置日志自动切分

log.SetAutoSplit(log.CronDaily)

根据对应的规则定时自动切分日志，定时任务参考[cron](https://godoc.org/github.com/robfig/cron)

## 设置应用根目录

log.SetAppRoot(2)

不设置应用根目录，日志会打印gopath完整路径。

如果不想打印长长的gopath路径，可以设置应用根目录。

有些应用main函数(log.SetAppRoot调用位置)不在根目录，传入parent参数，将根目录设为指定层级父目录，eg.. 0为当前目录，1为上层目录，2为上上层目录。

## 设置日志输出位置
og.SetLogPath("../logs/stone.log")

如果不想输出到stdout，可以设置输出位置

## 设置日志输出模板

log.ChangeFormat(format string)

通过组合内置的一些元素设置日志输出模板，改变日志输出格式

format := fmt.SprintF("%s %s %s %s:%d %s", "2006-01-02 15:04:05.000000", log.TagToken, log.LevelToken, log.FileToken, log.LineToken, log.DivideToken, log.MessageToken)

