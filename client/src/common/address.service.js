import ApiService from './api.service'

const AddressService = {
  getAddress(user_id) {
    return ApiService.get('/api/address', { params: user_id })
  },
  createAddress(address) {
    return ApiService.post('/api/address', address)
  },
  editAddress(address) {
    return ApiService.put('/api/address', address)
  },
  deleteAddress(payload) {
    return ApiService.delete('/api/address', { params: payload })
  },
}

export default AddressService
