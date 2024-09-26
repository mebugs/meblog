package baseModel

// PageReq 分页查询对象
type PageReq struct {
	Current  int `json:"current" example:"1"`   // 当前页
	PageSize int `json:"pageSize" example:"10"` // 单页数量
}

// PageInit 分页初始化 默认第一页 默认10条数据
func (p *PageReq) PageInit() {
	if p.Current == 0 {
		p.Current = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}

// PageRes 分页查询响应对象
type PageRes struct {
	List  any   `json:"list"`                // 分页数据
	Total int64 `json:"total" example:"100"` // 数据总量
}

// SetPageRes 生成分页响应对象
func SetPageRes(list any, total int64) PageRes {
	return PageRes{
		List:  list,
		Total: total,
	}
}
