<template>
  <div class="page-user">
    <div class="user-main">
      <MenuUser />
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Cookies from 'universal-cookie'
import MenuUser from '@/components/MenuUser.vue'
export default {
  name: 'User',
  components: {
    MenuUser,
  },
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
}
</script>

<style lang="scss" scoped>
@import '@/scss/variables.scss';

.page-user{
  background-color: $content-bg-color;
  .user-main{
    width: 80%;
    margin: 0 auto;
    display: flex;
    padding: $space-bar;
  }
}

@media screen and (min-width: $screen-md) {
  .page-user{
    .user-main{
      width: 65%;
    }
  }
}
</style>
