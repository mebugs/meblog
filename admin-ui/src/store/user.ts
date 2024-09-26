import { defineStore } from 'pinia'
import { RouteRecordRaw } from 'vue-router'
import routers from '@/router/routers'
import userMenuTree from '@/utils/hooks/menu'
//import { logout as userLogout, getUserInfo } from '@/api/user'
import { authMine, authLogout } from '@/api/base/auth'
import { clearToken } from '@/utils/hooks/token'
import { removeRouteListener } from '@/utils/routeListener'

interface UserState {
  accountId: Number
  name?: string
  avatar?: string
  permissions: string[] // 权限列表
  menusTree: RouteRecordRaw[] // 菜单树
}

const userStore = defineStore('user', {
  state: (): UserState => ({
    accountId: 0,
    name: undefined,
    avatar: undefined,
    permissions: [],
    menusTree: []
  }),
  getters: {
    userInfo(state: UserState): UserState {
      return { ...state }
    }
  },
  actions: {
    initMenusTree() {
      const { recursionRouters } = userMenuTree()
      // 初始化用户菜单树
      this.menusTree = recursionRouters(routers)
    },
    // Reset user's information
    resetInfo() {
      this.$reset()
    },
    // Get user's information
    async getUserInfo() {
      try {
        const res = await authMine()
        this.accountId = res.data.accountId
        this.name = res.data.name
        this.permissions = res.data.permissions
      } catch (err) {
        throw err
      }
    },
    // 登出
    async logout() {
      try {
        authLogout()
      } finally {
        this.resetInfo()
        clearToken()
        removeRouteListener()
        // 跳去着陆页
        window.location.replace('/login')
      }
    }
  }
})

export default userStore
