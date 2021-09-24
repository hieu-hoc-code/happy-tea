<template>
  <div class="catalog-container admin-item">
    <div class="title">
      <h2>Tạo danh mục</h2>
    </div>

    <div class="create">
      <div class="create-title">
        <h3>Tên danh mục</h3>
      </div>
      <input type="text" placeholder="Nhập tên danh mục" v-model="newCatalog" @keydown.enter="createCatalog"/>
      <div class="save-btn" @click="createCatalog">Lưu</div>
    </div>

    <div class="body">
      <div class="body-title">
        <h3>Tất cả danh mục</h3>
      </div>
      <input type="text" class="search" v-model="searchCatalog" placeholder="Tìm kiếm danh mục" @keydown.enter="search"/>
      <div v-for="catalog in catalogFilter" :key="catalog.id" class="catalog-item">
        <input type="text" class="edit" :data-id="catalog.id" :data-value="catalog.name" @keydown.enter="saveEdit">
        <span>{{ catalog.name }}</span>
        <button class="danger" @click="openModal(catalog.id)">Xóa</button>
        <button class="update" @click="updateCatalog">Sửa</button>
      </div>
    </div>
    <div class="modal" ref="modal">
      <div class="modal-title">Bạn muốn xóa danh mục ?</div>
      <div class="modal-body">
        <div class="close-btn" @click="deleteCatalog">Delete</div>
        <div class="confirm-btn" @click="closeModal">Cancel</div>
      </div>
    </div>
  </div>
</template>

<script>
import { POST_CATALOG, DELETE_CATALOG, UPDATE_CATALOG } from '../../store/actions.type'
import { mapState } from 'vuex'
export default {
  name: 'Catalog',
  data() {
    return {
      newCatalog: '',
      searchCatalog: '',
      catalogFilter: '',
    }
  },
  computed: {
    ...mapState({
      catalog: state => state.catalog.catalog,
    }),
  },
  beforeMount() {
    this.catalogFilter = this.$store.state.catalog.catalog
  },
  methods: {
    createCatalog() {
      this.$store.dispatch(POST_CATALOG, this.newCatalog)
      this.newCatalog = ""
    },
    openModal(id) {
      this.$refs.modal.style.display = "block"
      this.id = id
    },
    updateCatalog(e) {
      const input = e.target.closest('.catalog-item').firstChild
      input.style.display = "block"
      input.value = input.dataset.value
    },
    saveEdit(e){
      const input = e.target
      input.style.display = "none";
      this.$store.dispatch(UPDATE_CATALOG, {id: parseInt(input.dataset.id), name: input.value})
    },
    search() {
      this.catalogFilter = this.catalog
      if (!this.searchCatalog) {
         return
      }
      const result = this.catalogFilter.find(x => x.name === this.searchCatalog)
      if (result) {this.catalogFilter = [result]}
      else {this.catalogFilter = []}
    },
    closeModal() {
      this.$refs.modal.style.display = "none"
    },
    deleteCatalog() {
      this.$store.dispatch(DELETE_CATALOG, this.id)
      this.$refs.modal.style.display = "none"
    }
  },
  watch: {
    searchCatalog: function(val) {
      if (!val)
      this.search()
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../../scss/_variables.scss';
.catalog-container {
  .create {
    padding: 20px 20px;
    border: 1px solid rgb(209, 209, 209);
    border-radius: 8px;
    .create-title {
      h3 {
        font-size: 18px;
        font-weight: bolder;
      }
    }
  }

  .body{
    margin-top: 24px;
    padding: 20px 20px;
    border: 1px solid rgb(209, 209, 209);
    border-radius: 8px;
    .body-title {
       h3 {
        font-size: 18px;
        font-weight: bolder;
      }
    }

    .catalog-item {
      border: 1px solid $app-bg-color;
      border-radius: 8px;
      background-color: $content-bg-color;
      position: relative;

      padding: 10px;
      margin: 6px 0;
      input {
        display: none;
        position: absolute;
        top: -8px;
        left: 0;
        width: 100%;
        height: 100%;
        font-size: 14px;
        font-weight: 600;
      }

      span {
        font-size: 14px;
        font-weight: 600;
      }
      button{
        float: right;
        margin-top: -2px;
        padding: 4px 18px;
        margin-right: 20px;
        background: white;
        outline: none;
        border: 1px solid;
        border-radius: 8px;
        
        font-family: 'Quicksand', sans-serif;
        font-weight: 600;
        font-size: 14px;
        transition: all 0.3s linear;
        &.update{
          color:#2a8d67;
          &:hover {
            background:#2a8d67;
          }
        }
        &.danger{
          color: rgb(238, 96, 96);
          &:hover {
            background: rgb(238, 96, 96);
          }
        }
        &:hover {
          color: white;
          opacity: 0.7;
          cursor: pointer;
        }
      }
    }
  }
}
</style>
