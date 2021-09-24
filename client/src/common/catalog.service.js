import ApiService from './api.service'

const AdminService = {
  fetchCatalog() {
    return ApiService.get('/api/catalog')
  },
  postCatalog(newCatalog) {
    return ApiService.post('/api/catalog', {name: newCatalog})
  },
  updateCatalog(payload) {
    return ApiService.put('/api/catalog', payload)
  },
  deleteCatalog(id) {
    return ApiService.delete(`/api/catalog/${id}`)
  }
}

export default AdminService