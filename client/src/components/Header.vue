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

      <div class="drop-cart" @mouseover="getTotal">
        <router-link :to="{ name: 'cart' }" class="cart">
          <div class="img-card">
            <img :src="cartIcon" />
          </div>
          <p class="corner">{{ productQuantity }}</p>
        </router-link>
        <div class="dropdown-cart">
          <div>
            <div class="dropdown__cart">
              Giỏ hàng ({{ cart.amount }} sản phẩm)
            </div>
            <div class="dropdown__wrapper">
              <div v-for="c in cart.cart" :key="c.id" class="dropcart__content">
                <img src="@/assets/image/nui.jpg" alt="" />
                <div class="product__detail">
                  <span class="product__detail_name">{{ c.name }}</span>
                  <div class="product__price">
                    <span class="product_price__quantity">
                      x{{ c.quantity }}
                    </span>
                    <span class="product__price_total">
                      {{ priceMod(c.quantity * c.price) }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
            <div class="dropdown__footer">
              <div class="dropcart__total">
                Tổng cộng:
                <span>{{ priceMod(cart.total) }}</span>
              </div>
              <div class="dropcart__tocart">
                <router-link :to="{ name: 'cart' }">Xem giỏ hàng</router-link>
              </div>
            </div>
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
import methodMixins from '../mixins/methodMixin'
export default {
  name: 'Header',
  mixins: [methodMixins],
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
    productQuantity() {
      const cart = this.cart.cart.slice(0)
      let quantity = 0
      Array.prototype.forEach.call(cart, x => {
        quantity += x.quantity
      })
      return quantity
    },
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
    getTotal() {
      this.$store.getters.getTotal
    },
    async logout() {
      const isSuccess = await this.$store.dispatch(LOGOUT)
      if (isSuccess) {
        this.$store.state.cart.cart = []
        if (this.$route.path !== '/') {
          this.$router.push({ name: 'home' })
        }
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
