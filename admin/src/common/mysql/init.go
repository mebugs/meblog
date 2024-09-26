package mysql

import "siteol.com/smart/src/common/mysql/platDB"

// Init 初始化全部数据库
func Init(traceId string) {
	platDB.InitPlatFromDb(traceId)
}
