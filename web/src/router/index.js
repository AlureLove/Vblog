import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/common/LoginPage.vue')
    },
    {
      path: '/backend',
      name: 'backend',
      component: () => import('../views/backend/BackendLayout.vue'),
      children: [
        {
          path: 'blog',
          name: 'backend_blog',
          component: () => import('../views/backend/blog/ListPage.vue'),
        }
      ]
    }
  ],
})

export default router
