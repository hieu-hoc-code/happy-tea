import ProductService from '../../common/products.service'
import {
  CREATE_PRODUCT,
  DELETE_PRODUCT,
  FETCH_PRODUCTS,
  SEARCH_PRODUCT,
  UPDATE_PRODUCT,
  FETCH_A_PRODUCT,
} from '../actions.type'
import {
  ADD_PRODUCT,
  SET_ERRORS,
  SET_PRODUCT,
  SET_SEARCH_PRODUCT,
  SET_PRODUCT_DETAIL,
} from '../mutations.type'
import HomeService from '../../common/home.service'
const state = {
  searchProducts: '',
  products: '',
  response: '',
  err: '',
  productDetail: '',
}

const getters = {}

const actions = {
  async [FETCH_PRODUCTS]({ commit }, payload) {
    const response = await ProductService.fetchProducts(payload)
    const data = response.data
    commit(SET_PRODUCT, data)
  },
  async [CREATE_PRODUCT]({ commit }, product) {
    try {
      const response = await ProductService.createProduct(product)
      const data = response.data
      commit(ADD_PRODUCT, data)
      return true
    } catch {
      return false
    }
  },
  async [DELETE_PRODUCT]({ commit }, id) {
    try {
      await ProductService.deleteProduct(id)
      commit(DELETE_PRODUCT, id)
      return true
    } catch (err) {
      commit(SET_ERRORS, err)
      return false
    }
  },
  async [UPDATE_PRODUCT]({ commit }, product) {
    try {
      const response = await ProductService.updateProduct(product)
      const data = response.data
      commit(UPDATE_PRODUCT, data)
      return true
    } catch (err) {
      commit(SET_ERRORS, err)
      return false
    }
  },
  async [SEARCH_PRODUCT]({ commit }, query) {
    try {
      const response = await HomeService.searchProduct(query)
      const products = response.data
      commit(SET_SEARCH_PRODUCT, products)
    } catch (err) {
      console.log('err when search product: ', err)
      return false
    }
  },
  async [FETCH_A_PRODUCT]({ commit }, id) {
    try {
      const response = await ProductService.fetchAProduct(id)
      commit(SET_PRODUCT_DETAIL, response.data)
      return true
    } catch (err) {
      return false
    }
  },
}

const mutations = {
  [SET_PRODUCT](state, products) {
    state.products = products
  },
  [ADD_PRODUCT](state, product) {
    state.products.push(product)
  },
  [DELETE_PRODUCT](state, id) {
    const result = state.products.find(x => x.id === id)
    const index = state.products.indexOf(result)
    state.products.splice(index, 1)
  },
  [UPDATE_PRODUCT](state, product) {
    const index = state.products.findIndex(x => x.id === product.id)
    state.products.splice(index, 1, product)
  },
  [SET_SEARCH_PRODUCT](state, products) {
    state.searchProducts = products
    state.products = products
  },
  [SET_PRODUCT_DETAIL](state, product) {
    state.productDetail = product
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
