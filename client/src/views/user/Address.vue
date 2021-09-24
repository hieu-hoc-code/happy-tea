<template>
  <div class="address">
    <div class="address-new">
      <button @click="toggle" v-if="!popupFormAddress">
        <i class="fa fa-plus-square-o"></i>
        Thêm địa chỉ mới
      </button>
    </div>

    <router-view></router-view>
    <CreateAddress v-if="popupFormAddress" @send="toggle" :prop-title="index" />
    <div v-else class="form-address" v-for="(v, index) in address" :key="v.id">
      <div class="address-me">
        <div class="name">
          <h6>{{ v.name }}</h6>
          <i class="fa fa-check-square-o">Địa chỉ mặc định: {{ v.official }}</i>
        </div>

        <button @click="updateAddress(index)">Chỉnh sửa</button>
      </div>
      <div class="where-phone">
        <h6>Địa chỉ :</h6>
        <span>
          {{
            v.address + ', ' + v.village + ', ' + v.district + ', ' + v.province
          }}
        </span>
      </div>
      <div class="where-phone">
        <h6>Điện thoại :</h6>
        <span>{{ v.phone_number }}</span>
      </div>
      <button @click="deleteAddress(v.id)">Xoa</button>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { FETCH_ADDRESS, DELETE_ADDRESS } from '@/store/actions.type'
export default {
  data() {
    return {
      popupFormAddress: false,
      index: '',
    }
  },
  created() {
    this.$store.dispatch(FETCH_ADDRESS, { user_id: this.user.userid })
  },
  computed: {
    ...mapState({
      address: state => state.address.address,
      user: state => state.auth.user,
    }),
  },
  components: {
    CreateAddress: () => import('@/components/CreateAddress.vue'),
  },
  methods: {
    toggle() {
      this.popupFormAddress = !this.popupFormAddress
    },
    updateAddress(index) {
      this.index = index + ''
      this.popupFormAddress = !this.popupFormAddress
    },
    async deleteAddress(id) {
      let payload = {
        id: id,
        user_id: this.user.userid,
      }
      const isSuccess = await this.$store.dispatch(DELETE_ADDRESS, payload)
      if (isSuccess) {
        this.$store.dispatch(FETCH_ADDRESS, { user_id: this.user.userid })
      }
    },
  },
}
</script>
<style lang="scss" scoped>
@import '@/scss/_variables';
.address {
   width: 80%;
  background-color: white;
  border-radius: 8px;
  padding: $gap;
  .address-new {
    button {
      width: 100%;
      padding: $space-bar;
      border: 1px dashed;
      cursor: pointer;
    }
  }
  .form-address {
    width: 90%;
    margin: auto;
    .address-me {
      display: flex;
      justify-content: space-between;
      margin: $space-bar 0;
      .name {
        display: flex;
        align-items: center;
        h6 {
          color: rgb(129, 129, 129);
        }
        i {
          color: green;
          margin-left: $space-bar;
        }
      }
      button {
        cursor: pointer;
      }
    }
    .where-phone {
      display: flex;
      h6 {
        margin-right: $space-bar;
        color: rgb(129, 129, 129);
      }
      span {
        color: black;
      }
    }
  }
}
</style>
