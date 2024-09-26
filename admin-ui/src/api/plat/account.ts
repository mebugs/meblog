import axios from 'axios'
import { ref } from 'vue'

export type account = {
  id: number // 数据ID
  account: string // 登陆账号
  name: string // 姓名
  deptId: string // 部门ID
  deptName: string // 部门名
  isLeader: string // 部门职位，枚举：0_部门员工 1_部门领导
  permissionType: string // 权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局
  roleIds: Array<number> // 关联的角色
  roleNames: Array<string> // 关联的角色名
  status: string // 状态，枚举：0_正常 1_锁定 2_封存
}

export function accountInit() {
  return ref<account>({
    id: 0,
    account: '',
    name: '',
    deptId: '',
    deptName: '',
    isLeader: '0',
    permissionType: '0',
    roleIds: [],
    roleNames: [],
    status: '0'
  })
}

export function accountAdd(req: account) {
  return axios.post('/plat/account/add', req)
}

export function accountPage(req: any) {
  return axios.post('/plat/account/page', req)
}

export function accountGet(req: number) {
  return axios.post('/plat/account/get', { id: req })
}

export function accountEdit(req: account) {
  return axios.post('/plat/account/edit', req)
}

export function accountDel(req: any) {
  return axios.post('/plat/account/del', { id: req })
}

export function accountReset(req: any) {
  return axios.post('/plat/account/reset', { id: req })
}
