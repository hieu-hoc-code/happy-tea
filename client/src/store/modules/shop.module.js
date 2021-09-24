const API_URL = 'http://localhost:3000'
import axios from 'axios'

axios.defaults.baseURL = API_URL

const state = {
  products: '',
  catalog: '',
  brands: '',
  toppings: '',
}

const getters = {}

const actions = {
  async fetchBrands({ commit }) {
    const response = axios.get('api/brands')
    const data = (await response).data
    commit('SET_BRANDS', data)
  },
  async fetchTopping({ commit }) {
    const response = axios.get('api/toppings')
    const data = (await response).data
    commit('SET_TOPPING', data)
  },
}

const mutations = {
  SET_CATALOG(state, catalog) {
    state.catalog = catalog
  },
  SET_BRANDS(state, brands) {
    state.brands = brands
  },
  SET_TOPPING(state, topping) {
    state.toppings = topping
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
