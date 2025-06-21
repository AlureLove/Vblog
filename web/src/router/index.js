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
      redirect: '/backend/blog',
      children: [
        {
          path: 'blog',
          name: 'blog_management',
          redirect: '/backend/blog/blog_list',
          children: [
            {
              path: 'blog_list',
              name: 'backend_blog_list',
              component: () => import('../views/backend/blog/ListPage.vue'),
            },
            {
              path: 'tag_list',
              name: 'backend_tag_list',
              component: () => import('../views/backend/tag/ListPage.vue'),
            }
          ]
        },
        {
          path: 'comment',
          name: 'comment_management',
        }
      ]
    }
  ],
})

export default router
