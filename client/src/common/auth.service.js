import ApiService from './api.service'

const AuthService = {
  login({from, ...user}) {
    return ApiService.post(`api${from}`, user)
  },
  me() {
    return ApiService.get('/api/user')
  },
  admin() {
    return ApiService.get('api/admin')
  },
  logout() {
    return ApiService.get('/api/logout')
  },
  updateUser(user) {
    return ApiService.post('/api/user', user)
  },
  uploadImage(formData) {
    return ApiService.patch('/api/upload/user', formData)
  },
}

export default AuthService
