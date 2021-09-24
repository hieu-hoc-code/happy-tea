<template>
  <div class="admin-item">
    <div class="title">
      <h2>Tất cả sản phẩm</h2>
    </div>
    <div class="body">
      <div v-for="product in products" :key="product.id" class="card">
        <div class="card-body">
          <img src="../../assets/image/sp2.jpg" alt="ảnh sản phẩm" />
          <div class="card-item card-title">{{ product.name }}</div>
          <div class="card-item card-desc">{{ descriptionMod(product.desc) }}</div>
          <div class="card-item card-rating" v-html="ratingMod(product.rating, product.numOfRate)"></div>
          <div class="card-item card-price">{{ priceMod(product.price) }}</div>
          <div class="card-btn">
            <div class="update" @click="updateProduct(product)">Sửa</div>
            <div class="destroy" @click="openModal(product.id)">Xóa</div>
          </div>
        </div>
      </div>
    </div>
    <div class="modal" ref="modal">
      <div class="modal-title">Bạn muốn xóa danh mục ?</div>
      <div class="modal-body">
        <div class="close-btn" @click="deleteProduct">Delete</div>
        <div class="confirm-btn" @click="closeModal">Cancel</div>
      </div>
    </div>
  </div>
</template>

<script>
import methodMixin from '../../mixins/methodMixin'
import {  mapState } from 'vuex'
import { DELETE_PRODUCT, FETCH_PRODUCTS } from '../../store/actions.type'
export default {
  name: 'allProduct',
  mixins: [methodMixin],
  data() {
    return {
      id: ''
    }
  },
  computed: {
    ...mapState({
      products: state => state.product.products,
    }),
  },
  beforeMount() {
    this.$store.dispatch(FETCH_PRODUCTS)
  },
  methods: {
    openModal(id) {
      this.id = id
      this.$refs.modal.style.display = "block"
    },
    closeModal() {
      this.$refs.modal.style.display = "none"
    },
    deleteProduct() {
      this.$store.dispatch(DELETE_PRODUCT, this.id)
      this.closeModal()
    },
    updateProduct(product) {
      this.$router.push({name: 'product', params: {type: "update", product: product }})
    }
  },
}
</script>

<style lang="scss" scoped>
.body {
  display: flex;
  flex-wrap: wrap;
  padding: 15px 0;
  .card {
    width: 25%;
    padding: 0 10px;
    border: 1px solid rgb(242, 236, 236);
    .card-body {
      position: relative;
      margin: 4px;
      img {
        width: 100%;
        max-height: 220px;
      }
      .card-item {
        padding-left: 8px;
      }
      .card-title {
        text-align: left;
        font-weight: 600;
        font-size: 14px;
        padding-bottom: 4px;
      }

      .card-desc {
        margin-top: -4px;
        opacity: 0.5;
      }

      .card-desc,
      .card-rating {
        font-size: 12px;
      }

      .card-price {
        font-size: 15px;
        font-weight: 600;
      }

      .card-btn {
        position: absolute;
        background: transparent !important;
        left: 4px;
        top: 4px;
        display: none;
        flex-direction: row;
        div {
          width: 40px;
          border: 1px solid;
          border-radius: 8px;
          text-align: center;
          font-weight: 600;
          margin-right: 8px;
          overflow: hidden;
          transition: all 0.3s linear;
        }
        .update {
          color: #2a8d67;
          &:hover {
            color: white;
            background-color: #2a8d67;
          }
        }
        .destroy {
          color: rgb(238, 96, 96);
          &:hover {
            color: white;
            background-color: rgb(238, 96, 96);
          }
        }
      }
    }
    &:hover .card-btn {
      display: flex;
    }
    &:hover {
      cursor: pointer;
      box-shadow: 0 7px 29px 0 rgba(100, 100, 111, 0.4);
    }
  }
}
</style>
