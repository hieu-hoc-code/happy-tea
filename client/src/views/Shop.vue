<template>
  <div class="shop-container">
    <div class="shop-content">
      <div class="shop-filter">
        <h3>Bộ lọc tìm kiếm</h3>

        <div class="filter-item filter-by-price">
          <div class="filter-title" @click="choiceFilter">
            <i class="fas fa-search-dollar"></i>
            <span>Khoảng giá</span>
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <input
              v-model="filter.priceRange"
              type="range"
              min="0"
              max="100000"
            />
            <span>{{ priceRangeMod }}</span>
          </div>
        </div>

        <div class="filter-item filter-by-catalog">
          <div class="filter-title" @click="choiceFilter">
            <i class="fas fa-list"></i>
            Danh mục sản phẩm
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <div v-for="catalog in catalog" :key="catalog.id" class="check-box">
              <input
                type="checkbox"
                :value="catalog.id"
                v-model="filter.categoryId"
              />
              <span>{{ catalog.name }}</span>
            </div>
          </div>
        </div>

        <div class="filter-item filter-by-rate">
          <div class="filter-title" @click="choiceFilter">
            <i class="far fa-star"></i>
            Đánh giá
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <div @click="filterByRate(5)">
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
            </div>
            <div @click="filterByRate(4)">
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star"></i>
              trở lên
            </div>
            <div @click="filterByRate(3)">
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              trở lên
            </div>
            <div @click="filterByRate(2)">
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              trở lên
            </div>
            <div @click="filterByRate(1)">
              <i class="fas fa-star yellow-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              <i class="fas fa-star"></i>
              trở lên
            </div>
          </div>
        </div>

        <div class="filter-item filter-shipping">
          <div class="filter-title" @click="choiceFilter">
            <i class="fas fa-list"></i>
            Thương hiệu
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <div v-for="brand in brands" :key="brand.id" class="check-box">
              <input
                type="checkbox"
                :value="brand.id"
                v-model="filter.brandId"
              />
              <span>{{ brand.name }}</span>
            </div>
          </div>
        </div>

        <div class="filter-item filter-shipping">
          <div class="filter-title" @click="choiceFilter">
            <i class="fas fa-list"></i>
            Topping
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <div
              v-for="topping in toppings"
              :key="topping.id"
              class="check-box"
            >
              <input
                type="checkbox"
                :value="topping.id"
                v-model="filter.toppingId"
              />
              <span>{{ topping.name }}</span>
            </div>
          </div>
        </div>

        <div class="filter-item filter-shipping">
          <div class="filter-title" @click="choiceFilter">
            <i class="fas fa-shipping-fast"></i>
            Giao hàng
            <i class="submenu-arrow"></i>
          </div>
          <div class="filter-body">
            <div class="check-box">
              <input type="radio" :value="true" v-model="filter.shipping" />
              <span>Có</span>
            </div>
            <div class="check-box">
              <input type="radio" :value="false" v-model="filter.shipping" />
              <span>Không</span>
            </div>
          </div>
        </div>

        <div class="clear-filter-btn">
          <button @click="clearFilter">Xóa tất cả</button>
        </div>
      </div>
      <div class="filter-result">
        <h3>SẢN PHẨM</h3>
        <div v-if="!filterResult || filterResult.length === 0" class="not-found">
          <img
            src="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
            alt="Empty"
          />
          <h4>Không tìm thấy sản phẩm!</h4>
        </div>

        <div v-else class="filter-result-body">
          <div class="card" v-for="product in filterResult" :key="product.id" :data-id="product.id" @click="goDetailPage">
            <div class="card-body">
               <img src="../assets/image/sp1.jpg" alt="ảnh sản phẩm" />
              <div class="card-item card-title">{{ product.name }}</div>
              <div class="card-item card-desc">{{ descriptionMod(product.desc)}}</div>
              <div class="card-item card-rating" v-html="ratingMod(product.rating, product.numOfRate)"></div>
              <div class="card-item card-price">{{ priceMod(product.price) }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { debounce } from 'debounce'
import { mapState } from 'vuex'
import { FETCH_PRODUCTS } from '@/store/actions.type'
import methodMixin from '../mixins/methodMixin'
export default {
  name: 'shop',
  mixins: [methodMixin],
  data() {
    return {
      filter: {
        priceRange: '',
        categoryId: [],
        rating: '',
        brandId: [],
        toppingId: [],
        shipping: '',
      },
    }
  },
  computed: {
    ...mapState({
      products: state => state.product.products,
      catalog: state => state.catalog.catalog,
      brands: state => state.shop.brands,
      toppings: state => state.shop.toppings,
    }),
    filterResult() {
      if (this.products) {
        const data = this.products.slice(0)
        return data
      }
      return null
    },
    priceRangeMod() {
      let temp = this.filter.priceRange / 1000
      if (temp % 1 === 0) {
        return `${temp}.000 đ`
      } else {
        return `${temp} đ`
      }
    },
  },
  beforeMount() {
    console.log("before")
    this.$store.dispatch(FETCH_PRODUCTS, this.$route.params)
    this.$store.dispatch('fetchBrands')
    this.$store.dispatch('fetchTopping')
  },
  methods: {
    choiceFilter(e) {
      const target = e.target.closest('.filter-item')
      const isActive = target.classList.contains('in-active')
      target.classList.toggle('in-active', !isActive)
    },
    filterByRate(rate) {
      this.filter.rating = rate
    },
    clearFilter() {
      this.filter = {
        priceRange: '',
        categoryId: [],
        rating: '',
        brandId: [],
        toppingId: [],
        shipping: '',
      }
    },
  },
  watch: {
    filter: {
        handler: debounce(function(filter){
          console.log("fetch watch")
          this.$store.dispatch(FETCH_PRODUCTS, filter)
        },800),
        deep: true,
      },
  },
}
</script>

<style lang="scss" scoped>
@import '../scss/shop.scss';
</style>
