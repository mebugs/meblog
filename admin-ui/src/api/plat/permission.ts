import axios from 'axios'
import { ref } from 'vue'

export type permission = {
  id: number // 数据ID
  level: string // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现
  pid: number // 父级ID，默认为1
  static: string // 默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾
  name: string // 权限名称，界面展示，建议与界面导航一致
  alias: string // 权限别名，英文，规范如下：sys，sysAccount sysAccountAdd
  routerIds: number[] // 路由集，提交路由ID数组
  routers: any[] // 路由列表
}
export function permissionInit() {
  return ref<permission>({
    id: 0,
    level: '',
    pid: 1,
    static: '1',
    name: '',
    alias: '',
    routerIds: [],
    routers: []
  })
}

export function permissionAdd(req: permission) {
  return axios.post('/plat/permission/add', req)
}

export function permissionTree(req: any) {
  return axios.post('/plat/permission/tree', req)
}

export function permissionGet(req: number) {
  return axios.post('/plat/permission/get', { id: req })
}

export function permissionEdit(req: permission) {
  return axios.post('/plat/permission/edit', req)
}

export function permissionDel(req: any) {
  return axios.post('/plat/permission/del', { id: req })
}

export function permissionBro(req: any) {
  return axios.post('/plat/permission/bro', { id: req })
}

export function permissionSort(req: any) {
  return axios.post('/plat/permission/sort', req)
}
