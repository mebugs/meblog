import axios from 'axios'
import { ref } from 'vue'

export type role = {
  id: number // 数据ID
  name: string // 角色名称
  remark: string // 角色备注
  permissionIds: number[] // 权限集
  halfPermissionIds: number[] // 权限集（父级半选）
}
export function roleInit() {
  return ref<role>({
    id: 0,
    name: '',
    remark: '',
    permissionIds: [],
    halfPermissionIds: []
  })
}

export function roleAdd(req: role) {
  return axios.post('/plat/role/add', req)
}

export function rolePage(req: any) {
  return axios.post('/plat/role/page', req)
}

export function roleGet(req: number) {
  return axios.post('/plat/role/get', { id: req })
}

export function roleEdit(req: role) {
  return axios.post('/plat/role/edit', req)
}

export function roleDel(req: any) {
  return axios.post('/plat/role/del', { id: req })
}

export function roleList() {
  return axios.post('/plat/role/list')
}
