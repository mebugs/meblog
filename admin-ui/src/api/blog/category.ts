import axios from 'axios'
import { ref } from 'vue'

export type category = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  sourceShow: string
  source: Array<string> // 选择数组（资源选择）
}

export function categoryInit() {
  return ref<category>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    sourceShow: '',
    source: []
  })
}

export function categoryAdd(req: category) {
  return axios.post('/blog/category/add', req)
}

export function categoryPage(req: any) {
  return axios.post('/blog/category/page', req)
}

export function categoryGet(req: number) {
  return axios.post('/blog/category/get', { id: req })
}

export function categoryEdit(req: category) {
  return axios.post('/blog/category/edit', req)
}

export function categoryDel(req: any) {
  return axios.post('/blog/category/del', { id: req })
}

export function categoryMerge(req: any) {
  return axios.post('/blog/category/merge', req)
}

export function categoryList() {
  return axios.post('/blog/category/list')
}
