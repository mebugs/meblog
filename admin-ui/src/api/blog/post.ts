import axios from 'axios'
import { ref } from 'vue'

export type post = {
  id: number // 数据ID
  title: string // 标题
  summary: string // 摘要
  img: string // 主图
  group: string // 分组
  tags: Array<string> // 关联标签
  status: string // 状态，枚举：0_发布 1_草稿
}

export function postInit() {
  return ref<post>({
    id: 0,
    title: '',
    summary: '',
    img: '',
    group: '',
    tags: [],
    status: '0'
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
