import request from '@/utils/request'

export function getList(params) {
  return request({
    url: './api/message/page',
    method: 'get',
    data: {
      ...params
    }
  })
}
