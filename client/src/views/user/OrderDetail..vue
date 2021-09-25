<template>
  <div v-if="isDetailPage" class="order-detail-container">
    <div class="overlay" ref="overlay" @click="closeModal"></div>
    <button class="back btn" @click="back">Quay trờ lại</button>
    <table>
      <tr class="title-order">
        <td>Sản phẩm</td>
        <td>Số lượng</td>
        <td>Đơn giá</td>
      </tr>
      <tr v-for="product in allOrder" :key="product.id" class="tb-row">
        <td>{{ product.name }}</td>
        <td>{{ product.quantity }}</td>
        <td>{{ priceMod(product.price) }}</td>
          <button class="btn rebuy">Mua lại</button>
          <button class="btn rate" @click="openModal(product.product_id)">Đánh giá</button>
      </tr>
    </table>
    <div class="modal" ref="modal">
        <div>
          <h3 class="title">Hãy cho chúng mình biết <br> đánh giá của bạn ^^</h3>
        </div>
        <i class="fas fa-star" :class="{hovered: dataRate >= 1}" data-rate="1" @mouseleave="resetDataRate" @mouseover="setDataRate" @click="rate"></i>
        <i class="fas fa-star" :class="{hovered: dataRate >= 2}" data-rate="2" @mouseleave="resetDataRate" @mouseover="setDataRate" @click="rate"></i>
        <i class="fas fa-star" :class="{hovered: dataRate >= 3}" data-rate="3" @mouseleave="resetDataRate" @mouseover="setDataRate" @click="rate"></i>
        <i class="fas fa-star" :class="{hovered: dataRate >= 4}" data-rate="4" @mouseleave="resetDataRate" @mouseover="setDataRate" @click="rate"></i>
        <i class="fas fa-star" :class="{hovered: dataRate >= 5}" data-rate="5" @mouseleave="resetDataRate" @mouseover="setDataRate" @click="rate"></i>
    </div>
  </div>
</template>

<script>
import Cookies from 'universal-cookie'
import { mapState } from 'vuex'
import methodMixins from '../../mixins/methodMixin'
import { CREATE_RATE, FETCH_ORDER_DETAIL } from '../../store/actions.type'
export default {
  name: 'order-detail',
  mixins: [methodMixins],
  data() {
    return {
      dataRate : 0,
      productId: '',
    }
  },
  computed: {
    ...mapState({
      isDetailPage: state => state.order.isDetailPage,
      allOrder: state => state.order.orderDetail,
    }),
  },
  beforeMount(){
    if (this.$route.path === "/user/history/order-detail" && this.$route.params.id) {
      this.$store.state.order.isDetailPage = true
      this.$store.dispatch(FETCH_ORDER_DETAIL, this.$route.params.id)
    } else {
      this.$router.push({name: "history"})
    }
  },
  methods: {
      back() {
          this.$store.state.order.isDetailPage = !this.$store.state.order.isDetailPage 
          this.$router.push({name: "history"})
      },
      openModal(id) {
          this.$refs.overlay.style.display = "block"
          this.$refs.modal.style.display = "block"
          this.productId = id
      },
      setDataRate(e) {
        this.dataRate = e.target.dataset.rate
      },
      resetDataRate() {
        this.dataRate = 0
      },
      rate() {
        this.$refs.modal.style.display = "none"
        this.$refs.overlay.style.display = "none"
        const cookie = new Cookies()
        let payload ={
          rate : parseInt(this.dataRate),
          productId: this.productId,
          userID: parseInt(cookie.get("user_id"))
        }
        this.$store.dispatch(CREATE_RATE, payload)
        this.dataRate = 0
      },
      closeModal() {
          this.$refs.overlay.style.display = "none"
          this.$refs.modal.style.display = "none"
      }
  }
}
</script>

<style lang="scss" scoped>
.order-detail-container {
  .overlay {
    position: fixed;
    display: none;
    width: 1920px;
    height: 1024px;
    top: 0;
    left: 0;
    z-index: 1;
    background: rgba(1,1,1,0.1);
  }
    position: relative;
  padding: 10px 20px;
  .btn {
      background: white;
      border: 1px solid;
      border-radius: 4px;
      padding: 0 10px;
      font-family: 'Quicksand';
      font-size: 14px;
      font-weight: 600;
      transition: all 0.3s linear;
      &:hover{
          cursor: pointer;
      }
  }
    .back {
        border-radius: 100px 0 0 100px;
        border-color:rgb(191, 191, 191);
        padding: 4px 10px;
        color: rgb(99, 99, 99);
        margin-bottom: 30px;
        &:hover {
            opacity: 0.8;
            color: white;
            background-color: rgb(99, 99, 99);
        }
    }

  table {
    margin: 0 auto;
    width: 100%;
    margin-left:  140px;
    .title-order {
      font-weight: bold;
    }
    tr {
      td {
        font-size: 14px;
        border: 1px solid rgb(196, 196, 196);
        text-align: center;
      }
        .btn {
            display: inline-block;
            margin-left: 20px;
        }

        .rebuy {
            color: rgb(235, 94, 94);
            &:hover {
                background-color: rgb(235, 94, 94);
                color:white;
            }
        }

        .rate {
            color: #ff9609;
             &:hover {
                background-color: #ff9609;
                color:white;
            }
        }
    }
  }
  .modal {
      width: 320px;
      height: 150px;
      text-align: center;
      background: rgb(242, 219, 219);
      color: rgb(170, 169, 169);
      box-shadow: 1px 5px 10px rgb(136, 135, 135);
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      border-radius: 16px;
      .title {
        margin-top: 30px;
        font-size: 16px;
        color: rgb(240, 141, 157);
      }
      i {
        margin-top: 20px;
        font-size: 28px;
        &:hover {
          color: rgb(235, 235, 32)
        }
      }
      .hovered {
        color: rgba(216, 216, 44, 0.8);
      }
  }
}
</style>
