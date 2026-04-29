import request from './request'

export function getCommentList(params) {
  return request.get('/comment/list', { params })
}

export function getCommentByHomestay(homestayId, params) {
  return request.get(`/comment/homestay/${homestayId}`, { params })
}

export function getCommentByBusiness(businessId, params) {
  return request.get(`/comment/business/${businessId}`, { params })
}

export function createComment(data) {
  return request.post('/comment/create', data)
}

export function uploadCommentImage(file) {
  const formData = new FormData()
  formData.append('file', file)
  return request.post('/comment/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}
