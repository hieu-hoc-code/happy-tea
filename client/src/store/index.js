import Vue from 'vue'
import Vuex from 'vuex'
import auth from './modules/auth.module'
import cart from './modules/cart.module'
import history from './modules/history.module'
import order from './modules/order.module'
import home from './modules/home.module'
import shop from './modules/shop.module'
import catalog from './modules/catalog.module'
import address from './modules/address.module'
import product from './modules/product.module'
Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    auth,
    cart,
    history,
    order,
    home,
    shop,
    catalog,
    address,
    product,
  },
})
