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
}

const actions = {
  async [FETCH_ADDRESS]({ commit }, user_id) {
    try {
      const response = await AddressService.getAddress(user_id)

      const data = response.data
      commit(SET_ADDRESS, data)
    } catch (err) {
      console.log(err)
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
  },
}

export default {
  state,
  actions,
  mutations,
}
