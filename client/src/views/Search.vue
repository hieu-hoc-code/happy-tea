<template>
    <div class="search-container">
        <div class="search-content">
            <div class="search-count">
                <h3 class="title">Kết quả: {{productsCount}} sản phẩm</h3>
            </div>
            <div class="search-result">
                <div v-for="product in products" :key="product.id" class="card">
        <div class="card-body">
          <img src="../assets/image/sp2.jpg" alt="ảnh sản phẩm" />
          <div class="card-item card-title">{{ product.name }}</div>
          <div class="card-item card-desc">{{ descriptionMod(product.desc) }}</div>
          <div class="card-item card-rating" v-html="ratingMod(product.rating, product.numOfRate)"></div>
          <div class="card-item card-price">{{ priceMod(product.price) }}</div>
        </div>
      </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from 'vuex'
import methodMixins from '../mixins/methodMixin'
import {  SEARCH_PRODUCT } from '../store/actions.type'
export default {
    mixins: [methodMixins],
    computed: {
        ...mapState({
            products: state => state.product.searchProducts
        }),
        productsCount() {
            if (this.products) {
                return this.products.length
            } else {
                return 0
            }
        }
    },
    beforeMount() {
        this.$store.dispatch(SEARCH_PRODUCT, this.$route.params.query)
    },
    methods: {
    }
}
</script>

<style lang="scss" scoped>
@import '../scss/_variables.scss';
    .search-container {
        background-color: $content-bg-color;
        min-height: 600px;
        padding: 20px 0;
        .search-content{
            max-width: 1370px;
            background: white;
            min-height: 600px;
            margin: auto;
            .search-count{
                h3{
                    color: $app-bg-color;
                    margin: 0 20px;
                    padding: 30px 20px;
                    border-bottom: 1px solid rgb(220, 220, 220);
                }
            }
        }
         .search-result {
            display: flex;
            flex-wrap: wrap;
            padding: 20px 40px;
            .card {
                width: 25%;
                position: inherit;
                .card-item {
                    padding: 5px $gap;
                }
                .card-body {
                    margin: 16px;
                    img{
                        width: 100%;
                        max-height: auto;
                    }
                    .card-title {
                        text-align: left;
                        font-weight: 600;
                        font-size: 16px;
                    }

                    .card-desc {
                        margin-top: -$gap;
                        opacity: 0.5;
                    }
                 
                    .card-desc, .card-rating {
                        font-size: 14px;
                    }

                    .card-price {
                        font-size: 20px;
                        font-weight: 600;
                    }
                }
                &:hover {
                    cursor: pointer;
                    box-shadow: 0 7px 29px 0 rgba(100, 100, 111, 0.4);
                }
            }
    }
}

</style>