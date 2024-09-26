package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"sort"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkPermissionDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "alias_uni") {
			// 权限别名唯一
			return baseModel.Fail(constant.PermissionUniAliasNG)
		}
		if strings.Contains(errStr, "name_uni") {
			// 权限别名唯一
			return baseModel.Fail(constant.PermissionUniNameNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 递归权限树
func recursionPermissionTree(traceID string, treeNode *baseModel.Tree) (err error) {
	// 没有子级了
	if treeNode.Level == "3" {
		return
	}
	// 查询子集
	var permissionList platDB.PermissionArray
	permissionList, err = platDB.PermissionTable.GetByObject(&platDB.Permission{Pid: treeNode.Id})
	if err != nil {
		log.WarnTF(traceID, "RecursionPermissionTree Fail . PID %d . Err is : %s", treeNode.Id, err)
		return
	}
	if len(permissionList) == 0 {
		// 沒有子集推出
		return
	}
	// 数据排序
	sort.Sort(permissionList)
	treeNode.Children = make([]*baseModel.Tree, 0)
	// 组装子集
	for _, item := range permissionList {
		// 节点对象
		treeChild := &baseModel.Tree{
			Title:    item.Name,
			Key:      item.Alias,
			Expand:   item.Static == constant.StatusOpen, // 拓展用于选择框置灰
			Children: nil,
			Level:    item.Level,
			Id:       item.Id,
		}
		// 递归子集
		recursionPermissionTree(traceID, treeChild)
		// 加入子集
		treeNode.Children = append(treeNode.Children, treeChild)
	}
	return
}
