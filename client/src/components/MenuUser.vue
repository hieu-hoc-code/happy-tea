<template>
  <div class="user-menu">
    <div class="user-name">
      <img :src="user.image ? user.image : User" alt="Avatar" />
      <p>{{ user.name }}</p>
      <input value="Thay avatar" type="file" ref="myfile" class="upload-box" />
      <button @click="uploadImage">upload</button>
    </div>
    <div class="menu-content">
      <div class="router-group">
        <router-link :to="{ name: 'me' }" id="TTTK" active-class="show">
          <div class="active">Thông tin tài khoản</div>
        </router-link>
      </div>
      <div class="router-group">
        <router-link :to="{ name: 'history' }" class="LSMH" active-class="show">
          <div class="active">Lịch sử mua hàng</div>
        </router-link>
      </div>
      <div class="router-group">
        <router-link :to="{ name: 'address' }" id="SDC" active-class="show">
          <div class="active">Sổ địa chỉ của bạn</div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import User from '@/assets/icon/user.svg'
import Us from '@/assets/image/nui.jpg'
import { UPLOAD_IMAGE, CHECK_AUTHENTICATE } from '@/store/actions.type'
import { mapState } from 'vuex'
export default {
  name: 'MenuUser',
  data() {
    return {
      User,
      Us,
    }
  },
  computed: {
    ...mapState({
      user: state => state.auth.user,
    }),
  },
  methods: {
    async uploadImage() {
      let file = this.$refs.myfile.files[0]
      if (file === undefined) return

      const formData = new FormData()
      formData.append('image', file)
      formData.append('user_id', this.user.user_id)
      const isSuccess = await this.$store.dispatch(UPLOAD_IMAGE, formData)

      if (isSuccess) {
        await this.$store.dispatch(CHECK_AUTHENTICATE)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
@import '@/scss/menu_user.scss';
</style>
