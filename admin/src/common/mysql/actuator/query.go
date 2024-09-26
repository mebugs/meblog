package actuator

import (
	"bytes"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

/**
 *
 * MYSQL 查询SQL生成器
 * 调用方式1：InitQuery(tabelName).Like("name", req.Name).Eq("service_code", req.ServiceCode).Desc("id").LimitByPage(req.Current, req.PageSize)
 * 调用方式2：TODO 反射处理，传入InterFace，需要指定字段名和操作方式（推荐简单的查询结构使用）
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-07-06
 */

// MySQL 查询器
// 整型字段即使传递”也会走索引，因此这里直接通过固定 ”作为值的模板
const (
	In        = "`${name}` in ?"    // 包含
	Eq        = "`${name}` = ?"     // 等于
	Ne        = "`${name}` != ?"    // 不等于
	Gt        = "`${name}` > ?"     // 大于
	Lt        = "`${name}` < ?"     // 小于
	Ge        = "`${name}` >= ?"    // 大于等于
	Le        = "`${name}` <= ?"    // 小于等于
	Like      = "`${name}` like ?"  // 模糊检索
	StartWith = "`${name}` like ?"  // 起始于
	EndWith   = "`${name}` like ?"  // 结束于
	BaseSql   = "SELECT %s FROM %s" // 基础Sql
	CountSql  = "count(id)"         // 数据求和字句
)

// Query 分页查询器对象
type Query struct {
	Table      string   // 表名
	Where      []*Where // 查询对象
	Order      []*Order // 排序对象
	LimitRange *Limit   // 范围过滤
}

// Where 查询结构
type Where struct {
	Name     string // 数据库字段
	Value    any    // 值
	SubWhere string // 子句模板
	//Or       []*Where // TODO 尽量减少Or查询未来实装
}

// Order 排序结构
type Order struct {
	Name string // 数据库字段
	Desc bool   // 降序？
}

// Limit 排序结构
type Limit struct {
	Offset int // 跳过数据
	Limit  int // 数据量
}

// InitQuery 初始化空查询
func InitQuery() *Query {
	return &Query{}
}

// In 包含
func (q *Query) In(name string, value any) {
	q.makeWhere(In, name, value)
}

// Eq 等于
func (q *Query) Eq(name string, value any) {
	q.makeWhere(Eq, name, value)
}

// Ne 不等于
func (q *Query) Ne(name string, value any) {
	q.makeWhere(Ne, name, value)
}

// Gt 大于
func (q *Query) Gt(name string, value any) {
	q.makeWhere(Gt, name, value)
}

// Lt 小于
func (q *Query) Lt(name string, value any) {
	q.makeWhere(Lt, name, value)
}

// Ge 大于等于
func (q *Query) Ge(name string, value any) {
	q.makeWhere(Ge, name, value)
}

// Le 小于等于
func (q *Query) Le(name string, value any) {
	q.makeWhere(Le, name, value)
}

// Like 模糊查询
func (q *Query) Like(name string, value any) {
	q.makeWhere(Like, name, fmt.Sprintf("%%%v%%", value))
}

// StartWith 开始模糊
func (q *Query) StartWith(name string, value any) {
	q.makeWhere(StartWith, name, fmt.Sprintf("%v%%", value))
}

// EndWith 结束模糊
func (q *Query) EndWith(name string, value any) {
	q.makeWhere(EndWith, name, fmt.Sprintf("%%%v", value))
}

// 实装Where
func (q *Query) makeWhere(subWhere, name string, value any) {
	if len(q.Where) == 0 {
		q.Where = make([]*Where, 0)
	}
	thisWhere := &Where{Name: name, Value: value, SubWhere: subWhere}
	q.Where = append(q.Where, thisWhere)
}

// Asc 升序
func (q *Query) Asc(name string) {
	q.makeOrder(false, name)
}

// Desc 降序
func (q *Query) Desc(name string) {
	q.makeOrder(true, name)
}

// 实装Order
func (q *Query) makeOrder(desc bool, name string) {
	if len(q.Order) == 0 {
		q.Order = make([]*Order, 0)
	}
	q.Order = append(q.Order, &Order{
		Name: name,
		Desc: desc,
	})
}

// Limit 约定范围
func (q *Query) Limit(offset, limit int) {
	q.LimitRange = &Limit{
		Offset: offset,
		Limit:  limit,
	}
}

// LimitByPage 计算范围
func (q *Query) LimitByPage(current, pageSize int) {
	offset := 0
	// 第一页无需修改offset
	if current > 1 && pageSize != 0 {
		offset = (current - 1) * pageSize
	}
	q.LimitRange = &Limit{
		Offset: offset,
		Limit:  pageSize,
	}
}

// makeSqlByQuery 根据Query对象生成SQL（包内方法）
func (q *Query) makeSqlByQuery(count bool) (string, []any) {
	var whereArray []any
	sql := &bytes.Buffer{}
	selectInfo := "*"
	if count {
		selectInfo = CountSql
	}
	// ex:SELECT * FROM account
	sql.WriteString(fmt.Sprintf(BaseSql, selectInfo, q.Table))
	// Where 语句
	if len(q.Where) > 0 {
		whereArray = q.makeWhereSql(sql)
	}
	// Order 语句
	if len(q.Order) > 0 {
		q.makeOrderSql(sql)
	}
	return sql.String(), whereArray
}

// makeWhereSql 生成查询语句
func (q *Query) makeWhereSql(sql *bytes.Buffer) []any {
	whereArray := make([]any, len(q.Where))
	for i, where := range q.Where {
		if i == 0 {
			sql.WriteString(" WHERE ")
		} else {
			sql.WriteString(" AND ")
		}
		// 创建Where子句
		subWhere := where.SubWhere
		subWhere = strings.ReplaceAll(subWhere, "${name}", where.Name)
		whereArray[i] = where.Value
		sql.WriteString(subWhere)
	}
	return whereArray
}

// makeOrderSql 创建排序语句
func (q *Query) makeOrderSql(sql *bytes.Buffer) {
	for i, order := range q.Order {
		if i == 0 {
			sql.WriteString(" ORDER BY ")
		} else {
			// 逗号分隔
			sql.WriteString(", ")
		}
		sql.WriteString(order.Name)
		// 降序指定
		if order.Desc {
			sql.WriteString(" DESC")
		}
	}
}

// CountByQuery 自定义Query结构组装查询
func (q *Query) countByQuery(db *gorm.DB) (int64, error) {
	total := int64(0)
	// 不包含Limit
	sql, whereArray := q.makeSqlByQuery(true)
	r := db.Raw(sql, whereArray...).Find(&total)
	return total, r.Error
}

// GetByQuery 自定义Query结构组装查询
func (q *Query) GetByQuery(db *gorm.DB, res any) error {
	// 不包含Limit
	sql, whereArray := q.makeSqlByQuery(false)
	dbQuery := db.Raw(sql, whereArray...)
	// 如果有Limit
	if q.LimitRange != nil {
		dbQuery.Offset(q.LimitRange.Offset).Limit(q.LimitRange.Limit)
	}
	r := dbQuery.Find(res)
	return r.Error
}
