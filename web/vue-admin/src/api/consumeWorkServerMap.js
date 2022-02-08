import request from '@/utils/request'

export function getList(params) {
  return request({
    url: './api/consumeConfig/workList',
    method: 'get',
    data: { ...params }
  })
}

export function create(params) {
  return request({
    url: './api/consumeServerMap/create',
    method: 'get',
    data: {
      ...params
    }
  })
}

export function update(params) {
  return request({
    url: './api/consumeServerMap/update',
    method: 'get',
    data: {
      ...params
    }
  })
}

export function deleteAction(params) {
  return request({
    url: './api/consumeServerMap/delete',
    method: 'get',
    data: {
      ...params
    }
  })
}
