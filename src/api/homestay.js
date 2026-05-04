import request from './request'

export function getHomestayList(params) {
  return request.post('/travel/v1/homestay/homestayList', params)
}

export function getHomestayDetail(id) {
  return request.get(`/homestay/detail/${id}`)
}

export function searchHomestay(params) {
  return request.get('/homestay/search', { params })
}

export function getHomestayByCategory(category) {
  return request.get('/homestay/category', { params: { category } })
}

export function createBooking(data) {
  return request.post('/homestay/booking', data)
}

export function getCategories() {
  return request.get('/homestay/categories')
}
