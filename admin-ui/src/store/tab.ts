import { defineStore } from 'pinia'
import type { RouteLocationNormalized } from 'vue-router'

// 默认路由地址(工作中心)
export const DEFAULT_ROUTE_NAME = 'CenterIndex'
export const DEFAULT_ROUTE = {
  title: 'menu.CenterIndex',
  name: DEFAULT_ROUTE_NAME,
  fullPath: '/center/index',
  isHome: true // 影响左侧图标的展示
}

export interface TagProps {
  title: string
  name: string
  fullPath: string
  query?: any
  isHome: boolean
}

// 转换路由到Tag对象
const routerToTag = (route: RouteLocationNormalized): TagProps => {
  const { name, meta, fullPath, query } = route
  return {
    title: meta.locale || '',
    name: String(name),
    fullPath,
    query,
    isHome: false // 执行转换的路由不是Home页
  }
}

export interface TabState {
  tagList: TagProps[]
  cacheTabList: Set<string>
}
const tabStore = defineStore('tab', {
  state: (): TabState => ({
    cacheTabList: new Set([DEFAULT_ROUTE_NAME]),
    tagList: [DEFAULT_ROUTE]
  }),
  getters: {
    getTabList(): TagProps[] {
      return this.tagList
    },
    getCacheList(): string[] {
      return Array.from(this.cacheTabList)
    }
  },
  actions: {
    addTabList(route: RouteLocationNormalized) {
      // 添加Tab可能添加缓存
      this.tagList.push(routerToTag(route))
      this.cacheTabList.add(route.name as string)
    },
    removeTabList(name: string, index: number) {
      // 移除Tab
      this.tagList.splice(index, 1)
      this.cacheTabList.delete(name)
    }
  }
})

export default tabStore
