import { Plugin } from '@nuxt/types'
import { AxiosError, AxiosRequestConfig } from 'axios'
import { initializeAxios } from '~/plugins/axios'

const accessor: Plugin = ({ store, $axios }) => {
  initializeAxios($axios)

  $axios.onRequest((config: AxiosRequestConfig) => {
    config.baseURL = process.env.apiURL
    config.timeout = 10000 // 10sec
    config.withCredentials = true

    const token: string = store.getters['auth/getToken']
    if (token) {
      config.headers.common['Authorization'] = `Bearer ${token}`
    }

    return config
  })

  $axios.onError((err: AxiosError) => {
    return Promise.reject(err)
  })
}

export default accessor
