import request from './request'

export function loginByCode(data) {
  return request.post('/user/login/code', data)
}

export function loginByPassword(data) {
  return request.post('/usercenter/v1/user/login', data)
}

export function sendVerifyCode(phone) {
  return request.post('/user/send-code', { phone })
}

export function register(data) {
  return request.post('/usercenter/v1/user/register', data)
}

export function getUserInfo() {
  return request.get('/user/info')
}

export function logout() {
  return request.post('/user/logout')
}
