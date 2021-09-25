import Cookies from 'universal-cookie'

const ID_TOKEN_KEY = 'jwt'
const CART_ID = 'cart_id'
const USER_ID = 'user_id'

const cookies = new Cookies()

// token user
export const saveToken = (token, admin) => {
  cookies.set(
    ID_TOKEN_KEY,
    token,
    { path: '/' },
    { expires: '24h' },
    { httpOnly: true },
  )
  cookies.set(
    "admin",
    admin,
    { path: '/' },
    { expires: '24h' },
    { httpOnly: true },
  )
}

export const getToken = () => {
  return {jwt: cookies.get('jwt'), admin: cookies.get('admin')}
}

export const removeToken = () => {
  cookies.remove('admin')
  cookies.remove('jwt')
  return
}

// cart id
export const saveCartID = token => {
  window.localStorage.setItem(CART_ID, token)
}

export const getCartID = () => {
  return window.localStorage.getItem(CART_ID)
}

export const removeCartID = () => {
  return window.localStorage.removeItem(CART_ID)
}

// token user
export const saveUserID = token => {
  cookies.set(USER_ID, token, { path: '/' }, { httpOnly: true })
}

export default {
  saveToken,
  getToken,
  removeToken,
  saveCartID,
  getCartID,
  removeCartID,
  saveUserID,
}
