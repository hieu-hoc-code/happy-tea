<template>
  <div class="main">
    <div class="container">
      <div class="product-content">
        <div class="cart">
          <h2>Giỏ hàng ({{ carts.amount }})</h2>
          {{ order }}
        </div>
        <div class="product">
          <table>
            <tr class="title">
              <td class="img">Hình ảnh</td>
              <td>Tên SP</td>
              <td>Đơn giá</td>
              <td>Số lượng</td>
              <td>thành tiền</td>
            </tr>
            <tr
              v-for="product in order.order"
              :key="product.id"
              class="product"
            >
              <td class="img"><img :src="sp1" alt="Hinh anh" /></td>
              <td>{{ product.name }}</td>
              <td>{{ product.price }}</td>
              <td>
                <span>{{ product.quantity }}</span>
              </td>
              <td>{{ product.quantity * product.price }}</td>
            </tr>
          </table>
        </div>
      </div>
      <div class="price">
        <div class="pro">
          <div class="km">
            <div class="km-happytea">
              <span>Khuyến mãi</span>
              <p>
                Có thể chọn 2
                <i class="fa fa-exclamation-circle"></i>
              </p>
            </div>

            <router-link to="#">
              <i class="fa fa-pencil-square-o"></i>
              Nhập khuyến mãi
            </router-link>
          </div>
        </div>
        <div class="bill">
          <div class="bill-content">
            <table>
              <tr>
                <td>Tạm tính :</td>
                <td>{{ order.total }}đ</td>
              </tr>
              <tr>
                <td>Giảm giá :</td>
                <td>0đ</td>
              </tr>
              <tr class="sum">
                <td>Tổng cộng :</td>
                <td>{{ order.total }}đ</td>
              </tr>
            </table>
          </div>
        </div>
        <div class="order">
          <button @click="checkout">Đặt hàng</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import {
  UPDATE_CART_ITEM,
  REMOVE_CART_ITEM,
  CREATE_ORDER
} from '@/store/actions.type'
import sp1 from '@/assets/image/sp1.jpg'
export default {
  name: 'Checkout',
  data() {
    return {
      sp1,
    }
  },
  computed: {
    ...mapState({
      carts: state => state.cart,
      order: state => state.order,
    }),
  },
  methods: {
    ...mapActions(['updateOrder', 'createOrder']),
    addToCart(product_id, quantity) {
      let product = { product_id: product_id, quantity: quantity + 1 }
      this.$store.dispatch(UPDATE_CART_ITEM, product)
    },
    subToCart(product_id, quantity, id) {
      if (quantity > 1) {
        let product = { product_id: product_id, quantity: quantity - 1 }
        this.$store.dispatch(UPDATE_CART_ITEM, product)
      } else {
        this.removeToCart(id)
      }
    },
    removeToCart(id) {
      this.$store.dispatch(REMOVE_CART_ITEM, id)
    },
    checkout() {
      this.$store.dispatch(CREATE_ORDER)
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/variables';

.main {
  background-color: $content-bg-color;
  .container {
    font-size: 16px;
    display: flex;
    width: 90%;
    margin: 0 auto;
    padding: calc(#{$space-bar * 4}) 0;
    .product-content {
      flex: 70%;
      display: grid;

      .cart {
        padding: $space-bar $gap;
        margin: $gap 0;
        border-radius: $border-radius;
        background-color: white;
      }
      .product {
        padding: $space-bar 0;
        background-color: white;
        table {
          width: 100%;
          text-align: center;
          tr {
            td {
              img {
                width: 100%;
              }
              span {
                margin: 0 5px;
              }
              i {
                cursor: pointer;
              }
            }
            td.img {
              width: 70px;
            }
          }
          tr.title {
            height: 40px;
            line-height: 40px;
            font-size: 16px;
          }
        }
      }
    }
    .price {
      flex: 25%;
      background-color: white;
      margin-left: 5%;
      min-height: 350px;
      border-radius: $border-radius;
      .pro {
        height: 30%;
        display: flex;
        justify-content: center;
        .km {
          height: 100%;
          width: 90%;
          grid: 20%;
          display: grid;
          padding: $gap;
          border-bottom: 1px solid #bbb;
          color: $app-main-text-color;
          .km-happytea {
            display: flex;
            align-items: center;
            justify-content: space-between;
            span {
              background-color: white;
              display: flex;
              border-radius: calc(#{$border-radius * 2});
            }
            p {
              color: #bbb;
            }
          }

          a {
            text-decoration: none;
            display: flex;
            align-items: center;
            color: rgb(104, 104, 255);
            i {
              margin-right: 5px;
            }
            &:hover {
              cursor: pointer;
              color: blue;
            }
          }
        }
      }
      .bill {
        height: 50%;
        .bill-content {
          padding: 5%;
          height: 90%;
          table {
            width: 100%;
            height: 100%;
            tr {
              td {
                transform: translateY(50%);
              }
            }
            tr.sum {
              font-weight: bold;
            }
          }
        }
      }
      .order {
        height: 20%;
        width: 90%;
        display: flex;
        align-items: center;
        margin: auto;
        button {
          width: 100%;
          height: 50%;
          cursor: pointer;
          border: none;
          border-radius: $border-radius;
          box-shadow: 2px 2px 2px #aaa;
          text-transform: uppercase;
          background-image: linear-gradient(
            to right,
            rgb(252, 231, 234),
            $app-bg-color
          );
          &:hover {
            font-weight: bold;
          }
          &:active {
            transform: translateY(3px);
            box-shadow: none;
          }
        }
      }
    }
  }
}

@media screen and (min-width: $screen-md) {
  .main {
    .container {
      max-width: 1170px;
    }
  }
}
</style>
