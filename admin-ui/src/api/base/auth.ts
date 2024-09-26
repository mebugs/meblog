import axios from 'axios'

export function authLogin(req: any) {
  return axios.post('/auth/login', req)
}

export function authReset(req: any) {
  return axios.post('/auth/reset', req)
}

export function authMine() {
  return axios.post('/auth/mine')
}

export function authLogout() {
  return axios.post('/auth/logout')
}
