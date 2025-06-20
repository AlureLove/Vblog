import { useStorage } from '@vueuse/core'

export const token = useStorage(
  'token',
  {
    access_token: '',
    ref_user_name: '',
    refresh_token: '',
  },
  localStorage,
  {
    mergeDefaults: true,
  },
)
