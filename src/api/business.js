import request from './request'

export function getBusinessDetail(id) {
  return request.get(`/business/detail/${id}`)
}

export function getBusinessHomestayList(businessId, params) {
  return request.get(`/business/${businessId}/homestay`, { params })
}
