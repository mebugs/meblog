import axios from 'axios'
import { ref } from 'vue'

export type post = {
  id: number // 数据ID
  title: string // 标题
  url: string // 静态地址
  summary: string // 摘要
  img: string // 主图
  imgs: Array<string> // 主图缩略图
  group: string // 分组
  topic: string // 专题
  page: string // 页面
  tags: Array<string> // 关联标签
  status: string // 状态，枚举：0_发布 1_草稿
  md: string // markdown文本
  html: string // 编译的HTML
}

export function postInit() {
  return ref<post>({
    id: 0,
    title: '',
    url: '',
    summary: '',
    img: '',
    imgs: [],
    group: '',
    topic: '',
    page: '',
    tags: [],
    status: '',
    md: '',
    html: ''
  })
}

export function postAdd(req: post) {
  return axios.post('/plat/account/add', req)
}

export function postPage(req: any) {
  return axios.post('/plat/account/page', req)
}

export function postGet(req: number) {
  return axios.post('/plat/account/get', { id: req })
}

export function postEdit(req: post) {
  return axios.post('/plat/account/edit', req)
}

export function postDel(req: any) {
  return axios.post('/plat/account/del', { id: req })
}

export function postReset(req: any) {
  return axios.post('/plat/account/reset', { id: req })
}
