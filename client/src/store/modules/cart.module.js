import {
  FETCH_CART,
  UPDATE_CART_ITEM,
  REMOVE_CART_ITEM,
} from '@/store/actions.type'
import { SET_CART, SET_CART_ITEM, REMOVE_CART } from '@/store/mutations.type'
import CartService from '@/common/cart.service'
import { getCartID } from '@/common/jwt.service'

const state = {
  cart: [],
  amount: 0,
  total: 0,
}
const getters = {
  getTotal: state => {
    state.cart.forEach(item => {
      state.total += item.quantity * item.price
    })
  },
}
const actions = {
  async [FETCH_CART]({ commit }) {
    let cartID = getCartID()
    if (cartID) {
      try {
        const response = await CartService.fetchCart(cartID)
        const cart = response.data
        commit(SET_CART, cart)
      } catch (err) {
        console.log('err when fetch cart: ', err)
      }
    } else {
      console.log('khong')
    }
  },
  async [UPDATE_CART_ITEM]({ commit }, product) {
    let cartID = getCartID()
    if (cartID) {
      try {
        let cartItem = {
          cart_id: parseInt(cartID),
          product_id: product.product_id,
          quantity: product.quantity,
        }

        await CartService.addToCart(cartItem) //axios.post(

        commit(SET_CART_ITEM, product)
        return true
      } catch (err) {
        console.log('err when add cart: ', err)
        return false
      }
    } else {
      console.log('khong co cart')
    }
  },
  async [REMOVE_CART_ITEM]({ commit }, id) {
    let cartID = getCartID()
    if (cartID) {
      try {
        await CartService.removeCartItem({ params: { id: id } })

        commit(REMOVE_CART, id)
        return true
      } catch (err) {
        console.log('err when add cart: ', err)
        return false
      }
    } else {
      console.log('khong co cart')
    }
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
}

export default {
  state,
  getters,
  actions,
  mutations,
}
