<template>
  <div class="main-app-content">
    <Header v-if="isVisible" />
    <router-view />
    <Footer v-if="isVisible" />
    <Notif />
  </div>
</template>

<script>
import { CHECK_AUTHENTICATE, FETCH_CART } from '@/store/actions.type'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import Notif from './components/Notif.vue'
export default {
  name: 'App',
  data() {
    return {
      isVisible: true,
    }
  },
  components: {
    Header,
    Footer,
    Notif,
  },
  beforeCreate() {
    const isSuccess = this.$store.dispatch(CHECK_AUTHENTICATE)
    if (isSuccess) {
      this.$store.dispatch(FETCH_CART)
    }
  },
  watch: {
    $route() {
      if (this.$route.name === 'notfound') {
        this.isVisible = false
      } else {
        this.isVisible = true
      }
    },
  },
}
</script>
<style lang="scss">
body {
  font-family: 'Quicksand', sans-serif !important;
}
.main-app-content {
  position: relative;
}
</style>
