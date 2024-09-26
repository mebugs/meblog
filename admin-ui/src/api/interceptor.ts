import i18n from '@/locale'
import axios from 'axios'
import type { InternalAxiosRequestConfig } from 'axios'
import { Message, Modal } from '@arco-design/web-vue'
import { userStore } from '@/store'
import { getToken } from '@/utils/hooks/token'

let modalUnOpen = true

// 绑定请求URL
if (import.meta.env.VITE_API_BASE_URL) {
  axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL
}
// 超时时间（暂定100秒）
axios.defaults.timeout = 100000
// 请求拦截器
axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 为请求携带登陆Token
    const token = getToken()
    const lang = localStorage.getItem('arco-locale')
    // 提交Token
    config.headers.set('Token', `${token}`, true)
    // 提交语言
    config.headers.set('Lang', `${lang}`, true)
    return config
  },
  (error) => {
    // do something
    // eslint-disable-next-line no-console
    console.log(error) // for debug
    return Promise.reject(error)
  }
)
// 响应拦截器
axios.interceptors.response.use(
  (response) => {
    // 读取响应数据
    const res = response.data
    // 后端所有响应均进行了翻译处理
    // 特定响应码的处理逻辑
    // 成功类型响应码 S 开头 如S000
    if (res.code.indexOf('S') === 0) {
      // 成功是判断是否是静默Pop
      if (!res.unPop) {
        Message.success(res.msg)
      }
      return res
    }
    // 其他错误码均为业务错误码或500系统未知异常
    Message.error(res.msg)
    return Promise.reject(new Error(res.msg || 'Error'))
  },
  (error) => {
    // 读取响应数据
    const { response } = error
    const { status, data } = response
    // 错误文言
    let msg = data ? data.msg : 'Request Error'
    // 是Msg还是Box
    let isBox = false
    let boxTitle = 'base.nologin'
    // 处理错误
    switch (status) {
      // 400 请求数据不合法（校验失败）
      // 500 系统未知异常（意料之外的错误）
      case 400:
      case 500:
        break
      // 401 未授权、未登陆
      // 403 禁止访问（无权限）
      case 401:
        isBox = true
        break
      case 403:
        isBox = true
        boxTitle = 'base.noauth'
        break
      default:
        msg = error.msg
    }
    // 如果是box
    if (isBox) {
      if (modalUnOpen) {
        modalUnOpen = false
        Modal.error({
          title: i18n.global.t(boxTitle),
          content: '', // 服务端翻译好的提示文言
          okText: i18n.global.t('button.ok'),
          maskClosable: false,
          async onOk() {
            modalUnOpen = true
            // 登出处理
            const user = userStore()
            await user.logout()
          }
        })
      }
    } else {
      Message.error({
        content: msg || error.msg || 'Request Error',
        duration: 5 * 1000
      })
    }
    return Promise.reject(error)
  }
)
