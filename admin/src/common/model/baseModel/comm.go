package baseModel

// Req 空白请求对象
type Req struct {
}

// IdReq ID查询对象
type IdReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
}

// SelectRes 通用下拉响应
type SelectRes struct {
	Label string `json:"label" example:"Name"`  // 展示名
	Value string `json:"value" example:"Value"` // 展示值
}

// SelectNumRes 通用下拉数值响应
type SelectNumRes struct {
	Label string `json:"label" example:"Name"` // 展示名
	Value uint64 `json:"value" example:"1"`    // 展示值
}

// SortReq 排序对象
type SortReq struct {
	ID   uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	Sort uint16 `json:"sort" example:"1"`                  // 序号
}

// SortRes 通用排序前置响应列表
type SortRes struct {
	Id   uint64 `json:"id" example:"1"`         // 默认数据ID
	Name string `json:"name" example:"Account"` // 权限名称，界面展示，建议与界面导航一致
	Sort uint16 `json:"sort" example:"1"`       // 权限排序
}

// Tree 树对象
type Tree struct {
	Title    string  `json:"title" example:"ROOT"` // 树标题
	Key      string  `json:"key" example:"ROOT"`   // 树键
	Children []*Tree `json:"children"`             // 子树
	Level    string  `json:"level" example:"1"`    // 表示树等级
	Expand   any     `json:"expand"`               // 拓展信息
	Id       uint64  `json:"id" example:"1"`       // 表示树数据ID
}
