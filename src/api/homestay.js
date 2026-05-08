import request from './request'

export function getHomestayList(params) {
  return request.post('/travel/v1/homestay/homestayList', params)
}

export function getHomestayDetail(params) {
  return request.get(`/travel/v1/homestay/homestay-detail`,{params})
}

export function getHomestayBossWithHomestayId(params) {
  return request.get(`/travel/v1/homestay/homestay-bussiness-info/with-homestay-id`,{params})
}

export function getHomestayBossWithId(params) {
  return request.get(`/travel/v1/homestay/homestay-bussiness-info/with-id`,{params})
}

export function getHomestayCommentList(params) {
  return request.get(`/travel/v1/homestay/homestay-comment-list`,{params})
}

export function searchHomestay(params) {
  return request.get('/homestay/search', { params })
}

export function getHomestayByCategory(category) {
  return request.get('/homestay/category', { params: { category } })
}

export function createBooking(data) {
  return request.post('/order/v1/order/create', data)
}

export function getCategories() {
  return request.get('/homestay/categories')
}
