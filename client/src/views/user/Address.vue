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
          <h6>Họ và tên: <span>{{ v.name }}</span></h6>
          <i class="fa fa-check-square-o"></i>
          <span> Địa chỉ mặc định: {{ v.official }}</span>
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
      <button class="delete-btn" @click="deleteAddress(v.id)">Xóa</button>
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
  padding: 20px;
  .address-new {
    button {
      width: 100%;
      padding: 6px 0;
      border: 1px dashed;
      cursor: pointer;
    }
  }
  .form-address {
    
    border: 1px solid rgb(208, 203, 203);
    padding: 10px; 
    border-radius: 4px;
    margin: auto;
    margin-top: 20px;
    .address-me {
      display: flex;
      justify-content: space-between;
      margin: $space-bar 0;
      .name {
        display: flex;
        align-items: center;
        input {
          display: none;
        }
        h6 {
          font-weight: 600;
          color: #333;
          span{
            font-weight: 400;
          }
        }
        i {
          color: rgb(37, 181, 37);
          margin-left: 40px;
        }
        span {
          margin-left: 10px;
          font-weight: 600;
        }
      }
      button {
      cursor: pointer;
      padding: 5px $gap;
      color: rgb(74, 181, 110);
      background: white;
      border: 0.5px solid;
      border-radius: 8px;
      font-family: 'Quicksand';
      font-size: 15px;
      font-weight: 600;
      transition: all 0.3s linear;
      &:hover {
       color: white;
       background: rgb(74, 181, 110);
      }
      }
    }
    .where-phone {
      display: flex;
      h6 {
        margin-right: 5px;
         font-weight: 600;
          color: #333;
      }
      span {
        color: #333;
      }
    }
  }
  .delete-btn {
      margin-top: 10px;
       cursor: pointer;
      padding: 5px $gap;
      color: rgb(181, 74, 74);
      background: white;
      border: 0.5px solid;
      border-radius: 8px;
      font-family: 'Quicksand';
      font-size: 15px;
      font-weight: 600;
      transition: all 0.3s linear;
      &:hover {
       color: white;
       background: rgb(181, 74, 74);
      }
  }
}
</style>
