import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
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
    }
  ]
})

export default router
