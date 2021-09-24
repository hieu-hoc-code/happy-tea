<template>
  <div class="admin-container">
    <div class="console">
      <router-link
        v-for="(console, index) in consoleItem"
        :key="index"
        class="console-item"
        active-class="active"
        :to="{name: `${console.slug}`}"
      >
        {{ console.name }}
      </router-link>
    </div>

    <div class="console-content">
      <div class="console-content-title"></div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  data() {
    return {
      consoleItem: [
        {name :'Bảng điều khiển', slug: "console"},
        {name :'Sản phẩm', slug: "product"},
        {name :'Tất cả sản phẩm', slug: "products"},
        {name :'Danh mục', slug: "catalog"},
        {name :'Mã giảm giá', slug: "discount"},
        {name :'Đổi mật khẩu', slug: "changePass"},
      ],
    }
  },
  computed: {
    ...mapState({
      user: state => state.auth.user
    })
  }
}
</script>

<style lang="scss">
 @import '../scss/admin.scss'
</style>