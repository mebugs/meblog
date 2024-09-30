import { MAIN_LAYOUT } from './base'
import type { RouteRecordRaw } from 'vue-router'

// 博客路由
const BLOG: RouteRecordRaw = {
  path: '/blog', // 博客管理
  name: 'Blog', // 充当权限Alias
  component: MAIN_LAYOUT, // 布局
  meta: { locale: 'menu.Blog', icon: 'icon-book', requiresAuth: false }, // 国际化、图标、鉴权标识
  children: [
    {
      path: 'post', // 文章
      name: 'BlogPost',
      component: () => import('@/views/blog/post/index.vue'),
      meta: { locale: 'menu.BlogPost', requiresAuth: false }
    },
    {
      path: 'group', // 分类
      name: 'BlogGroup',
      component: () => import('@/views/blog/group/index.vue'),
      meta: { locale: 'menu.BlogGroup', requiresAuth: false }
    },
    {
      path: 'topic', // 专题
      name: 'BlogTopic',
      component: () => import('@/views/blog/topic/index.vue'),
      meta: { locale: 'menu.BlogTopic', requiresAuth: false }
    },
    {
      path: 'tag', // 标签
      name: 'BlogTag',
      component: () => import('@/views/blog/tag/index.vue'),
      meta: { locale: 'menu.BlogTag', requiresAuth: false }
    },
    {
      path: 'page', // 页面
      name: 'BlogPage',
      component: () => import('@/views/blog/page/index.vue'),
      meta: { locale: 'menu.BlogPage', requiresAuth: false }
    },
    {
      path: 'Comment', // 评论
      name: 'BlogComment',
      component: () => import('@/views/blog/comment/index.vue'),
      meta: { locale: 'menu.BlogComment', requiresAuth: false }
    }
  ]
}

export default BLOG
