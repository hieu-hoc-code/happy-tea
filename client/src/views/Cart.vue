<template>
  <div class="main">
    <div class="container">
      <div class="product-content">
        <div class="cart">
          <h2>Giỏ hàng ({{ carts.amount }})</h2>
        </div>
        <div>
          <div>
            <div class="title product">
              <div class="buy">Mua</div>
              <div class="img">Sản phẩm</div>
              <div class="single-price">Đơn giá</div>
              <div class="quantity">Số lượng</div>
              <div class="total">thành tiền</div>
              <div class="remove">Xóa</div>
            </div>

            <div v-for="cart in carts.cart" :key="cart.id" class="product">
              <div class="buy">
                <input
                  v-model="ordered"
                  type="checkbox"
                  name="choose"
                  :value="{ cart }"
                  @change="checkboxHandler"
                />
              </div>
              <div class="img">
                <img :src="cart.image" alt="Hinh anh" />
                <span>{{ cart.name }}</span>
              </div>

              <div class="single-price">{{ priceMod(cart.price) }}</div>
              <div class="quantity">
                <i
                  class="fa fa-caret-left"
                  @click="subToCart(cart.product_id, cart.quantity, cart.id)"
                ></i>
                <span>{{ cart.quantity }}</span>
                <i
                  class="fa fa-caret-right"
                  @click="addToCart(cart.product_id, cart.quantity)"
                ></i>
              </div>
              <div class="total">
                {{ priceMod(cart.quantity * cart.price) }}
              </div>
              <div class="remove">
                <button @click="removeToCart(cart.id)">xóa</button>
              </div>
            </div>
          </div>
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
                <td>{{ priceMod(order.total) }}</td>
              </tr>
              <tr>
                <td>Giảm giá :</td>
                <td>0đ</td>
              </tr>
              <tr class="sum">
                <td>Tổng cộng :</td>
                <td>{{ priceMod(order.total) }}</td>
              </tr>
            </table>
          </div>
        </div>
        <div class="order">
          <button @click="orderHandler">Tiến hành đặt hàng</button>
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
  CREATE_MESSAGE,
  ADD_CURRENT_SELECTED,
} from '@/store/actions.type'
import sp1 from '@/assets/image/sp1.jpg'
import methodMixins from '../mixins/methodMixin'
export default {
  name: 'Cart',
  data() {
    return {
      ordered: [],
      sp1,
    }
  },
  mixins: [methodMixins],
  computed: {
    ...mapState({
      carts: state => state.cart,
      order: state => state.order,
    }),
  },
  methods: {
    ...mapActions(['updateOrder']),
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
    checkboxHandler() {
      let currentSelected = []
      this.ordered.forEach(item => {
        currentSelected.push(item.cart.id)
      })
      this.$store.dispatch(ADD_CURRENT_SELECTED, currentSelected)
      this.updateOrder(this.ordered)
    },
    orderHandler() {
      if (this.order.orders.length < 1) {
        let message = 'Bạn chưa chọn sản phẩm để mua'
        this.$store.dispatch(CREATE_MESSAGE, message)
        return
      }
      this.$router.push({ name: 'order' })
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/variables';

.main {
  font-family: 'Quicksand';
  background-color: $content-bg-color;
  .container {
    font-size: 16px;
    display: flex;
    width: 1370px;
    margin: 0 auto;
    padding: 20px 0;
    .product-content {
      flex: 70%;
      display: grid;

      .cart {
        padding: $space-bar $gap;
        margin-bottom: 10px;
        border-radius: 8px;
        background-color: white;
        h2 {
          font-weight: 600;
          font-size: 20px;
        }
      }
      .product.title {
        font-size: 16px;
        line-height: 24px;
        margin-bottom: 20px;
        height: 54px;
        .img {
          text-align: left;
        }
        .single-price,
        .quantity,
        .total,
        .remove {
          width: 15%;
          line-height: 24px;
          font-size: 16px;
          padding-top: 0;
        }
      }

      .product {
        margin-bottom: 10px;
        padding: 15px 0;
        background-color: white;
        border-radius: 8px;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        text-align: left;
        padding-left: 15px;
        box-shadow: 1px 3px 3px rgb(226, 226, 226);
        img {
          width: 106px;
        }
        .buy {
          width: 10%;
          input {
            margin-left: 20px;
            margin-top: 45px;
            width: 20px;
            height: 20px;
            &:hover {
              cursor: pointer;
            }
          }
        }

        .img {
          text-align: left;
          width: 30%;
          height: 106px;
          span {
            font-size: 16px;
            font-weight: 600;
            padding-left: 16px;
            vertical-align: top;
          }
        }

        .single-price,
        .quantity,
        .total,
        .remove {
          width: 15%;
          padding-top: 40px;
        }

        .single-price,
        .total {
          font-weight: 600;
          font-size: 18px;
          color: black;
        }

        .quantity {
          font-size: 20px;
          i {
            padding: 8px;
            &:hover {
              color:black;
              cursor: pointer;
            }
          }
        }
        .remove {
          padding-left: 40px;
          button {
            margin-left: -10px;
            font-size: 16px;
            font-weight: 600;
            color: rgb(216, 74, 74);
            border: 1px solid;
            border-radius: 8px;
            padding: 2px 20px;
            background: white;
            font-family: 'Quicksand';
            transition: all 0.3s linear;
            &:hover {
              cursor: pointer;
              color: white;
              background: rgb(216, 74, 74);
            }
          }
        }
      }
    }
    .price {
      flex: 20%;
      background-color: white;
      margin-left: 10px;
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
              color: black;
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

          font-weight: bolder;
          font-size: 16px;
          color: white;
          background-color: $app-bg-color;
          transition: background-color 0.3s linear;
          font-family: 'Quicksand';
          &:hover {
            cursor: pointer;
            background: #e5858d !important;
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
