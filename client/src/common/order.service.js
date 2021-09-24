import ApiService from './api.service'

const OrderService = {
  createOrder(order) {
    return ApiService.post('/api/orders', order)
  },
  fetchOrder() {
    return ApiService.get('api/orders')
  },
  fetchOrderDetail(id) {
    return ApiService.get(`/api/order-detail/${id}`)
  },
  createRate(payload) {
    return ApiService.post('api/ratings', payload)
  }
}

export default OrderService
