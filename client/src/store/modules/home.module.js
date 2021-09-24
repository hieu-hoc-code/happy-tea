import {
  FETCH_CATALOG,
} from '@/store/actions.type'
import { SET_PRODUCT, SET_CATALOG } from '@/store/mutations.type'
import HomeService from '@/common/home.service'

const state = {
  allProducts: '',
  catalog: '',
}

const getters = {}

const actions = {
  async [FETCH_CATALOG]({ commit }) {
    try {
      const response = await HomeService.fetchCatalog()
      const catalog = response.data

      commit(SET_CATALOG, catalog)
    } catch (err) {
      console.log('err when fetch catalog: ', err)
    }
  },
}

const mutations = {
  [SET_PRODUCT](state, products) {
    state.allProducts = products
  },
  [SET_CATALOG](state, catalog) {
    state.catalog = catalog
  },
  [SET_PRODUCT](state, products) {
    state.allProducts = products
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
