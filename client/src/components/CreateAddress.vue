<template>
  <div class="createaddress">
    <div class="can" @click="send">
      <i class="fa fa-close"></i>
    </div>
    <form @submit.prevent="update">
      <div class="information">
        <span>Họ và tên</span>
        <input type="text" v-model="form.name" required />
      </div>
      <div class="information">
        <label for="congty">Công ty :</label>
        <input
          id="congty"
          type="text"
          placeholder="company"
          v-model="form.company"
        />
      </div>
      <div class="information">
        <label for="sdt">Số điện thoại :</label>
        <input
          id="sdt"
          type="text"
          placeholder="sdt"
          v-model="form.phone_number"
          required
        />
      </div>
      <div class="information">
        <span>Tỉnh/thành phố</span>
        <select name="tinh" v-model="form.province" required>
          <option value>Chọn tỉnh / thành phố</option>
          <option value="lam dong">Lam Dong</option>
        </select>
      </div>
      <div class="information">
        <span>Quận/huyện</span>
        <select name="huyen" v-model="form.district" required>
          <option value>Chọn quận / huyện</option>
          <option value="lam ha">lam ha</option>
        </select>
      </div>
      <div class="information">
        <span>Xã</span>
        <select name="xa" v-model="form.village" required>
          <option value>Chọn xã</option>
          <option value="nam ha">nam ha</option>
        </select>
      </div>
      <div class="information">
        <label for="diachi">Địa chỉ</label>
        <input id="diachi" type="text" v-model="form.address" />
      </div>
      <div class="information">
        <span>Loại địa chỉ</span>
        <div class="where">
          <input
            type="radio"
            id="home"
            value="home"
            v-model="form.type_address"
          />
          <label for="home">Nhà ở</label>
          <input
            type="radio"
            id="workplace"
            value="company"
            v-model="form.type_address"
          />
          <label for="workplace">Cơ quan</label>
        </div>
      </div>
      <div class="check">
        <input type="checkbox" v-model="form.official" />
        <span>Đặt làm mặc định</span>
      </div>
      <div class="btn">
        <button type="submit">Cập nhật</button>
        <button>Thêm địa chỉ</button>
      </div>
    </form>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import {
  CREATE_ADDRESS,
  FETCH_ADDRESS,
  EDIT_ADDRESS,
} from '@/store/actions.type'
export default {
  data() {
    return {
      form: {
        name: 'name',
        company: '',
        phone_number: '',
        province: '',
        district: '',
        village: '',
        address: '',
        type_address: '',
        official: false,
      },
      isCreate: '',
    }
  },
  props: {
    propTitle: String,
  },
  computed: {
    ...mapState({
      user: state => state.auth.user,
      addr: state => state.address.address,
    }),
  },
  methods: {
    send() {
      this.$emit('send')
    },
    async update() {
      let isSuccess = false
      if (this.isCreate) {
        isSuccess = await this.$store.dispatch(CREATE_ADDRESS, this.form)
      }
      if (!this.isCreate) {
        let payload = {
          id: parseInt(this.addr[this.propTitle].id),
          ...this.form,
        }
        isSuccess = await this.$store.dispatch(EDIT_ADDRESS, payload)
      }
      if (isSuccess) {
        this.$store.dispatch(FETCH_ADDRESS)
        this.$emit('send')
      }
    },
    getAdd(propTitle) {
      if (propTitle !== '') {
        let ad = this.addr[propTitle]
        this.form = {
          user_id: parseInt(this.user.userid),
          name: ad.name,
          company: ad.company,
          phone_number: ad.phone_number,
          province: ad.province,
          district: ad.district,
          village: ad.village,
          address: ad.address,
          type_address: ad.type_address,
          official: ad.official,
        }
        this.isCreate = false
      } else {
        this.form = {
          user_id: parseInt(this.user.userid),
          name: 'name',
          company: '',
          phone_number: '',
          province: '',
          district: '',
          village: '',
          address: '',
          type_address: '',
          official: false,
        }
        this.isCreate = true
      }
    },
  },
  created() {
    this.getAdd(this.propTitle)
  },
}
</script>

<style lang="scss" scoped>
@import '@/scss/_variables';

.createaddress {
  input,select {
    flex: 60%;
    color: rgb(138, 138, 138);
    border-radius: 8px;
    border: 1px solid rgb(199, 195, 195);
    padding: 3px 8px;
    font-size: 14px;
    font-family: 'Quicksand';
    &:focus {
      border-color: $app-bg-color;
    }
  }
  label,span {
    margin-right: 20px;
    text-align: right;
    flex:20%;
    font-weight: 600;
  }
  .can {
    width: 90%;
    margin: $gap auto;
    display: flex;
    justify-content: flex-end;
    padding: 5px;
    background-color: rgb(245, 195, 176);
    border-radius: 8px;
    color: rgb(235, 107, 107);
    transform: translateY(10px);
    cursor: pointer;
    box-shadow: 2px 2px 3px rgb(145, 145, 145);
    transition: all .3s linear;
   &:hover {
      opacity: 0.7;
    }
    &:active {
      transform: translateY(4px);
      box-shadow: none;
    }
  }

  form {
    width: 80%;
    height: 100%;
    padding: 0 0 $space-bar 0;
    margin: auto;
    .information {
      display: flex;
      justify-content: flex-end;
      padding: 5px 0;
      input {
        &:focus {
          outline: none;
        }
        &:last-child {
          width: 50%;
        }
      }
    }
  }

  .where {
    input[id='workplace'] {
      margin-left: calc(#{$space-bar * 2});
    }
    label {
      margin-left: $gap;
    }
  }
  .check {
    span {
      margin-left: $gap;
    }
    &:hover {
      cursor: pointer;
    }
  }
  .commit {
    button[id='cancel'] {
      margin-left: $gap;
    }
  }
  .btn {
    display: flex;
    justify-content: space-around;
    margin-top: $space-bar;
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
      &:focus {
        transform: translateY(3px);
        box-shadow: none;
      }
    }
  }
}
</style>
