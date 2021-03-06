import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './scss/_reset.scss'
import ApiService from './common/api.service'
import 'font-awesome/css/font-awesome.min.css'

import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.use(Buefy)

Vue.config.productionTip = false
ApiService.init()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
