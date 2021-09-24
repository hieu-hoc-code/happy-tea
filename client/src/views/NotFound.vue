<template lang="html">
  <div class="main">
    <div class=notfound>
      <img :src="nui" />
    </div>
    <div class="content">
      <div class="not">
        <font>404</font>
      </div>
      <h2>Không tìm thấy trang yêu cầu !</h2>
      <div class="search">
        <input type="text" placeholder="Tìm kiếm...">
        <button type="button">Tìm kiếm</button>
      </div>
      <div class="social">
        <i class="fa fa-facebook"></i>
        <i class="fa fa-twitter"></i>
        <i class="fa fa-pinterest"></i>
        <i class="fa fa-google-plus"></i>
      </div>
      <router-link to="/">Quay lại trang chủ</router-link>
    </div>
  </div>
  
</template>
<script>
import nui from '@/assets/image/nui.jpg'
export default {
  name: 'NotFound',
  data () {
    return {
      nui,
    }
  }
}
</script>
<style lang="scss">
@import '@/scss/_variables';
.main{
  display: flex;
  align-items: center;
  justify-content: center;
  .notfound{
    width: 100%;
    position: relative;
    img{
      width: 100%;
      height: 100vh;
    }
  }
  .content{
    position: absolute;
    background-color: rgb(219, 219, 219);
    z-index: 2;
    width: 600px;
    height: 600px;
    text-align: center;
    .not{
      font{
        font-size: 220px;
        color: #222225;
        text-transform: uppercase;
        font-family: passion one,cursive;
      }
    }
    h2{
      font-family: muli,sans-serif;
      font-size: 26px;
      font-weight: 400;
      text-transform: uppercase;
      color: #222225;
    }
    .search{
      margin-bottom: calc(#{$space-bar*3});
      input{
        padding: $gap;
        border: none;
        color: #fff;
        background: #222225;
      }
      button{
        padding: $gap;
        background-color: #ff00b4;
        border: none;
      }
    }
    .social{
      margin: $space-bar;
      display: flex;
      justify-content: space-around;
      i{
        background-color: #222225;
        color: white;
        width: 40px;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        &:hover{
          background-color: #ff00b4;
        }
      }
    }
    a{
      color: #222225;
    }
  }
}
</style>