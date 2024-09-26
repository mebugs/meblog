const TOKEN_KEY = 'token'
const TOKEN_ACC = 'acc'

const isLogin = () => {
  return !!localStorage.getItem(TOKEN_KEY)
}

const getToken = () => {
  return localStorage.getItem(TOKEN_KEY)
}

const setToken = (token: string) => {
  localStorage.setItem(TOKEN_KEY, token)
}

const clearToken = () => {
  localStorage.removeItem(TOKEN_KEY)
}

const getAcc = () => {
  const acc = localStorage.getItem(TOKEN_ACC)
  if (acc) {
    return JSON.parse(window.atob(acc))
  }
  return null
}

const setAcc = (acc: any) => {
  localStorage.setItem(TOKEN_ACC, window.btoa(JSON.stringify(acc)))
}

const clearAcc = () => {
  localStorage.removeItem(TOKEN_ACC)
}

export { isLogin, getToken, setToken, clearToken, getAcc, setAcc, clearAcc }
