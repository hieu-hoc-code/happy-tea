import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import API_URL from './config'
const credentials = { withCredentials: true }
import JwtService from '@/common/jwt.service'

const ApiService = {
  init() {
    Vue.use(VueAxios, axios)
    Vue.axios.defaults.baseURL = API_URL
  },
  setHeader() {
    Vue.axios.defaults.headers.common[
      'Authorization'
    ] = `Token ${JwtService.getCartID()}`
  },
  post(resource, params) {
    return Vue.axios.post(`${resource}`, params, credentials)
  },
  put(resource, params) {
    return Vue.axios.put(`${resource}`, params, credentials)
  },
  get(resource, params) {
    return Vue.axios.get(`${resource}`, { ...params, ...credentials })
  },
  delete(resource, params) {
    console.log({ ...params, ...credentials })
    return Vue.axios.delete(`${resource}`, { ...params, ...credentials })
  },
  patch(resource, params) {
    return Vue.axios.patch(`${resource}`, params, credentials)
  },
}

export default ApiService
