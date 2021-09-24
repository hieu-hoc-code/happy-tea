import ApiService from './api.service'

const HomeService = {
  searchProduct(query) {
    return ApiService.get('/api/search', { params: { q: query } })
  },
  fetchProducts() {
    return ApiService.get('/api/products')
  },
  fetchCatalog() {
    return ApiService.get('/api/catalog')
  },
}

export default HomeService
