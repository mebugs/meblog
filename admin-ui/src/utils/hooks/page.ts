// 通用的分页构造器
import { reactive } from 'vue'

export default function usePage() {
  // 默认分页对象
  const page = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showTotal: true
  })
  let query: Function
  const setQuery = (q: Function) => {
    query = q
  }
  // 搜索方法
  const search = () => {
    page.current = 1
    query()
  }
  // 切换分页
  const changePage = (current: number) => {
    page.current = current
    query()
  }
  // 重置分页
  const resetPage = () => {
    page.current = 1
    page.total = 0
  }
  return {
    page,
    setQuery,
    search,
    changePage,
    resetPage
  }
}
