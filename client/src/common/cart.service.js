import ApiService from './api.service'

const CartService = {
  fetchCart(cart_id) {
    return ApiService.get(
      '/api/cartitems',
      { params: { cart_id: cart_id } },
      { withCredentials: true },
    )
  },
  addToCart(cartItem) {
    return ApiService.post('/api/cartitems', cartItem)
  },
  createCart() {
    return ApiService.post('/api/cart')
  },
  removeCartItem(product_id) {
    return ApiService.delete('/api/cartitems', product_id)
  },
}

export default CartService
