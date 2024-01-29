import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RoomsView from '../views/RoomsView.vue'

const routes = [
  {
    path: '/home',
    redirect: '/',
    component: HomeView
  },
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/rooms',
    name: 'rooms',
    component: RoomsView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
