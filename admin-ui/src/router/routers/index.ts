import type { RouteRecordRaw } from 'vue-router'
import BASE from './base' // 基础路由
import BLOG from './blog' // 博客路由
import PLAT from './plat' // 平台路由

// 导出路由
const routers: RouteRecordRaw[] = [...BASE, BLOG, PLAT]
export default routers
