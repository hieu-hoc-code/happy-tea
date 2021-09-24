<template>
  <div v-if="product" class="cart">
    <div class="main-cart">
      <div class="path">
        <router-link to="/">Trang chủ</router-link>
        <p>/</p>
        <span>{{ product.name }}</span>
      </div>
      <div class="add-to-cart">
        <div class="img-sp">
          <img :src="product.image" />
        </div>
        <div class="detail-product">
          <h2>{{ product.name }}</h2>
          <div class="price">
            <span>{{ priceMod(product.price) }}</span>
            <p>{{ product.desc }}</p>
          </div>
          <span
            class="rating"
            v-html="ratingMod(product.rating, product.numOfRate)"
          ></span>
          <div class="brand">
            <table>
              <tr>
                <td>Thương Hiệu</td>
                <td>{{ brand }}</td>
              </tr>
              <tr>
                <td>Topping</td>
                <td>{{ topping }}</td>
              </tr>
              <tr>
                <td>Vận Chuyển</td>
                <td>{{ product.shipping == true ? 'có' : 'Không' }}</td>
              </tr>
              <tr>
                <td>Đã Bán</td>
                <td>113</td>
              </tr>
            </table>
          </div>
          <div class="to-add">
            <i
              class="fa fa-caret-left"
              :disabled="quantity <= 1"
              @click="subQuantity"
            ></i>
            <span>{{ quantity }}</span>
            <i class="fa fa-caret-right" @click="addQuantity"></i>
          </div>

          <button @click="addToCart">Thêm vào giỏ hàng</button>
        </div>
      </div>
      <div v-if="product" class="related">
        <h4>Sản phẩm liên quan</h4>
        <div class="related-sp">
          <img :src="related" />
          <p>Không có sản phẩm liên quan</p>
        </div>
      </div>
      <div v-else class="filter-result-body">
        <div
          class="card"
          v-for="product in products"
          :key="product.id"
          :data-id="product.id"
          @click="goDetailPage"
        >
          <div class="card-body">
            <img :src="product.image" alt="ảnh sản phẩm" />
            <div class="card-item card-title">{{ product.name }}</div>
            <div class="card-item card-desc">
              {{ descriptionMod(product.desc) }}
            </div>
            <div
              class="card-item card-rating"
              v-html="ratingMod(product.rating, product.numOfRate)"
            ></div>
            <div class="card-item card-price">
              {{ priceMod(product.price) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { UPDATE_CART_ITEM } from '@/store/actions.type'
import { mapState } from 'vuex'

import related from '@/assets/icon/empty.svg'
import {
  FETCH_A_PRODUCT,
  FETCH_CART,
  FETCH_PRODUCTS,
  FETCH_TOPPING,
  FETCH_BRAND,
} from '../store/actions.type'
import methodMixins from '../mixins/methodMixin'

export default {
  name: 'Detail',
  mixins: [methodMixins],
  data() {
    return {
      quantity: 1,
      related,
    }
  },
  computed: {
    ...mapState({
      product: state => state.product.productDetail,
      cart: state => state.cart.cart,
      brands: state => state.shop.brands,
      toppings: state => state.shop.toppings,
      products: state => state.product.products,
    }),
    brand() {
      const brands = this.brands.slice(0)
      const result = Array.prototype.find.call(
        brands,
        x => x.id === this.product.brandId,
      )
      if (result) {
        return result.name
      } else {
        return 'Không có'
      }
    },
    topping() {
      const toppings = this.toppings.slice(0)
      const result = Array.prototype.find.call(
        toppings,
        x => x.id === this.product.toppingId,
      )
      if (result) {
        return result.name
      } else {
        return 'Không có'
      }
    },
  },
  beforeMount() {
    const id = this.$route.params.id
    this.$store.dispatch(FETCH_A_PRODUCT, id)
    if (!this.brands) {
      this.$store.dispatch(FETCH_BRAND)
    }
    if (!this.toppings) {
      console.log(this.toppings)
      this.$store.dispatch(FETCH_TOPPING)
    }
    this.$store.dispatch(FETCH_PRODUCTS, { categoryId: [2] })
  },
  methods: {
    subQuantity() {
      if (this.quantity > 1) this.quantity = this.quantity - 1
    },
    addQuantity() {
      this.quantity = this.quantity + 1
    },
    async addToCart() {
      // get product id
      let productID = parseInt(this.$route.params.id)
      let quantity = this.quantity

      let isInCart = this.cart.find(item => item.product_id === productID)
      if (isInCart) quantity += isInCart.quantity

      let product = { product_id: productID, quantity: quantity }
      const isSuccess = await this.$store.dispatch(UPDATE_CART_ITEM, product)
      if (isSuccess) this.$store.dispatch(FETCH_CART)
      this.scrollToTop()
    },
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: 'smooth',
      })
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/detail.scss';
</style>
