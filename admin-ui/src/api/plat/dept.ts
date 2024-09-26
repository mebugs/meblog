import axios from 'axios'
import { ref } from 'vue'

export type dept = {
  id: number // 数据ID
  name: string // 部门名称
  pid: number // 父级部门ID，租户创建时默认创建根部门，父级ID=1
  permissionType: string // 权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局 4_指定部门 5_指定人
  accounts: Array<Array<any>> // 部门账号
}
export function deptInit() {
  return ref<dept>({
    id: 0,
    name: '',
    pid: 1,
    permissionType: '0',
    accounts: [[], []]
  })
}

export function deptAdd(req: dept) {
  return axios.post('/plat/dept/add', req)
}

export function deptTree(req: any) {
  return axios.post('/plat/dept/tree', req)
}

export function deptGet(req: number) {
  return axios.post('/plat/dept/get', { id: req })
}

export function deptEdit(req: dept) {
  return axios.post('/plat/dept/edit', req)
}

export function deptDel(req: any) {
  return axios.post('/plat/dept/del', { id: req })
}

export function deptBro(req: any) {
  return axios.post('/plat/dept/bro', { id: req })
}

export function deptSort(req: any) {
  return axios.post('/plat/dept/sort', req)
}

export function deptTo(req: any) {
  return axios.post('/plat/dept/to', req)
}
