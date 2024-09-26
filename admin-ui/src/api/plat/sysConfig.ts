import axios from 'axios'
import { ref } from 'vue'

export type sysConfig = {
  loginSwitch: string // 并发限制开关，0限制 1不限制
  loginNum: number // 最大登陆并发量，最小为1
  loginFailSwitch: string // 登陆失败限制开关，0限制 1不限制
  loginFailUnit: string // 登陆失败限制 1秒 2分 3时 4天
  loginFailNum: number // 登陆失败最大尝试次数，最小为1
  loginFailLockUnit: string // 登陆失败锁定 1秒 2分 3时 4天
  loginFailLockNum: number // 登陆失败锁定数量，最小为1
  loginFailTryNum: number // 失败尝试次数
  logoutSwitch: string // 登陆过期开关，0限制 1不限制
  logoutUnit: string // 登陆过期单位，1秒 2分 3时 4天
  logoutNum: number // 登陆过期长度数量，最小为1
}
export function sysConfigInit() {
  return ref<sysConfig>({
    loginSwitch: '1',
    loginNum: 1,
    loginFailSwitch: '1',
    loginFailUnit: '0',
    loginFailNum: 1,
    loginFailLockUnit: '0',
    loginFailLockNum: 1,
    loginFailTryNum: 1,
    logoutSwitch: '1',
    logoutUnit: '0',
    logoutNum: 1
  })
}

export function sysConfigGet() {
  return axios.post('/plat/sysConfig/get')
}

export function sysConfigEdit(req: sysConfig) {
  return axios.post('/plat/sysConfig/edit', req)
}
