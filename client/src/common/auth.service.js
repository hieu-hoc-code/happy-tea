import ApiService from './api.service'

const AuthService = {
  login(user) {
    return ApiService.post(`api/login`, user)
  },
  loginAdmin(user) {
    return ApiService.post(`api/login/admin`, user)
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
  register(user) {
    return ApiService.post('/api/register', user)
  },
}

export default AuthService
