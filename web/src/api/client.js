import axios from 'axios';
import { Message } from "@arco-design/web-vue";

export const client = axios.create({
  baseURL: '',
  timeout: 3000,
})

client.interceptors.response.use(
  (value) => {
    return value.data
  },
  (error) => {
    console.log(error)
    let msg = error.message
    try {
      const data = error.response.data
      if (typeof data === 'string') {
        msg = data
      } else if (typeof data === 'object' && data !== null) {
        msg = data.message || data.msg || data.error || JSON.stringify(data)
      }
    } catch (e) {
      console.log(e)
    }
    Message.error(msg)
  }
)
