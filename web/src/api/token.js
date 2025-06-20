import { client } from './client'

export const LOGIN = (data) => client({
  method: 'POST',
  url: '/api/vblog/v1/tokens',
  data: data,
})
