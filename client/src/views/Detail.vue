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
      <div v-if="!products" class="related">
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
@import '@/scss/variables';
.cart {
  font-family: 'Quicksand';
  background-color: $content-bg-color;
  .main-cart {
    display: grid;
    width: 1370px;
    margin: 0 auto;
    .path {
      display: flex;
      background-color: white;
      width: 100%;
      height: 5rem;
      margin: $space-bar auto;
      align-items: center;
      font-size: $app-font-size;
      border-radius: calc(#{$border-radius * 2});
      a {
        text-decoration: none;
        color: $app-main-text-color;
        margin: 0 $space-bar 0 calc(#{$space-bar * 3});
        &:hover {
          cursor: pointer;
          color: $app-bg-color;
        }
      }
      span {
        margin: 0 $space-bar;
        color: $app-bg-color;
      }
    }
    .add-to-cart {
      display: flex;
      width: 100%;
      margin: 0 auto;
      border-radius: 16px;

      background-color: white;
      .img-sp {
        flex: 40%;
        padding: calc(#{$space-bar * 2});
        border-right: 1px solid rgb(214, 214, 214);
        img {
          width: 100%;
          height: 100%;
        }
      }
      .detail-product {
        flex: 60%;
        padding: 40px 20px;
        display: grid;
        h2 {
          font-weight: bold;
          font-size: calc(#{$app-font-size * 2});
          margin-bottom: $gap;
        }

        .price {
          display: grid;
          background-color: $content-bg-color;
          padding: $space-bar calc(#{$space-bar * 2});
          font-size: 14px;
          border-radius: calc(#{$border-radius * 2});
          span {
            font-weight: bold;
            font-size: calc(#{$app-font-size * 2});
            font-size: 30px;
            margin: 0;
          }
        }

        .rating {
          font-size: 18px;
        }
        span {
          margin: 10px 0;
          display: flex;
          align-items: center;
          font-size: $app-font-size;
        }
        .brand {
          font-size: $app-font-size;
          padding: 0 0 $gap;
          table {
            width: 40%;
            font-size: 14px;
            color: #333;
            tr {
              display: flex;
              padding: 8px 0;
              td {
                flex: 50%;
              }
            }
          }
        }
        .to-add {
          display: flex;
          align-items: center;
          i {
            font-size: $i-font-size;
            color: $app-bg-color;
            &:hover {
              cursor: pointer;
              transition: 0.2s;
              color: darken($app-bg-color, 15);
            }
          }
          span {
            font-size: 18px;
            font-weight: 600;
            margin: 0 $gap;
          }
        }
        button {
          width: 40%;
          cursor: pointer;
          padding: $gap;
          border: none;
          border-radius: $border-radius;
          box-shadow: 2px 2px 2px #aaa;

          font-size: 16px;
          font-family: 'Quicksand';
          background: #ff929b;
          color: white;
          margin-top: $gap;
          &:hover {
            background: #f07e88;
            cursor: pointer;
            transform: scale(1.05);
            transition: 0.2s;
          }
        }
      }
    }
    .related {
      background-color: white;
      width: 100%;
      border-radius: 16px;
      margin: calc(#{$space-bar * 2}) auto;
      padding: calc(#{$space-bar * 3}) 0;
      h4 {
        font-size: calc(#{$app-font-size * 2});
        color: $app-bg-color;
        margin-left: $space-bar;
        text-transform: uppercase;
      }
      .related-sp {
        display: grid;
        justify-content: center;
        img {
          width: 100%;
        }
        p {
          font-size: calc(#{$app-font-size * 1.2});
          color: $app-main-text-color;
        }
      }
    }
      .filter-result-body {
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        padding: 20px;
        .card {
          width: 20%;
          .card-item {
            padding: 5px $gap;
          }
          .card-body {
            margin: 16px;
            img {
              width: 100%;
              max-height: 220px;
            }
            .card-title {
              text-align: left;
              font-weight: 600;
              font-size: 16px;
            }

            .card-desc {
              margin-top: -$gap;
              opacity: 0.5;
            }

            .card-desc,
            .card-rating {
              font-size: 14px;
            }

            .card-price {
              font-size: 20px;
              font-weight: 600;
            }
          }
          &:hover {
            cursor: pointer;
            box-shadow: 0 7px 29px 0 rgba(100, 100, 111, 0.4);
          }
        }
      }
  }
}
</style>
