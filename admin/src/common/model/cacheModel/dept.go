package cacheModel

import (
	"encoding/json"
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"sort"
	"time"
)

// SyncDeptTreeCache 同步刷新部门树
func SyncDeptTreeCache(traceID string) (err error) {
	// 先移除缓存
	_ = redis.Del(constant.CacheDeptTrees)
	// 开始构建树对象
	// 查询根节点
	rootPerm, err := platDB.DeptTable.GetOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "SyncDeptTreeCache GetRoot Fail . Err Is : %s", err)
	}
	treeNode := &baseModel.Tree{
		Title:    rootPerm.Name,
		Key:      fmt.Sprintf("%d", rootPerm.Id),
		Children: nil,
		Expand:   rootPerm.Name,
		Level:    constant.StatusOpen, // 跟层不可移动
		Id:       rootPerm.Id,
	}
	// 递归部门树
	_ = recursionDeptTree(traceID, treeNode)
	// 加入缓存
	err = redis.Set(constant.CacheDeptTrees, treeNode, 0)
	if err != nil {
		log.ErrorTF(traceID, "SyncDeptTreeCache SetCache Fail . Err Is : %s", err)
	}
	return
}

// GetDeptTreeCache 获取部门树
func GetDeptTreeCache(traceID string) (treeNode *baseModel.Tree, err error) {
	deptTreeStr := ""
	for {
		str, errG := redis.Get(constant.CacheDeptTrees)
		if errG != nil {
			if errG != redis.ErrNotFound {
				log.ErrorTF(traceID, "GetDeptTreeCache Fail . Err Is : %s", errG)
				err = errG
				return
			}
		}
		if str != "" {
			deptTreeStr = str
			break
		}
		// 刚好遇到缓存处理中
		time.Sleep(50 * time.Millisecond)
	}
	treeNode = &baseModel.Tree{}
	err = json.Unmarshal([]byte(deptTreeStr), treeNode)
	if err != nil {
		log.ErrorTF(traceID, "UnmarshalDeptTreeCache Fail . Err Is : %s", err)
	}
	return
}

// 递归部门树
func recursionDeptTree(traceID string, treeNode *baseModel.Tree) (err error) {
	// 查询子集
	var deptList platDB.DeptArray
	deptList, err = platDB.DeptTable.GetByObject(&platDB.Dept{Pid: treeNode.Id})
	if err != nil {
		log.WarnTF(traceID, "RecursionDeptTree Fail . PID %d . Err is : %s", treeNode.Id, err)
		return
	}
	if len(deptList) == 0 {
		// 沒有子集推出
		return
	}
	// 数据排序
	sort.Sort(deptList)
	treeNode.Children = make([]*baseModel.Tree, 0)
	// 组装子集
	for _, item := range deptList {
		// 节点对象
		treeChild := &baseModel.Tree{
			Title:    item.Name,
			Key:      fmt.Sprintf("%d", item.Id),
			Children: nil,
			Expand:   item.Name,
			Level:    constant.StatusLock, // 可以移动
			Id:       item.Id,
		}
		// 递归子集
		_ = recursionDeptTree(traceID, treeChild)
		// 加入子集
		treeNode.Children = append(treeNode.Children, treeChild)
	}
	return
}
