import {
  FETCH_ADDRESS,
  CREATE_ADDRESS,
  DELETE_ADDRESS,
  EDIT_ADDRESS,
} from '@/store/actions.type.js'
import { SET_ADDRESS } from '@/store/mutations.type'
import AddressService from '@/common/address.service'

const state = {
  address: [],
  default: null,
}

const actions = {
  async [FETCH_ADDRESS]({ commit }) {
    try {
      const response = await AddressService.getAddress()

      const data = response.data
      commit(SET_ADDRESS, data)
      return true
    } catch (err) {
      console.log(err)
      return false
    }
  },
  async [CREATE_ADDRESS]({ commit }, payload) {
    try {
      const response = await AddressService.createAddress(payload)
      const data = response.data
      console.log(data)
      return true
    } catch (err) {
      console.log(err)
      commit(SET_ADDRESS)
      return false
    }
  },
  async [EDIT_ADDRESS]({ commit }, payload) {
    try {
      const response = await AddressService.editAddress(payload)
      const data = response.data
      console.log(data)
      return true
    } catch (err) {
      console.log(err)
      commit(SET_ADDRESS)
      return false
    }
  },
  async [DELETE_ADDRESS]({ commit }, payload) {
    try {
      const response = await AddressService.deleteAddress(payload)
      const data = response.data
      console.log(data)
      return true
    } catch (err) {
      console.log(err)
      commit(SET_ADDRESS)
      return false
    }
  },
}

const mutations = {
  [SET_ADDRESS](state, data) {
    state.address = data
    state.default = data.filter(item => item.official === true)
  },
}

export default {
  state,
  actions,
  mutations,
}
