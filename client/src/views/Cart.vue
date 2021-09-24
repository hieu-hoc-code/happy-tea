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
                <img :src="sp1" alt="Hinh anh" />
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
import { mapState } from 'vuex'
import {
  UPDATE_CART_ITEM,
  REMOVE_CART_ITEM,
  CREATE_MESSAGE,
  ADD_CURRENT_SELECTED,
  UPDATE_ORDER,
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
      this.$store.dispatch(UPDATE_ORDER, this.ordered)
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
@import '@/scss/cart.scss';
</style>
