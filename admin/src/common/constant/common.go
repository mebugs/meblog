package constant

import "siteol.com/smart/src/config"

// TransLangSupport 支持更多语言请添加
var TransLangSupport = []string{"zh-CN", "en-US"}

const (
	HeaderToken = "Token" // 固定请求头（登陆Token）

	ContextLang     = "Lang"       // 语言
	ContextTraceID  = "TraceID"    // 日志链路跟踪ID
	ContextRouterC  = "RouterConf" // 路由配置对象
	ContextRouterI  = "RouterInDb" // 路由入库对象
	ContextAuthUser = "AuthUser"   // 授权用户对象

	StatusOpen  = "0" // 正常 启动 可变更 选择
	StatusLock  = "1" // 禁用 锁定 登出 不可变更 半选
	StatusClose = "2" // 移除 弃用 踢出

	DataPermissionTree = "0" // 数据权限继承部门或本部门与子部门
	DataPermissionDept = "1" // 数据权限本部门
	DataPermissionSelf = "2" // 数据权限个人
	DataPermissionAll  = "3" // 数据权限全部

	DBDuplicateErr = "Error 1062 (23000): Duplicate entry" // 唯一索引错误
)

var (
	ProjectName            = config.JsonConfig.Server.Name           // 项目名
	CacheSysConf           = ProjectName + ":Comm:SysConf"           // 系统配置缓存
	CacheResTrans          = ProjectName + ":Comm:ResTrans"          // 响应码缓存
	CacheRouters           = ProjectName + ":Comm:Routers"           // 路由缓存
	CacheRouterUrls        = ProjectName + ":Comm:RouterUrls"        // 路由地址缓存
	CachePermissions       = ProjectName + ":Comm:Permissions"       // 权限缓存
	CachePermissionsNormal = ProjectName + ":Comm:PermissionsNormal" // 默认权限缓存
	CacheRoles             = ProjectName + ":Comm:Roles"             // 角色缓存
	CacheDeptTrees         = ProjectName + ":Comm:DeptTrees"         // 部门缓存树缓存
	CacheAuth              = ProjectName + ":Auth:%s"                // 登陆授权缓存
)
