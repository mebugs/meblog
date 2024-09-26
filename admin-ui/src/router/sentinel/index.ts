/**
 * 路由哨兵
 * 监听路由变化
 * 传递路由通信
 * 判定路由权限
 */
import { setRouteEmitter } from '@/utils/routeListener'
import type { Router } from 'vue-router'
import { Modal } from '@arco-design/web-vue'
import i18n from '@/locale'
import NProgress from 'nprogress' // progress bar
import { userStore } from '@/store'
import userPermission from '@/utils/hooks/permission'

const permission = userPermission()
// 路由哨兵发布路由变化消息
function routerSentinelPush(router: Router) {
  router.beforeEach(async (to) => {
    // emit route change
    setRouteEmitter(to)
  })
}

// 路由哨兵检查路由访问 - 登陆 - 授权
function routerSentinelCheck(router: Router) {
  router.beforeEach(async (to, from, next) => {
    console.log(from)
    NProgress.start()
    const user = userStore()
    // 没有账号信息
    if (to.meta?.requiresAuth && user.accountId == 0) {
      // 获取用户权限
      await user.getUserInfo()
    }
    // 没有权限，退出登录
    if (!permission.permissionRouter(to)) {
      Modal.error({
        title: i18n.global.t('base.noauth'),
        content: '', // 服务端翻译好的提示文言
        okText: i18n.global.t('button.ok'),
        maskClosable: false,
        async onOk() {
          // 登出处理
          const user = userStore()
          await user.logout()
        }
      })
      next({ name: 'login' })
      return
    }
    next()
  })
}

// 路由哨兵跳转后的处理
function routerSentinelAfter(router: Router) {
  router.afterEach(() => {
    NProgress.done()
  })
}

// 创建路由处理中间件
export default function routerSentinel(router: Router) {
  routerSentinelPush(router) // 路由建传递路有变化信息
  routerSentinelCheck(router) // 验证是否有权限
  routerSentinelAfter(router) // 跳转后的处理
}
