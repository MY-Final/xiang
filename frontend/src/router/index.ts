import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/auth/LoginView.vue'),
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/views/auth/RegisterView.vue'),
    },
    {
      path: '/posts',
      name: 'posts',
      component: () => import('@/views/post/PostList.vue'),
    },
    {
      path: '/post/:id',
      name: 'post-detail',
      component: () => import('@/views/post/PostDetail.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/AboutView.vue'),
    },
  ],
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()

  // 需要认证的路由
  const requiresAuth = ['/admin', '/dashboard']
  const needsAuth = requiresAuth.some((path) => to.path.startsWith(path))

  if (needsAuth && !userStore.isLoggedIn) {
    next({ name: 'login', query: { redirect: to.fullPath } })
  } else if (to.name === 'login' && userStore.isLoggedIn) {
    // 已登录用户访问登录页，重定向到首页
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router
