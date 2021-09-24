<template>
  <header class="header">
    <div class="header_row_1">
      <div class="row_left">
        <router-link to="/" :class="{ active: isHome }">Trang chủ</router-link>
        <router-link :to="{ name: 'shop' }" :class="{ active: !isHome }">
          Sản phẩm
        </router-link>
      </div>

      <div v-if="isAuthenticated" class="row-right">
        <b-dropdown aria-role="list">
          <template #trigger="{ active }">
            <i
              class="fa fa-user-circle-o"
              :icon-right="active ? 'menu-up' : 'menu-down'"
            >
              <span>{{ user.name }}</span>
            </i>
          </template>
          <b-dropdown-item aria-role="listitem">
            <router-link :to="routerControl">
              Hi :
              {{ user.name }}
            </router-link>
          </b-dropdown-item>
          <b-dropdown-item aria-role="listitem">
            <p @click="logout">Đăng xuất</p>
          </b-dropdown-item>
        </b-dropdown>
      </div>

      <div v-else class="row_right">
        <router-link :to="{ name: 'login' }">Đăng nhập / Đăng ký</router-link>
      </div>
    </div>
    <div class="header_row_2">
      <router-link to="/" class="row-2-logo">
        <img src="@/assets/logo.svg" class="img-logo" />
      </router-link>
      <div class="dropdown">
        <p>
          Danh mục sản phẩm
          <i class="fa fa-sort-down"></i>
        </p>
        <div class="dropdown-content">
          <a
            v-for="catalog in catalog"
            :key="catalog.id"
            @click="getProductByCatalog(catalog.id)"
          >
            {{ catalog.name }}
          </a>
        </div>
      </div>
      <div id="search-block">
        <div class="search-box">
          <input
            v-model="query"
            type="text"
            class="search-txt"
            placeholder="Tìm sản phẩm ..."
            @keydown.enter="searchProduct"
          />
          <button @click="searchProduct" class="search-btn">
            <i class="fa fa-search"></i>
            Tìm kiếm
          </button>
        </div>
      </div>
      <router-link :to="{ name: 'cart' }" class="cart">
        <div class="img-card">
          <img :src="cartIcon" />
        </div>
        <p class="corner">{{ cart.amount }}</p>
      </router-link>
    </div>
    <div class="search-hint">
      <a href="#">Latte Sữa</a>
      <a href="#">Trà sữa hạt</a>
      <a href="#">Sữa tươi thạch</a>
      <a href="#">Sữa tươi trân châu</a>
      <a href="#">Premium Latte</a>
    </div>
  </header>
</template>

<script>
import cartIcon from '@/assets/icon/cart.svg'
import { mapState } from 'vuex'
import {
  LOGOUT,
  SEARCH_PRODUCT,
  FETCH_CATALOG,
  FETCH_PRODUCTS,
} from '@/store/actions.type'
import Cookies from 'universal-cookie'
export default {
  name: 'Header',
  data() {
    return {
      query: '',
      cartIcon,
      routerControl: '',
    }
  },
  computed: {
    ...mapState({
      catalog: state => state.catalog.catalog,
      isAuthenticated: state => state.auth.isAuthenticated,
      user: state => state.auth.user,
      cart: state => state.cart,
    }),
    isHome() {
      return this.$route.path === '/'
    },
  },
  beforeMount() {
    this.$store.dispatch(FETCH_CATALOG)
    const cookie = new Cookies()
    const admin = cookie.get("admin")
    admin === "true" ? this.routerControl = {path:'/admin'} : this.routerControl = {name:'me'}
  },
  methods: {
    async logout() {
      const isSuccess = await this.$store.dispatch(LOGOUT)
      if (isSuccess && this.$route.path !== '/') {
        this.$router.push({ name: 'home' })
      }
    },
    searchProduct() {
      if (this.$route.path === '/search') {
        this.$store.dispatch(SEARCH_PRODUCT, this.query)
      } else {
        this.$router.push({ name: 'search', params: { query: this.query } })
      }
    },
    getProductByCatalog(id) {
      if (this.$route.path === '/shop') {
        this.$store.dispatch(FETCH_PRODUCTS, { categoryId: [id] })
      } else {
        this.$router.push({ name: 'shop', params: { categoryId: [id] } })
      }
    },
  },
}
</script>
<style lang="scss">
@import '../scss/header.scss';
</style>
