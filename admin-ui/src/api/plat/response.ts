import axios from 'axios'
import { ref } from 'vue'

export type response = {
  id: number // 数据ID
  serviceCode: string // 业务ID，来源于字典，指定响应码归属业务
  type: string // 响应类型，该字段用于筛选，可配置S和F
  code: string // 响应码，仅示例，实际入库实时计算
  zhCn: string // 中文响应文言
  enUs: string // 英文响应文言
  remark: string // 其他备注信息
}
export function responseInit() {
  return ref<response>({
    id: 0,
    serviceCode: '',
    type: '',
    code: '',
    zhCn: '',
    enUs: '',
    remark: ''
  })
}

export function responseNextVal(req: any) {
  return axios.post('/plat/response/nextVal', req)
}

export function responseAdd(req: response) {
  return axios.post('/plat/response/add', req)
}

export function responsePage(req: any) {
  return axios.post('/plat/response/page', req)
}

export function responseGet(req: number) {
  return axios.post('/plat/response/get', { id: req })
}

export function responseEdit(req: response) {
  return axios.post('/plat/response/edit', req)
}

export function responseDel(req: any) {
  return axios.post('/plat/response/del', { id: req })
}

export function responseRead(req: any) {
  return axios.post('/plat/response/read', req)
}
