import request from './request'

export function getOrderDetail(sn) {
  return request.get('/travel/v1/order/detail', { params: { sn } })
}

export function payOrder(data) {
  return request.post('/travel/v1/order/pay', data)
}

export function cancelOrder(sn) {
  return request.post('/travel/v1/order/cancel', { sn })
}
