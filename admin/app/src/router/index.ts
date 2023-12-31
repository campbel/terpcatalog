import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  linkActiveClass: 'bg-gray-900',
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/strains',
      name: 'strains',
      component: () => import('../views/StrainsView.vue')
    },
    {
      path: '/strains/new',
      name: 'strains-new',
      component: () => import('../views/StrainsNewView.vue')
    },
    {
      path: '/producers',
      name: 'producers',
      component: () => import('../views/ProducersView.vue')
    },
    {
      path: '/producers/new',
      name: 'producers-new',
      component: () => import('../views/ProducersNewView.vue')
    },
  ]
})

export default router
