import { MAIN_LAYOUT } from './base'
import type { RouteRecordRaw } from 'vue-router'

// 博客路由
const BLOG: RouteRecordRaw = {
  path: '/blog', // 博客管理
  name: 'Blog', // 充当权限Alias
  component: MAIN_LAYOUT, // 布局
  meta: { locale: 'menu.Plat', icon: 'icon-book', requiresAuth: false }, // 国际化、图标、鉴权标识
  children: [
    {
      path: 'post', // 文章管理
      name: 'BlogPost',
      component: () => import('@/views/blog/post/index.vue'),
      meta: { locale: 'menu.BlogPost', requiresAuth: false }
    }
  ]
}

export default BLOG
