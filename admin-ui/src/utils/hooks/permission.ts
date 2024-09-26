import { RouteLocationNormalized, RouteRecordRaw } from 'vue-router'
import { userStore } from '@/store'

// 路由权限判定
export default function userPermission() {
  return {
    async permissionRouter(route: RouteLocationNormalized | RouteRecordRaw) {
      const user = userStore()
      const routerName = route.name ? route.name.toString() : ''
      // 无路由名或免授权返回成功
      if (routerName === '' || !route.meta?.requiresAuth) {
        return true
      }
      // 授权情况
      return user.permissions.includes(routerName)
    }
  }
}
