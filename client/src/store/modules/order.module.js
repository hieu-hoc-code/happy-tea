import OrderService from '../../common/order.service'
import {
  FETCH_ODER,
  CREATE_ORDER,
  FETCH_ORDER_DETAIL,
  CREATE_RATE,
  UPDATE_ORDER,
} from '../actions.type'
import {
  ADD_MSG,
  SET_ORDER,
  SET_ORDER_DETAIL,
  SET_UPDATE_ORDER,
  SET_TOTAL_ZERO,
} from '../mutations.type'

const state = {
  orders: [],
  total: 0,
  payment_id: 1,
  orderDetail: [],
  address_id: null,
  isDetailPage: false,
}

const actions = {
  async [FETCH_ODER]({ commit }) {
    const response = await OrderService.fetchOrder()
    const data = response.data
    commit(SET_ORDER, data)
  },
  async [FETCH_ORDER_DETAIL]({ commit }, id) {
    const response = await OrderService.fetchOrderDetail(id)
    const data = response.data
    commit(SET_ORDER_DETAIL, data)
  },
  async [CREATE_RATE]({ commit }, payload) {
    await OrderService.createRate(payload)
    commit()
  },
  [UPDATE_ORDER]({ commit }, ordered) {
    commit(SET_UPDATE_ORDER, ordered)
  },
  async [CREATE_ORDER]({ commit }) {
    let products = []
    if (state.orders.length > 0) {
      state.orders.forEach(product => {
        products.push({
          name: product.name,
          product_id: product.product_id,
          quantity: product.quantity,
          price: product.price,
        })
      })
      let payload = {
        total: state.total,
        payment_id: state.payment_id,
        address_id: state.address_id,
        products: products,
      }
      await OrderService.createOrder(payload)
      commit(SET_ORDER)
      commit(SET_TOTAL_ZERO)
      return true
    } else {
      let msg = 'Vui lòng chọn sản phẩm bạn muốn mua'
      commit(ADD_MSG, msg)
      return false
    }
  },
}

const mutations = {
  [SET_UPDATE_ORDER](state, payload) {
    state.orders = []
    payload.forEach(item => {
      state.orders.push(item.cart)
    })
    let tmp = 0
    state.orders.forEach(item => {
      tmp += item.price * item.quantity
    })
    state.total = tmp
  },
  [SET_ORDER](state, orders) {
    state.orders = orders
  },
  [SET_ORDER_DETAIL](state, orderDetail) {
    state.orderDetail = orderDetail
  },
  [SET_TOTAL_ZERO](state) {
    state.total = 0
  },
}

export default {
  state,
  actions,
  mutations,
}
