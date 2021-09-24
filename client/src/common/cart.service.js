import ApiService from './api.service'

const CartService = {
  fetchCart(id) {
    return ApiService.get(`/api/cartitems/${id}`)
  },
  addToCart(payload) {
    return ApiService.post('/api/cartitems', payload)
  },
  createCart() {
    return ApiService.post('/api/cart')
  },
  removeCartItem(id) {
    return ApiService.delete(`/api/cartitems/${id}`)
  },
}

export default CartService
