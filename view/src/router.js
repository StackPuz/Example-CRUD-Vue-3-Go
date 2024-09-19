import { createWebHistory, createRouter } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/product'
  },
  {
    path: '/product',
    name: 'product',
    component: () => import('./components/product/Index.vue')
  },
  {
    path: '/product/create',
    component: () => import('./components/product/Create.vue')
  },
  {
    path: '/product/:id/',
    component: () => import('./components/product/Detail.vue')
  },
  {
    path: '/product/edit/:id/',
    component: () => import('./components/product/Edit.vue')
  },
  {
    path: '/product/delete/:id/',
    component: () => import('./components/product/Delete.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router