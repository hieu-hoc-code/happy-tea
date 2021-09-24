import {
  FETCH_CART,
  UPDATE_CART_ITEM,
  REMOVE_CART_ITEM,
  ADD_CURRENT_SELECTED,
} from '../actions.type'
import {
  SET_CART,
  SET_CART_ITEM,
  REMOVE_CART,
  SET_CURRENT_SELECTED,
} from '../mutations.type'
import CartService from '@/common/cart.service'
import { getCartID } from '@/common/jwt.service'

const state = {
  cart: [],
  amount: 0,
  total: 0,
  currentSelected: [],
}
const getters = {
  getTotal: state => {
    state.total = 0
    state.cart.forEach(item => {
      state.total += item.quantity * item.price
    })
  },
}
const actions = {
  async [FETCH_CART]({ commit }) {
    const response = await CartService.fetchCart(getCartID())
    const cart = response.data
    commit(SET_CART, cart)
  },
  async [UPDATE_CART_ITEM]({ commit }, product) {
    let cartItem = {
      cart_id: parseInt(getCartID()),
      product_id: product.product_id,
      quantity: product.quantity,
    }

    await CartService.addToCart(cartItem)
    commit(SET_CART_ITEM, product)
    return true
  },
  async [REMOVE_CART_ITEM]({ commit }, id) {
    await CartService.removeCartItem(id)
    commit(REMOVE_CART, id)
  },
  [ADD_CURRENT_SELECTED]({ commit }, currentSelected) {
    commit(SET_CURRENT_SELECTED, currentSelected)
  },
}

const mutations = {
  [SET_CART](state, cart) {
    state.cart = cart
    state.amount = cart.length
  },
  [SET_CART_ITEM](state, product) {
    Array.prototype.forEach.call(state.cart, p => {
      if (p.product_id === product.product_id) {
        p.quantity = product.quantity
      }
    })
  },
  [REMOVE_CART](state, id) {
    let updateCart = state.cart.filter(p => p.id !== id)
    state.cart = updateCart
    state.amount = state.cart.length
  },
  [SET_CURRENT_SELECTED](state, currentSelected) {
    state.currentSelected = currentSelected
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
