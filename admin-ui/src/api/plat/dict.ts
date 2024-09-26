import axios from 'axios'
import { ref } from 'vue'

export type dict = {
  id: number // 数据ID
  groupKey: string // 字典分组KEY
  label: string // 字段名称
  labelEn: string // 字段名称（英文）
  choose: string // 是否可被选择 0可选择 1不可选择
  val: string // 字典值（字符型）
  remark: string // 字典描述
}
export function dictInit() {
  return ref<dict>({
    id: 0,
    groupKey: '',
    label: '',
    labelEn: '',
    choose: '0',
    val: '',
    remark: ''
  })
}

export function dictNextVal(req: string) {
  return axios.post('/plat/dict/nextVal', { groupKey: req })
}

export function dictAdd(req: dict) {
  return axios.post('/plat/dict/add', req)
}

export function dictPage(req: any) {
  return axios.post('/plat/dict/page', req)
}

export function dictGet(req: number) {
  return axios.post('/plat/dict/get', { id: req })
}

export function dictEdit(req: dict) {
  return axios.post('/plat/dict/edit', req)
}

export function dictBro(req: string) {
  return axios.post('/plat/dict/bro', { groupKey: req })
}

export function dictSort(req: any) {
  return axios.post('/plat/dict/sort', req)
}

export function dictDel(req: any) {
  return axios.post('/plat/dict/del', { id: req })
}

export function dictRead(req: any) {
  return axios.post('/plat/dict/read', req)
}
