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
        <div class="profile">
          <div class="profile-icon">
            <img
              :src="user.image"
              :icon-right="active ? 'menu-up' : 'menu-down'"
            />
          </div>
          <div class="profile-name">
            <span class="profile-name__note">Tài khoản:</span>
            <span class="profile-name__name">{{ user.name }}</span>
          </div>
        </div>
        <ul class="dropdown__menu_right">
          <li class="dropdown__item">
            <router-link :to="routerControl">Tài khoản của tôi</router-link>
          </li>
          <li @click="logout" class="dropdown__item">Đăng xuất</li>
        </ul>
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

      <div class="drop-cart">
        <router-link :to="{ name: 'cart' }" class="cart">
          <div class="img-card">
            <img :src="cartIcon" />
          </div>
          <p class="corner">{{ cart.amount }}</p>
        </router-link>
        <div class="dropdown-cart">
          <div class="content-drop-cart">
            <span>aaaaaaaaaa</span>
            <span>bbbbbbbbbb</span>
            <span>cccccccccc</span>
          </div>
        </div>
      </div>
    </div>
    <div class="search-hint">
      <a href="#">Latte Sữa</a>
      <a href="#">Trà sữa hạt</a>
      <a href="#">Sữa tươi thạch</a>
      <a href="#">Sữa tươi trân châu</a>
      <a href="#">Premium latte</a>
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
      active: '',
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
    const admin = cookie.get('admin')
    admin === 'true'
      ? (this.routerControl = { path: 'admin' })
      : (this.routerControl = { name: 'me' })
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
