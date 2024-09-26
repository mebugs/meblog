import { RouteRecordRaw } from 'vue-router'
import userPermission from './permission'
// 用户可访问的路由树
export default function userMenuTree() {
  const permission = userPermission()
  const recursionRouters = (_routes: RouteRecordRaw[]) => {
    if (!_routes) return []
    let uRouters: RouteRecordRaw[] = []
    _routes.forEach((r) => {
      let nR: RouteRecordRaw = { ...r }
      // 没有权限
      if (!permission.permissionRouter(r)) {
        return
      }
      // 不显示菜单
      if (r.meta?.hideInMenu) {
        return
      }
      // 默认不显示子节点
      nR.children = []
      if (!r.meta?.hideChildrenInMenu && r.children) {
        // 处理子节点 过滤不展示的子节点
        nR.children = r.children.filter((s) => s.meta?.hideInMenu !== true)
        // 递归处理子节点
        if (r.children.length > 0) {
          nR.children = recursionRouters(nR.children)
          // 只有一个子节点，子节点上升
          if (nR.children.length == 1) {
            let subR = nR.children[0]
            subR.path = nR.path + '/' + subR.path
            if (subR.meta && nR.meta) {
              subR.meta.icon = nR.meta.icon
            }
            nR = subR
          }
        }
      }
      uRouters.push(nR)
    })
    return uRouters
  }
  return { recursionRouters }
}
