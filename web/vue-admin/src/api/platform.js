import request from '@/utils/request'

export function getPlatformList(params) {
  return request({
    url: './api/platform/page',
    method: 'get',
    data: {
      ...params
    }
  })
}
export function create(params) {
  return request({
    url: './api/platform/create',
    method: 'get',
    data: {
      ...params
    }
  })
}
export function deleteAction(params) {
  return request({
    url: './api/platform/delete',
    method: 'get',
    data: {
      ...params
    }
  })
}

