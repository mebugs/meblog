package actuator

import (
	"siteol.com/smart/src/common/model/baseModel"

	"gorm.io/gorm"
)

// MYSQL 执行器
// 常见数据库执行API

// Executor 泛型接口
type Executor interface {
	TableName() string
	DataBase() *gorm.DB
}

// Table 泛型执行器
type Table[T Executor] string

// Executor 返回执行本体
func (t Table[T]) Executor() T {
	var exe T
	return exe
}

// GetAll 查询表所有数据
func (t Table[T]) GetAll() (res []*T, err error) {
	var exe T
	r := exe.DataBase().Find(&res)
	err = r.Error
	return
}

// GetOneById 查询表指定ID的数据
func (t Table[T]) GetOneById(id uint64) (res T, err error) {
	var exe T
	r := exe.DataBase().Where("id = ?", id).First(&res)
	err = r.Error
	return
}

// GetOneByObject 查询表指定对象的数据（传入指针）
func (t Table[T]) GetOneByObject(req any) (res T, err error) {
	var exe T
	r := exe.DataBase().Where(req).First(&res)
	err = r.Error
	return
}

// GetByIds 查询表指定ID的数据集
func (t Table[T]) GetByIds(ids []uint64) (res []*T, err error) {
	var exe T
	r := exe.DataBase().Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}

// GetByObject 查询表指定对象的数据集
func (t Table[T]) GetByObject(req any) (res []*T, err error) {
	var exe T
	r := exe.DataBase().Where(req).Find(&res)
	err = r.Error
	return
}

// CountByObject 查询表指定对象的数量
func (t Table[T]) CountByObject(req any) (res int64, err error) {
	var exe T
	r := exe.DataBase().Model(req).Where(req).Count(&res)
	err = r.Error
	return
}

// GetByObjectSort 查询表指定对象，默认使用sort字段排序的数据集
func (t Table[T]) GetByObjectSort(req any) (res []*T, err error) {
	var exe T
	r := exe.DataBase().Where(req).Order("sort").Find(&res)
	err = r.Error
	return
}

// CountByQuery 查询表指定对象，默认使用sort字段排序的数据集
func (t Table[T]) CountByQuery(query *Query) (total int64, err error) {
	var exe T
	// 提交表名
	query.Table = exe.TableName()
	// 先查询数量
	total, err = query.countByQuery(exe.DataBase())
	return
}

// Page 平台的公共分页查询方法
func (t Table[T]) Page(query *Query) (total int64, list []*T, err error) {
	var exe T
	// 提交表名
	query.Table = exe.TableName()
	// 先查询数量
	total, err = query.countByQuery(exe.DataBase())
	if err != nil || total == 0 {
		return
	}
	// 开始查询分页数据
	err = query.GetByQuery(exe.DataBase(), &list)
	if err != nil {
		return
	}
	return
}

// InsertOne 插入一条数据
func (t Table[T]) InsertOne(req any) (err error) {
	var exe T
	r := exe.DataBase().Create(req)
	err = r.Error
	return
}

// InsertBatch 插入多条数据
func (t Table[T]) InsertBatch(req any) (err error) {
	var exe T
	r := exe.DataBase().Create(req)
	err = r.Error
	return
}

// UpdateOne 更新一条数据（根据ID）
func (t Table[T]) UpdateOne(req any) (err error) {
	var exe T
	r := exe.DataBase().Save(req)
	err = r.Error
	return
}

// UpdateByIds 更新多条数据（保持一致的更新）
func (t Table[T]) UpdateByIds(ids []uint64, req any) (err error) {
	var exe T
	r := exe.DataBase().Table(exe.TableName()).Where("id IN ?", ids).Updates(req)
	err = r.Error
	return
}

// DeleteOne 删除一条数据（根据ID）
func (t Table[T]) DeleteOne(id uint64) (err error) {
	var exe T
	r := exe.DataBase().Delete(&exe, id)
	//.Exec(fmt.Sprintf("DELETE FROM %s WHERE `id`= ?", exe.TableName()), id)
	err = r.Error
	return
}

// DeleteByIds 删除多条数据（根据ID）
func (t Table[T]) DeleteByIds(ids []uint64) (err error) {
	var exe T
	r := exe.DataBase().Delete(&exe, ids)
	//.Exec(fmt.Sprintf("DELETE FROM %s WHERE `id`IN ?", exe.TableName()), ids)
	err = r.Error
	return
}

// SortWithTransaction 事务+排序
func (t Table[T]) SortWithTransaction(req []*baseModel.SortReq) error {
	var exe T
	db := exe.DataBase()
	mod := new(T)
	// 启用事务
	return db.Transaction(func(tx *gorm.DB) error {
		for _, item := range req {
			// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
			if err := tx.Model(mod).Where("id = ?", item.ID).Update("sort", item.Sort).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
}
