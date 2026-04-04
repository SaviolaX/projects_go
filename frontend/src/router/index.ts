import LoginCmp from '@/components/Auth/LoginCmp.vue'
import RegisterCmp from '@/components/Auth/RegisterCmp.vue'
import PostDetailCmp from '@/components/Posts/PostDetailCmp.vue'
import PostFormCmp from '@/components/Posts/PostFormCmp.vue'
import PostsCmp from '@/components/Posts/PostsCmp.vue'
import { useAuthStore } from '@/stores/auth'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'Posts', component: PostsCmp, meta: { requiresAuth: false } },
    {
      path: '/posts/:id',
      name: 'PostDetail',
      component: PostDetailCmp,
      meta: { requiresAuth: false },
    },
    {
      path: '/posts/create',
      name: 'CreatePost',
      component: PostFormCmp,
      meta: { requiresAuth: true },
    },

    { path: '/auth/login', name: 'Login', component: LoginCmp, meta: { requiresAuth: false } },
    {
      path: '/auth/register',
      name: 'Register',
      component: RegisterCmp,
      meta: { requiresAuth: false },
    },
  ],
})

router.beforeEach((to, _from) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    return '/auth/login'
  }
})

export default router
