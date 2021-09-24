<template>
  <div class="me-container">
    <form class="user-menu" @submit.prevent="update">
      <div class="user-group">
        <label for="name">Tên :</label>
        <input type="text" v-model="userName" id="name" required />
      </div>
      <div class="user-group">
        <label for="sdt">SDT :</label>
        <input type="tel" id="sdt" v-model="phoneNumber" />
      </div>
      <div class="user-group">
        <label for="email">Email :</label>
        <input type="email" id="email" v-model="user.email" disabled />
      </div>
      <div class="user-group">
        <div class="title">Giới tính :</div>
        <div class="check-sex">
          <div class="check-sex-item">
            <input
              type="radio"
              name="sex"
              id="male"
              value="male"
              v-model="sex"
            />
            <label for="male">Nam</label>
          </div>
          <div class="check-sex-item">
            <input
              type="radio"
              name="sex"
              id="female"
              value="female"
              v-model="sex"
            />
            <label for="female">Nữ</label>
          </div>
        </div>
      </div>
      <div class="user-group">
        <label for="birthday">Ngày sinh:</label>
        <input type="date" id="birthday" v-model="birthday" />
      </div>
      <div class="checkpass">
        <input type="checkbox" id="checkpass" v-model="isUpdatePassword" />
        <label class="change-pass-btn" for="checkpass">Đổi mật khẩu</label>
        <div v-if="isUpdatePassword" class="updatepass">
          <div class="input-group">
            <label for="old_password">Mật khẩu cũ :</label>
            <input
              type="password"
              id="old_password"
              placeholder="Mật khẩu cũ"
              required
              v-model="oldPassword"
            />
          </div>
          <div class="input-group">
            <label for="new-password">Mật khẩu mới :</label>
            <input
              type="password"
              id="new_password"
              placeholder="Password"
              required
              name="new_password"
              v-model="newPassword"
            />
          </div>
          <div class="input-group">
            <label for="confirm_password">nhập lại :</label>
            <input
              type="password"
              id="confirm_password"
              placeholder="Comfirm Password"
              required
              v-model="confirmPassword"
            />
          </div>
        </div>
      </div>
      <div class="save">
        <button type="submit">Cập nhật</button>
      </div>
    </form>
  </div>
</template>

<style lang="scss" scoped>
@import '@/scss/me';
</style>

<script>
import { UPDATE_USER, CHECK_AUTHENTICATE } from '@/store/actions.type'
import { mapState } from 'vuex'

export default {
  data() {
    return {
      userName: this.$store.state.auth.user.name,
      phoneNumber: this.$store.state.auth.user.phone_number,
      sex: this.$store.state.auth.user.sex,
      birthday: this.$store.state.auth.user.birthday,
      isUpdatePassword: false,
      oldPassword: '',
      newPassword: '',
      confirmPassword: '',
      errors: '',
    }
  },
  computed: {
    ...mapState({
      user: state => state.auth.user,
    }),
  },
  created() {
    if (this.user.birthday) {
      this.birthday = this.birthday.slice(0, 10)
    }
  },
  methods: {
    async update() {
      let payload = {
        name: this.userName,
        phone_number: this.phoneNumber,
        sex: this.sex,
        password: this.oldPassword,
        new_password: this.newPassword,
        confirm_password: this.confirmPassword,
      }
      if (this.birthday != '') {
        Object.assign(payload, { birthday: this.birthday + 'T00:00:00Z' })
      }
      const isSuccess = this.$store.dispatch(UPDATE_USER, payload)
      if (isSuccess) {
        this.$store.dispatch(CHECK_AUTHENTICATE)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.me-container {
  width: 80%;
  background-color: white;
  border-radius: 8px;
  padding: 10px;
  font-family: 'Quicksand', inherit;
  .user-menu {
    width: 100%;
    .user-group {
      width: 90%;
      justify-content: flex-end;
      padding: 0 40px;
      height: 36px;
      input {
        width: 80%;
        margin-left: 20px;
        border-radius: 8px;
        border: 1px solid rgb(194, 194, 194);
        padding: 8px 12px;
        font-size: 14px;
        font-family: 'Quicksand';
        color: rgb(82, 82, 82);
        &:focus {
          border: 2px solid rgb(250, 200, 209);
        }
      }
      input[type='email'] {
        opacity: 0.8;
        background-color: rgb(218, 213, 213);
      }
      label,
      .title {
        font-weight: 600;
        font-size: 16px;
      }
      .check-sex {
        display: flex;
        flex-direction: row;
        width: 630px;
        .check-sex-item {
          display: flex;
          flex-direction: row;
          width: 80px;
          input {
            width: 30px;
            margin-top: 5px;
            height: 14px;
            vertical-align: middle;
          }
        }
      }
    }

    .checkpass {
      width: 82%;
      margin-top: 24px;
      border: 1px solid rgb(204, 202, 202);
    border-radius: 8px;
    padding: 20px 30px;
      label {
        transition: all 0.3s linear;
        font-weight: 600;
      }
      #checkpass  {
        display: none;

        &:checked + label {
          color: white;
          background: rgb(230, 98, 98);
        }
      }
      .change-pass-btn {
        background: rgb(177, 174, 174);
        border: 1px solid rgb(177, 174, 174);
        border-radius: 4px;
        padding: 4px 12px;
      }

      .input-group {
        display: block;
        width: 100%;
        text-align: right;
        input{
          width: 70%;
          margin-left: 20px;
        border-radius: 8px;
        border: 1px solid rgb(194, 194, 194);
        padding: 8px 12px;
        font-size: 14px;
        font-family: 'Quicksand';
        color: rgb(82, 82, 82);
        outline: none;
        &:focus {
          border: 2px solid rgb(250, 200, 209);
        }
          
        } 
      }
    } 
  }

   .save {
     display: block;
     text-align: right;
     button {
       font-weight: bolder;
       width: unset;
       color: white;
       padding: 8px 16px;
       font-size: 24px;
       margin-right: 80px;
       text-transform: uppercase;
       background-image: none;
       background-color: rgb(226, 126, 144);
       box-shadow: 1px 5px 10px rgb(143, 141, 141);
       transition: all 0.3s linear;
       &:hover {
         background: rgb(226, 38, 72);
       }
     }
   }
}
</style>
