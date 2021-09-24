<template>
  <div class="page-user">
    <div class="user-content">
      <div class="order">
        <h2>Đơn hàng của tôi</h2>
      </div>
      <div class="order-content">
        <table v-if="!isDetailPage">
          <tr class="title-order">
            <td>Mã đơn hàng</td>
            <td>Ngày mua</td>
            <td>Tổng tiền</td>
            <td>Trạng thái đơn hàng </td>
          </tr>
          <tr v-for="order in orders" :key="order.id" class="tb-row">
            <td>{{ order.id }}</td>
            <td>{{ new Date(order.created_at) }}</td>
            <td>{{ priceMod(order.total) }}</td>
            <td>{{ orderStatusMod(order.status) }}</td>
            <button class="detail-order-btn" @click="getOrderDetail(order.id)">Chi tiết</button>
          </tr>
        </table>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import methodMixins from '../../mixins/methodMixin'
import { FETCH_ODER } from '../../store/actions.type'
export default {
  name: 'History',
  mixins: [methodMixins],
  computed: {
    ...mapState({
      isDetailPage: state => state.order.isDetailPage,
      orders: state => state.order.orders,
    }),
  },
  beforeMount() {
    if (this.$route.path === "/user/history") {
      this.$store.state.order.isDetailPage = false 
    }
    this.$store.dispatch(FETCH_ODER)
  },
  methods: {
    getOrderDetail(id) {
      this.$store.state.order.isDetailPage = !this.$store.state.orderisDetailPage
      this.$router.push({name: 'order-detail', params: {id: id}})
    },
    orderStatusMod(o) {
      let str = ''
      o === false ? (str = 'Đơn hàng đang thực hiện') : (str = 'Hoàn thành')
      return str
    },
  },
}
</script>

<style lang="scss" scoped>
@import '@/scss/variables';
.page-user {
  width: 80%;
  background-color: white;
  border-radius: 8px;
  padding: $gap;
  .user-content {
    .order {
      margin-left: $gap;
      text-transform: uppercase;
      color: $app-bg-color;
      margin-bottom: $space-bar;
      font-size: calc(#{$app-font-size * 1.5});
      h2 {
        font-weight: bold;
      }
    }
    .order-content {
      table {
        
        width: 100%;
        .title-order {
          font-weight: bold;
        }
        tr {
          td {
            font-size: 14px;
            border: 1px solid rgb(196, 196, 196);
            text-align: center;
          }
          .detail-order-btn {
            font-family: 'Quicksand';
            font-size: 14px;
            font-weight: 600;
            border:none;
            text-decoration: underline;
            background-color: white;
            padding-left: 10px;
            transition: color 0.2s linear;
            &:hover{
              cursor: pointer;
              color: rgb(204, 114, 114);
            }
          }
        }
      }
    }
  }
}
</style>
