import { createRouter, createWebHistory } from 'vue-router'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css'
import routers from './routers'
import routerSentinel from './sentinel' // 路由哨兵
NProgress.configure({ showSpinner: false }) // NProgress Configuration

// 初始化路由
const router = createRouter({
  history: createWebHistory(),
  routes: routers,
  scrollBehavior() {
    return { top: 0 } // 切换路由滚动到顶部
  }
})

routerSentinel(router) // 启用路由哨兵

export default router
