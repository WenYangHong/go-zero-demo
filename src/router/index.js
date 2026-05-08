import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/DefaultLayout.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
        meta: { title: '发现民宿' },
      },
      {
        path: 'homestay/:id',
        name: 'HomestayDetail',
        component: () => import('@/views/HomestayDetail.vue'),
        meta: { title: '民宿详情' },
      },
      {
        path: 'business/:id',
        name: 'BusinessDetail',
        component: () => import('@/views/BusinessDetail.vue'),
        meta: { title: '房东主页' },
      },
      {
        path: 'comment/:homestayId',
        name: 'Comment',
        component: () => import('@/views/Comment.vue'),
        meta: { title: '评价' },
      },
      {
        path: 'order/:sn',
        name: 'OrderPay',
        component: () => import('@/views/OrderPay.vue'),
        meta: { title: '订单支付' },
      },
    ],
  },
  {
    path: '/auth',
    component: () => import('@/layouts/AuthLayout.vue'),
    children: [
      {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue'),
        meta: { title: '登录' },
      },
      {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/Register.vue'),
        meta: { title: '注册' },
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

router.beforeEach((to) => {
  document.title = `${to.meta.title || '归栖民宿'} - 归栖民宿`
})

export default router
