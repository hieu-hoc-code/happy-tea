<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="customer-store-pay-follow">
        <h6 class="title">Chăm sóc khách hàng</h6>
        <router-link to="#">Trung tâm trợ giúp</router-link>
        <router-link to="#">Happy Tea Blog</router-link>
        <router-link to="#">Hướng dẫn mua hàng</router-link>
        <router-link to="#">Thanh toán</router-link>
        <router-link to="#">Vận chuyển</router-link>
        <router-link to="#">Trung tâm trợ giúp</router-link>
      </div>
      <div class="customer-store-pay-follow">
        <h6 class="title">Về cửa hàng</h6>
        <router-link to="#">Giới thiệu về Happy Tea</router-link>
        <router-link to="#">Giấy phép kinh doanh</router-link>
        <router-link to="#">Nguồn góc sản phẩm</router-link>
        <router-link to="#">Flash Sales</router-link>
        <router-link to="#">Đối tác</router-link>
      </div>
      <div class="customer-store-pay-follow">
        <h6 class="title">Thanh toán</h6>
        <router-link to="#">Tại cửa hàng</router-link>
        <router-link to="#">Khi nhận hàng</router-link>
        <router-link to="#">Thanh toán trực tuyến</router-link>
        <div class="pay">
          <div class="pay-1">
            <img :src="momo" />
            <img :src="card" />
            <img :src="zalopay" />
          </div>
          <div class="pay-1">
            <img :src="mastercard" />
            <img :src="visa" />
            <img :src="paypal" />
          </div>
        </div>
      </div>
      <div class="customer-store-pay-follow">
        <h6 class="title">Liên hệ</h6>
        <router-link to="#">
          <i class="fa fa-facebook-square"> Facebook</i>
        </router-link>
        <router-link to="#">
          <i class="fa fa-instagram"> Instagram</i>
        </router-link>
        <router-link to="#">
          <i class="fa fa-twitter-square"> Twitter</i>
        </router-link>
        <router-link to="#">
          <i class="fa fa-phone-square"> 01991000</i>
        </router-link>
      </div>
    </div>
  </footer>
</template>

<script>
import momo from '@/assets/icon/momo.svg'
import card from '@/assets/icon/card.svg'
import mastercard from '@/assets/icon/mastercard.svg'
import visa from '@/assets/icon/visa.svg'
import zalopay from '@/assets/image/zalopay.png'
import paypal from '@/assets/image/paypal.png'

export default {
  data() {
    return {
      momo,
      card,
      mastercard,
      visa,
      zalopay,
      paypal,
    }
  },
}
</script>

<style lang="scss" scoped>
@import '@/scss/variables';

.footer {
  padding: $space-bar 0;
  .footer-content {
    display: flex;
    justify-content: space-between;
    width: 90%;
    margin: calc(#{$space-bar} * 1) auto;
    
    .customer-store-pay-follow {
      display: grid;
      padding: 0 $space-bar;
    }
    .title {
      color: red;
      font-size: calc(#{$app-font-size * 1.2});
      margin-bottom: $space-bar;
      color: darken($app-main-text-color, 15);
      text-transform: uppercase;
    }
    a {
      text-decoration: none;
      font-size: calc(#{$app-font-size});
      color: $app-main-text-color;
      margin-bottom: $space-bar;
    }
    i.fa-facebook-square{
      color: blue;
    }
    i.fa-instagram{
      color: rgb(187, 57, 100);
    }
    i.fa-twitter-square{
      color: rgb(13, 13, 192);
    }
    i.fa-phone-square{
      color: rgb(255, 102, 0);
    }
    i:hover{
      transform: scale(1.2);
      transition: 0.5s;
    }
    .pay {
      display: flex;
      justify-content: space-between;
      .pay-1 {
        display: grid;
        width: 50%;
      }
      img {
        width: 4rem;
      }
    }
    i {
      margin-right: $gap;
    }
  }
}

@media screen and (min-width: $screen-md) {
  .footer {
    .footer-content {
      max-width: 1170px;
    }
  }
}
</style>
