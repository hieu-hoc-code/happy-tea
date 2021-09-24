import {
  LOGIN,
  CHECK_AUTHENTICATE,
  LOGOUT,
  UPDATE_USER,
  UPLOAD_IMAGE,
} from '../actions.type'
import { SET_USER, SET_ERRORS, CLEAR_AUTH, ADD_MSG } from '../mutations.type'
import AuthService from '@/common/auth.service'
import CartService from '@/common/cart.service'
import {
  saveToken,
  getToken,
  removeToken,
  removeCartID,
  saveUserID,
  getCartID,
  saveCartID,
} from '@/common/jwt.service'

const state = {
  isAuthenticated: false,
  errors: null,
  user: {},
  msg: [
    'Title này mấy đứa ơi,fdsaf dsafd afdsa fdsfds fdsfdsa fdsfdsa dsa vcvxaedsf Czsgw efsdxvc',
  ],
}

const getters = {
  errors(state) {
    if (!state.errors) {
      return []
    }

    return Object.keys(state.errors).map(key => {
      return `${state.errors[key]}`
    })
  },
}

const actions = {
  async [LOGIN]({ commit }, user) {
    try {
      const response = await AuthService.login(user)
      const data = response.data
      commit(SET_USER, data)
      // check cart
      let cartID = getCartID()
      if (!cartID || cartID !== data.cart_id) {
        const res = await CartService.createCart()
        saveCartID(res.data)
      }
      let str = `Hi ${data.name}, have a good day`
      commit(ADD_MSG, str)
      return true
    } catch (err) {
      commit(SET_ERRORS, err.response.data.errors)
      return false
    }
  },
  async [CHECK_AUTHENTICATE]({ commit }) {
    if (getToken()) {
      try {
        let data
        if (getToken().admin === 'true') {
          const response = await AuthService.admin()
          data = response.data
        } else {
          const response = await AuthService.me()
          data = response.data
        }
        commit(SET_USER, data)

        // check cart
        let cartID = getCartID()
        if (!cartID || cartID !== data.cart_id) {
          const res = await CartService.createCart()
          saveCartID(res.data)
        }

        return true
      } catch (err) {
        commit(CLEAR_AUTH)
      }
    } else {
      commit(CLEAR_AUTH)
    }
  },
  async [LOGOUT]({ commit }) {
    const response = await AuthService.logout()
    const data = response.data
    if (data !== '') {
      commit(CLEAR_AUTH)
      return true
    }
  },
  async [UPDATE_USER]({ commit }, user) {
    try {
      const response = await AuthService.updateUser(user)
      const data = response.data

      commit(ADD_MSG, data.msg)
      if (data.clear_token) {
        commit(CLEAR_AUTH)
      }
    } catch (err) {
      console.log(err)
      commit(ADD_MSG, err.errors)
    }
  },
  async [UPLOAD_IMAGE]({ commit }, formData) {
    try {
      const response = await AuthService.uploadImage(formData)
      const data = response.data
      commit(ADD_MSG, data.url)
      if (data.clear_token) {
        commit(CLEAR_AUTH)
      }
      return true
    } catch (err) {
      console.log(err)
      commit(ADD_MSG, err.errors)
      return false
    }
  },
}

const mutations = {
  [SET_USER](state, user) {
    state.user = user
    state.errors = null
    state.isAuthenticated = true
    saveToken(user.token, user.admin)
    saveUserID(user.user_id)
  },
  [SET_ERRORS](state, errors) {
    state.errors = errors
  },
  [CLEAR_AUTH](state) {
    state.user = {}
    state.errors = null
    state.isAuthenticated = false
    removeToken()
    removeCartID()
  },
  [ADD_MSG](state, msg) {
    state.msg.push(msg)
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
