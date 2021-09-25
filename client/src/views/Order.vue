<template>
  <div class="main">
    <div class="container">
      <div class="product-content">
        <div class="cart">
          <h2>Giỏ hàng ( {{ carts.amount }} sản phẩm )</h2>
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
              <td class="img"><img :src="product.image" alt="Hinh anh" /></td>
              <td>{{ product.name }}</td>
              <td>{{ priceMod(product.price) }}</td>
              <td>
                <span>{{ product.quantity }}</span>
              </td>
              <td>{{ priceMod(product.quantity * product.price) }}</td>
            </tr>
          </table>
        </div>
        <div class="pay">
          <h2>Chọn hình thức thanh toán :</h2>
          <div class="main-form-pay">
            <div class="form-pay">
              <input
                type="radio"
                id="payment1"
                value="1"
                v-model="payment_id"
              />
              <label for="payment1">Thanh toán khi nhận hàng</label>
            </div>
            <div class="form-pay">
              <input
                type="radio"
                id="payment2"
                value="2"
                v-model="payment_id"
                @click.prevent="noServe"
              />
              <label for="payment2">Thanh toán bằng Stripe</label>
            </div>
          </div>
        </div>
      </div>

      <div class="sidebar__right">
        <div class="address" v-if="addr.address.length > 0">
          <div class="address__title">
            <p>Giao tới :</p>
            <button>Sửa</button>
          </div>
          <div class="addressd__detail">
            <div class="who">
              <h2>{{ addr.default[0].name }}</h2>
              <span>Điện thoại : {{ addr.default[0].phone_number }}</span>
            </div>
            <span>
              Địa chỉ :
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
import methodMixins from '../mixins/methodMixin'
export default {
  name: 'Checkout',
  mixins: [methodMixins],
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
      // this.order.address_id = this.addr.default[0].id
      const isSuccess = await this.$store.dispatch(CREATE_ORDER)
      if (isSuccess) {
        let message = 'Đặt hàng thành công <3'
        this.$store.dispatch(CREATE_MESSAGE, message)
        this.carts.currentSelected.forEach(id => {
          this.$store.dispatch(REMOVE_CART_ITEM, id)
        })
        this.$router.push({ name: 'home' })
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
