import axios from 'axios'
import { ref } from 'vue'

export type router = {
  id: number // 数据ID
  type: string // 免授权路由 0 免授权 1 授权
  url: string // 路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容
  name: string // 路由名称，用于界面展示，与权限关联
  serviceCode: string // 业务编码（字典），为接口分组
  logInDb: string // 日志入库 0 启用 1 默认不启用
  reqLogPrint: string // 请求日志打印 0 打印 1 不打印
  reqLogSecure: string // 请求日志脱敏字段，逗号分隔，打印时允许配置
  resLogPrint: string // 响应日志打印 0 打印 1 不打印
  resLogSecure: string // 响应日志脱敏字段，逗号分隔，打印时允许配置`
  remark: string // 备注信息
}
export function routerInit() {
  return ref<router>({
    id: 0,
    type: '0',
    url: '',
    name: '',
    serviceCode: '',
    logInDb: '1',
    reqLogPrint: '0',
    reqLogSecure: '',
    resLogPrint: '0',
    resLogSecure: '',
    remark: ''
  })
}

export function routerAdd(req: router) {
  return axios.post('/plat/router/add', req)
}

export function routerPage(req: any) {
  return axios.post('/plat/router/page', req)
}

export function routerGet(req: number) {
  return axios.post('/plat/router/get', { id: req })
}

export function routerEdit(req: router) {
  return axios.post('/plat/router/edit', req)
}

export function routerDel(req: any) {
  return axios.post('/plat/router/del', { id: req })
}
