package gen

// TableConfig 数据库配置结构体
type TableConfig struct {
	PackName   string      // 包名
	DbName     string      // 库名
	TbName     string      // 表名 = 文件名
	ObjName    string      // 结构体名 = 大写驼峰
	Router     string      // 路由名 = 小写驼峰
	Remark     string      // 表注释
	Columns    [][7]string // 二维数组分别是 0 字段名 1 类型 2 JSON名 3 注释 4 非空 5 长度 6 源类型
	TimeImport bool        // 是否需要导入Time包
}
