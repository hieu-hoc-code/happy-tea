import ApiService from './api.service'

const AdminService = {
  fetchProducts(params) {
    return ApiService.get('/api/products', { params: params })
  },
  createProduct(newProduct) {
    return ApiService.post('/api/products', newProduct)
  },
  updateProduct(payload) {
    return ApiService.put('/api/products', payload)
  },
  deleteProduct(id) {
    return ApiService.delete(`/api/products/${id}`)
  },
  fetchAProduct(id) {
    return ApiService.get(`/api/products/${id}`)
  },
}

export default AdminService
