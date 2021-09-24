<template>
  <div class="main-content">
    <div class="content">
      <div class="slide">
        <div class="main-slide">
          <template>
            <b-carousel>
              <b-carousel-item v-for="(carousel, i) in carousels" :key="i">
                <div class="hero-body has-text-centered">
                  <img :src="carousel.color" />
                </div>
              </b-carousel-item>
            </b-carousel>
          </template>
        </div>
        <div class="advertise">
          <div class="news"><img :src="freeship" /></div>
          <div class="news"><img :src="giamgia" /></div>
        </div>
      </div>
      <div class="list-products">
        <p>
          <i class="fa fa-star"></i>
          Được yêu thích nhất
        </p>

        <div class="filter-result-body">
          <div
            class="card"
            v-for="product in favoritedProductsInPage"
            :key="product.id"
            :data-id="product.id"
            @click="goDetailPage"
          >
            <div class="card-body">
              <img src="../assets/image/sp1.jpg" alt="ảnh sản phẩm" />
              <div class="card-item card-title">{{ product.name }}</div>
              <div class="card-item card-desc">{{ descriptionMod(product.desc)}}</div>
              <div class="card-item card-rating" v-html="ratingMod(product.rating, product.numOfRate)"></div>
              <div class="card-item card-price">{{ priceMod(product.price) }}</div>
            </div>
          </div>
        </div>
        <ul class="page">
          <div
            class="previous-page page-item"
            :class="{ inActive: favoritedPage === 1 }"
            @click="prePage"
          >
            <i class="fa fa-chevron-left"></i>
          </div>
          <li
            v-for="product in favoritedProducts"
            :key="product.page"
            class="page-item"
            :class="{ active: product.page === favoritedPage }"
            @click="changeFavoritedPage(product.page)"
          >
            {{ product.page }}
          </li>
          <div
            class="next-page page-item"
            :class="{
              inActive: favoritedPage === 2,
            }"
            @click="nextPage"
          >
            <i class="fa fa-chevron-right"></i>
          </div>
        </ul>
      </div>
      <div class="list-products">
        <p>
          <i class="fa fa-star"></i>
          Sản phẩm mới nhất
        </p>

        <div class="filter-result-body">
          <div
            class="card"
            v-for="product in newestProductsInPage"
            :key="product.id"
            :data-id="product.id"
            @click="goDetailPage"
          >
            <div class="card-body">
              <img src="../assets/image/sp1.jpg" alt="ảnh sản phẩm" />
              <div class="card-item card-title">{{ product.name }}</div>
              <div class="card-item card-desc">{{ descriptionMod(product.desc)}}</div>
              <div class="card-item card-rating" v-html="ratingMod(product.rating, product.numOfRate)"></div>
              <div class="card-item card-price">{{ priceMod(product.price) }}</div>
            </div>
          </div>
        </div>

        <ul class="page">
          <div
            class="previous-page page-item"
            :class="{ inActive: newestPage === 1 }"
            @click="preNewestPage"
          >
            <i class="fa fa-chevron-left"></i>
          </div>
          <li
            v-for="product in newestProducts"
            :key="product.page"
            class="page-item"
            :class="{ active: product.page === newestPage }"
            @click="changeNewestPage(product.page)"
          >
            {{ product.page }}
          </li>
          <div
            class="next-page page-item"
            :class="{
              inActive: newestPage === 2,
            }"
            @click="nextNewestPage"
          >
            <i class="fa fa-chevron-right"></i>
          </div>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import slide2 from '@/assets/image/slide2.jpg'
import slide3 from '@/assets/image/slide3.jpg'
import freeship from '@/assets/image/free-ship.jpg'
import giamgia from '@/assets/image/giamgia20.jpg'
import sp1 from '@/assets/image/sp1.jpg'
import { mapState } from 'vuex'
import { FETCH_PRODUCTS } from '@/store/actions.type'
import methodMixin from '../mixins/methodMixin'

export default {
  name: 'Home',
  mixins: [methodMixin],
  data() {
    return {
      freeship,
      giamgia,
      sp1,
      test: '',
      favoritedPage: 1,
      newestPage: 1,
      maxProduct: 4,
      carousels: [{ color: slide2 }, { color: slide3 }],
    }
  },
  computed: {
    ...mapState({
      allProducts: state => state.product.products,
    }),
    // favorited product
    favoritedProducts() {
      const data = this.allProducts.slice()
      Array.prototype.sort.call(data, (a, b) => b.rating - a.rating)
      let fProducts = data.slice(0, 8)
      let t = this.splitPage(fProducts)
      return t
    },
    favoritedProductsInPage() {
      const data = this.favoritedProducts.find(
        e => e.page === this.favoritedPage,
      )
      if (data) {
        return data.products
      }
      return null
    },
    //newest product
    newestProducts() {
      const data = this.allProducts.slice(0)
      Array.prototype.sort.call(
        data,
        (a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt),
      )
      let nProducts = data.slice(0, 8)
      let t = this.splitPage(nProducts)
      return t
    },
    newestProductsInPage() {
      const data = this.newestProducts.find(e => e.page === this.newestPage)
      if (data) {
        return data.products
      }
      return null
    },
  },
  beforeMount() {
    this.$store.dispatch(FETCH_PRODUCTS)
  },
  methods: {
    splitPage(page) {
      const products = page
      let data = []
      let numPage = products.length / this.maxProduct
      for (let i = 0; i < numPage; i++) {
        let temp = []
        for (let j = 0; j < this.maxProduct; j++) {
          if (products.length == 0) {
            break
          }
          temp.push(products.shift())
        }
        data.push({ page: i + 1, products: temp })
      }
      return data
    },
    // favorited product page control
    changeFavoritedPage(page) {
      this.favoritedPage = page
    },
    prePage() {
      if (this.favoritedPage === 1) {
        return
      }
      this.favoritedPage--
    },
    nextPage() {
      if (this.favoritedPage === 2) {
        return
      }
      this.favoritedPage++
    },
    //newest product page control
    changeNewestPage(page) {
      this.newestPage = page
    },
    preNewestPage() {
      if (this.newestPage === 1) {
        return
      }
      this.newestPage--
    },
    nextNewestPage() {
      if (this.newestPage === 2) {
        return
      }
      this.newestPage++
    },
  },
}
</script>
<style lang="scss" scoped>
@import '../scss/home.scss';
</style>
