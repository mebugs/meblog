import { defineStore } from 'pinia'

interface AppState {
  collapsed: boolean // 折叠
  theme: string // 主题
  headIconSize: any // 顶部图标的字体大小
}
const appStore = defineStore('app', {
  state: (): AppState => ({
    collapsed: false,
    theme: localStorage.getItem('theme') ? (localStorage.getItem('theme') as string) : 'light',
    headIconSize: { fontSize: '22px' }
  }),
  getters: {
    getHeadIconSize(): any {
      return this.headIconSize
    }
  },
  actions: {
    initTheme() {
      // 初始化主题
      document.body.setAttribute('arco-theme', this.theme)
    },
    setTheme() {
      // 主题状态更新
      this.theme = this.theme == 'dark' ? 'light' : 'dark'
      document.body.setAttribute('arco-theme', this.theme)
      localStorage.setItem('theme', this.theme)
    },
    setCollapsed() {
      // 折叠状态更新
      this.collapsed = !this.collapsed
    }
  }
})

export default appStore
