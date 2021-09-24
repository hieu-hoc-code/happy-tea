import Cookies from 'universal-cookie'
import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/Home.vue'),
  },
  {
    path: '/shop',
    name: 'shop',
    component: () => import('../views/Shop.vue'),
  },
  {
    path: '/login/',
    name: 'login',
    component: () => import('../views/Login.vue'),
    alias: '/login/admin',
  },

  {
    path: '/notif',
    name: 'notif',
    component: () => import('../components/Notif.vue'),
  },

  {
    path: '/detail/:id',
    name: 'detail',
    component: () => import('../views/Detail.vue'),
  },
  {
    path: '/cart',
    name: 'cart',
    component: () => import('../views/Cart.vue'),
  },
  {
    path: '/order',
    name: 'order',
    component: () => import('../views/Order.vue'),
  },
  {
    path: '/user',
    component: () => import('../views/User.vue'),
    children: [
      {
        path: 'me',
        name: 'me',
        component: () => import('../views/user/Me.vue'),
      },
      {
        path: 'history',
        name: 'history',
        component: () => import('../views/user/History.vue'),
        children: [
          {
            path:'order-detail',
            name:'order-detail',
            component: () => import('../views/user/OrderDetail..vue')
          }
        ]
      },
      {
        path: 'address',
        name: 'address',
        component: () => import('../views/user/Address.vue'),
      },
    ],
  },
  {
    path: '/admin',
    name: 'admin',
    component: () => import('../views/Admin.vue'),
    children: [
      {
        path: '/admin/console',
        name: 'console',
        component: () => import('../views/admin/Console.vue'),
      },
      {
        path: '/admin/product',
        name: 'product',
        component: () => import('../views/admin/Product.vue'),
      },
      {
        path: '/admin/all-products',
        name: 'products',
        component: () => import('../views/admin/AllProducts.vue'),
      },
      {
        path: '/admin/catalog',
        name: 'catalog',
        component: () => import('../views/admin/Catalog.vue'),
      },
      {
        path: '/admin/discount',
        name: 'discount',
        component: () => import('../views/admin/Discount.vue'),
      },
      {
        path: '/admin/change-password',
        name: 'changePass',
        component: () => import('../views/admin/ChangPassword.vue'),
      },
    ],
  },
  {
    path: '/search',
    name: 'search',
    component: () => import('../views/Search.vue'),
  },
  {
    path: '/:catchAll(.*)',
    name: 'notfound',
    component: () => import('../views/NotFound.vue'),
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})

router.beforeEach((to, from, next) => {
  if (to.path === '/admin') {
    const cookie = new Cookies()
    if (cookie.get('admin') !== 'true') {
      router.push({ name: 'notfound' })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
