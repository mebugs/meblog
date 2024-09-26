package log

// 可以用这些串和日期、时间（包含毫秒数）任意组合，拼成各种格式的日志，如 csv/json/xml
const (
	LevelToken   string = "info"
	TagToken            = "tag"
	FileToken           = "path/main.go"
	LineToken    int    = 88
	DivideToken         = "|"
	MessageToken string = "message"
	TraceID             = "traceID"
)

// DEFAULT_FORMAT 默认日志格式
const DEFAULT_FORMAT = "2006-01-02 15:04:05 info path/main.go:88 | message"

// DEFAULT_FORMAT_TAG 默认日志格式带标签
const DEFAULT_FORMAT_TAG = "2006-01-02 15:04:05 tag info path/main.go:88 | message"

// DEFAULT_FORMAT_TRACED 默认日志格式带TraceID
const DEFAULT_FORMAT_TRACEID = "2006-01-02 15:04:05.000 traceID info path/main.go:88 | message"

var (
	DEFAULT_DIVIDE = boldColorStr(colorGreen, DivideToken)
	DEFAULT_TAG    = "-"
)
