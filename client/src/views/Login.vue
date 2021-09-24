<template>
  <div class="main-login">
    <div class="content-login">
      <div class="img-login">
        <img :src="happytea" />
      </div>
      <div class="data-login">
        <b-tabs>
          <b-tab-item label="Đăng nhập">
            <form class="form-login" @submit.prevent="submit">
              <div class="input-group">
                <input
                  type="email"
                  v-model="email"
                  placeholder="Email"
                  name="email"
                  required
                />
              </div>
              <div class="input-group">
                <input
                  type="password"
                  v-model="password"
                  placeholder="Mật khẩu"
                  name="password"
                  required
                />
              </div>
              <div class="error" v-for="error in errors" :key="error">
                {{ error }}
              </div>
              <div class="input-group">
                <span>Quên mật khẩu ?</span>
              </div>
              <div class="input-group">
                <button class="btn-login" type="submit">Đăng nhập</button>
              </div>
              <div class="or">
                <p>Hoặc đăng nhập bằng :</p>
              </div>
              <div class="input-group-btn">
                <button class="fb" type="button"><i class="fa fa-facebook-square"></i> Facebook</button>
                <button class="gg" type="button"><i class="fa fa-google-plus-square"></i> Google</button>
              </div>
            </form>
          </b-tab-item>

          <b-tab-item label="Đăng ký">
            <form class="register">
              <div class="input-group">
                <input type="email" placeholder="Email" name="email" required />
              </div>
              <div class="input-group">
                <input
                  type="password"
                  placeholder="mật khẩu ..."
                  name="password"
                  required
                />
              </div>
              <div class="input-group">
                <input
                  type="password"
                  placeholder="Nhập lại mật khẩu ..."
                  name="cfpassword"
                  required
                />
              </div>
              <div class="input-group">
                <button class="btn">Đăng ký</button>
              </div>
            </form>
          </b-tab-item>
        </b-tabs>
      </div>
    </div>
  </div>
</template>
<script>
import happytea from '@/assets/image/happytea.jpg'
import shipper from '@/assets/image/shipper.png'
import { mapGetters } from 'vuex'
import { LOGIN, FETCH_CART } from '@/store/actions.type'
import Cookies from 'universal-cookie'
export default {
  name: 'Auth',
  data() {
    return {
      happytea,
      shipper,
      email: '',
      password: '',
    }
  },
  computed: {
    ...mapGetters(['errors']),
  },
  methods: {
    async submit() {
      const isSuccess = await this.$store.dispatch(LOGIN, {
        from: this.$route.path,
        email: this.email,
        password: this.password,
      })
      if (isSuccess) {
        const cookie = new Cookies()
        const admin  = cookie.get("admin")
        if (admin === "true") {
          this.$router.push({name: 'admin'})
          return
        }
        this.$router.push({ name: 'home' })
        this.$store.dispatch(FETCH_CART)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
@import '../scss/login.scss';
</style>
