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
            <span>{{ product.price }} 000 đ</span>
            <p>{{ product.desc }}</p>
          </div>
          <span>Số đánh giá: {{ product.numOfRate }}</span>
          <div class="brand">
            <table>
              <tr>
                <td>Thương hiệu</td>
                <td>Sharetea</td>
              </tr>
              <tr>
                <td>Popping</td>
                <td>Cà phê đen</td>
              </tr>
              <tr>
                <td>Vận chuyển</td>
                <td>Có</td>
              </tr>
              <tr>
                <td>Đã bán</td>
                <td>113</td>
              </tr>
            </table>
          </div>
          <div class="to-add">
            <i
              class="fa fa-caret-left"
              :disabled="quantity <= 1"
              @click="sub_quantity"
            ></i>
            <span>{{ quantity }}</span>
            <i class="fa fa-caret-right" @click="add_quantity"></i>
          </div>

          <button @click="addToCart">Thêm vào giỏ hàng</button>
        </div>
      </div>
      <div class="related">
        <h4>Sản phẩm liên quan</h4>
        <div class="related-sp">
          <img :src="related" />
          <p>Không có sản phẩm liên quan</p>
          <p>{{ cart }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { UPDATE_CART_ITEM } from '@/store/actions.type'
import { mapState } from 'vuex'
import Cookies from 'universal-cookie'

import related from '@/assets/icon/empty.svg'
import { FETCH_A_PRODUCT, FETCH_CART } from '../store/actions.type'

export default {
  name: 'Detail',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
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
    }),
  },
  created() {
    const id = this.$route.params.id
    this.$store.dispatch(FETCH_A_PRODUCT, id)
  },
  methods: {
    sub_quantity: function () {
      if (this.quantity > 1) {
        this.quantity = this.quantity - 1
      }
    },
    add_quantity: function () {
      this.quantity = this.quantity + 1
    },
    addToCart() {
      let id = parseInt(this.$route.params.id)
      let isInCart = false
      this.cart.forEach(c => {
        if (c.product_id === id) {
          let quantity = parseInt(c.quantity) + this.quantity
          let product = {
            product_id: id,
            quantity: quantity,
          }
          const isSuccess = this.$store.dispatch(UPDATE_CART_ITEM, product)
          if (isSuccess) this.$store.dispatch(FETCH_CART)
          return (isInCart = true)
        }
        if (isInCart) return
      })
      if (isInCart) return
      let product = { product_id: id, quantity: this.quantity }
      const isSuccess = this.$store.dispatch(UPDATE_CART_ITEM, product)
      if (isSuccess) this.$store.dispatch(FETCH_CART)
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/variables';
.cart {
  background-color: $content-bg-color;
  .main-cart {
    display: grid;
    width: 90%;
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
        padding: calc(#{$space-bar * 2});
        display: grid;
        h2 {
          font-size: calc(#{$app-font-size * 2});
          margin-bottom: $gap;
        }
        .price {
          display: grid;
          background-color: $content-bg-color;
          padding: $space-bar calc(#{$space-bar * 2});
          font-size: $app-font-size;
          border-radius: calc(#{$border-radius * 2});
          span {
            font-weight: bold;
            font-size: calc(#{$app-font-size * 2});
          }
        }
        span {
          display: flex;
          align-items: center;
          font-size: $app-font-size;
          margin: $space-bar 0;
        }
        .brand {
          font-size: $app-font-size;
          padding: 0 0 $gap;
          table {
            width: 100%;
            tr {
              display: flex;
              padding: $gap 0;
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
            margin: 0 $gap;
          }
        }
        button {
          width: 50%;
          text-transform: uppercase;
          cursor: pointer;
          padding: $gap;
          border: none;
          border-radius: $border-radius;
          box-shadow: 2px 2px 2px #aaa;
          background-image: linear-gradient(
            $app-button-color,
            darken($app-button-color, 15)
          );
          margin-top: $gap;
          &:hover {
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
  }
}
@media screen and (min-width: $screen-md) {
  .cart {
    .main-cart {
      max-width: 1170px;
    }
  }
}
</style>
