<template>
  <div class="admin-item">
    <div class="title">
      <h2>Tạo sản phẩm</h2>
    </div>
    <div class="upload-image body">
      <div class="title">Chọn ảnh sản phẩm</div>
      <button @click="uploadImgLink">
        Tải ảnh lên
      </button>
      <input
        type="file"
        ref="fileInput"
        placeholder="Chọn file"
        @change="previewImg"
      />
      <div class="img" ref="img" :class="{hide: this.image === '' }">
        <div class="remove-btn" @click="removeImg">
          <p>x</p>
        </div>
        <img ref="preview" :src="this.image" alt="Ảnh" />
      </div>
    </div>

    <div class="body">
      <div class="title">Nhập thông tin sản phẩm</div>
      <div class="input">
        <span>Tên:</span>
        <input type="text" v-model="newProduct.name" />
      </div>
      <div class="input">
        <span>Mô tả:</span>
        <input type="text" v-model="newProduct.desc" />
      </div>
      <div class="input">
        <span>Giá tiền:</span>
        <input type="text" v-model="newProduct.price" />
      </div>

      <div class="input">
        <span>Vận chuyển:</span>
        <select v-model="newProduct.shipping">
          <option disabled value="">Vui lòng chọn</option>
          <option :value="true">Có</option>
          <option :value="false">Không</option>
        </select>
      </div>

      <div class="input">
        <span>Topping:</span>
        <select v-model="newProduct.toppingId">
          <option disabled value="">Vui lòng chọn</option>
          <option
            v-for="topping in toppings"
            :key="topping.id"
            :value="topping.id"
          >
            {{ topping.name }}
          </option>
        </select>
      </div>

      <div class="input">
        <span>Thương hiệu:</span>
        <select v-model="newProduct.brandId">
          <option disabled value="">Vui lòng chọn</option>
          <option v-for="brand in brands" :key="brand.id" :value="brand.id">
            {{ brand.name }}
          </option>
        </select>
      </div>

      <div class="input">
        <span>Danh mục:</span>
        <select v-model="newProduct.categoryId">
          <option disabled value="">Vui lòng chọn</option>
          <option
            v-for="catalog in catalog"
            :key="catalog.id"
            :value="catalog.id"
          >
            {{ catalog.name }}
          </option>
        </select>
      </div>
      <div class="save save-btn" @click="createProduct">Lưu</div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { CREATE_PRODUCT, UPDATE_PRODUCT } from '../../store/actions.type'
export default {
  name: 'Product',
  data() {
    return {
      newProduct: {
        name: '',
        desc: '',
        price: '',
        shipping: '',
        toppingId: '',
        brandId: '',
        categoryId: '',
      },
      image: '',
      form: new FormData(),
    }
  },
  computed: {
    ...mapState({
      catalog: state => state.catalog.catalog,
      brands: state => state.shop.brands,
      toppings: state => state.shop.toppings,
    }),
  },
  beforeMount() {
    if (this.$route.params.product) {
      this.newProduct = this.$route.params.product
    }
    if (this.newProduct.image) {
       this.image = this.newProduct.image
    }
    this.$store.dispatch('fetchBrands')
    this.$store.dispatch('fetchTopping')
  },
  methods: {
    uploadImgLink() {
      const input = this.$refs.fileInput
      input.click()
    },
    async createProduct() {
      let isSuccess
      this.newProduct.price = parseFloat(this.newProduct.price)

      const type = this.$route.params.type
      this.form.append('image', this.$refs.fileInput.files[0])
      this.form.append('newProduct', JSON.stringify(this.newProduct))
      if (type == 'update') {
        isSuccess = await this.$store.dispatch(UPDATE_PRODUCT, this.form)
      } else {
        isSuccess = await this.$store.dispatch(CREATE_PRODUCT, this.form)
      }
      if (isSuccess) {
        this.newProduct = {
          name: '',
          desc: '',
          price: '',
          shipping: '',
          toppingId: '',
          brandId: '',
          categoryId: '',
        }
        this.removeImg()
      }
    },
    previewImg(e) {
      const [file] = e.target.files
      if (file) {
        this.image = URL.createObjectURL(file)
      }
    },
    removeImg() {
      this.image = ''
      this.$refs.fileInput.value = ""
    },
  },
}
</script>

<style lang="scss" scoped>
@import '../../scss/_variables.scss';
.admin-item {
  .upload-image {
    .hide {
      display: none;
    }
    button {
      font-family: 'Quicksand', sans-serif;
      padding: 4px 16px;
      font-size: 13px;
      font-weight: 600;
      background: white;
      border: solid 1px rgb(172, 172, 172);
      border-radius: 8px;
      color: #333;
      &:hover {
        cursor: pointer;
         opacity: 0.9;
      }
    }
    input {
      display: none;
    }
    .img {
      position: relative;
      margin-top: 10px;
      padding: 10px;
      width: 120px;
      border: 1px solid rgb(199, 196, 196);
      border-radius: 16px;
      .remove-btn {
        position: absolute;
        top: -10px;
        right: -10px;
        font-weight: bolder;
        color: white;
        width: 20px;
        font-size: 14px;
        height: 20px;
        background: rgb(252, 90, 90);
        border-radius: 9999999px;
        text-align: center;
        vertical-align: middle;
        p {
          margin-top: -1px;
        }
        &:hover {
          cursor: pointer;
          opacity: 0.8;
        }
      }
      img {
        margin: 0 auto;
        border-radius: 8px;
        width: 100px;
      }
    }
  }

  .body {
    .input {
      display: flex;
      justify-content: flex-end;
      border: none;
      margin: 10px 0;
      span {
        font-size: 14px;
        font-weight: 600;
        margin-right: 20px;
      }
      input,
      select {
        width: 700px;
      }
      &:active {
        box-shadow: none;
      }
    }
    .save-btn {
      width: 80px;
      margin-left: calc(100% - 90px);
    }
  }
}
</style>
