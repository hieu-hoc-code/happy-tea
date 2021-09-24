<template>
  <div class="main">
    <div class="container">
      <div class="product-content">
        <div class="cart">
          <h2>Giỏ hàng ({{ carts.amount }}sản phẩm)</h2>
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
              v-for="product in order.orders"
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
        <div>
          <h2>Chọn hình thức thanh toán</h2>
          <label for="payment">Thanh toán khi nhận hàng</label>
          <input type="radio" name="payment" value="1" v-model="payment_id" />
          <label for="payment">Thanh toán bằng Stripe</label>
          <input
            type="radio"
            name="payment"
            value="2"
            v-model="payment_id"
            @click.prevent="noServe"
          />
        </div>
      </div>

      <div class="sidebar__right">
        <div class="address" v-if="addr.address.length > 0">
          <div class="address__title">
            <p>Địa chỉ giao hàng</p>
            <button>Sửa</button>
          </div>
          <div class="addressd__detail">
            <h2>{{ addr.default[0].name }}</h2>
            <span>
              {{
                addr.default[0].address +
                ', ' +
                addr.default[0].village +
                ', ' +
                addr.default[0].district +
                ', ' +
                addr.default[0].province
              }}
            </span>
            <span>Điện thoại: {{ addr.default[0].phone_number }}</span>
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
            <button @click="checkoutOrder">Đặt mua</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import {
  REMOVE_CART_ITEM,
  CREATE_ORDER,
  FETCH_ADDRESS,
} from '@/store/actions.type'
import sp1 from '@/assets/image/sp1.jpg'
import { CREATE_MESSAGE } from '../store/actions.type'
export default {
  name: 'Checkout',
  data() {
    return {
      sp1,
      default: null,
      payment_id: 1,
    }
  },
  computed: {
    ...mapState({
      carts: state => state.cart,
      order: state => state.order,
      addr: state => state.address,
    }),
  },
  created() {
    this.$store.dispatch(FETCH_ADDRESS)
  },
  methods: {
    async checkoutOrder() {
      console.log('vao day')
      // this.order.address_id = this.addr.default[0].id
      const isSuccess = await this.$store.dispatch(CREATE_ORDER)
      if (isSuccess) {
        let message = 'Đặt hàng thành công <3'
        this.$store.dispatch(CREATE_MESSAGE, message)
        this.carts.currentSelected.forEach(id => {
          this.$store.dispatch(REMOVE_CART_ITEM, id)
        })
        this.$router.push({ name: 'cart' })
      } else {
        let message = 'Đặt hàng thất bại :(('
        this.$store.dispatch(CREATE_MESSAGE, message)
      }
    },
    noServe() {
      let message = 'Chức năng năng này đang được update!'
      this.$store.dispatch(CREATE_MESSAGE, message)
      this.order.address_id = 0
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/order.scss';
</style>
