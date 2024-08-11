import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/UserStore'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'LoginView',
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/',
      name: 'HomeView',
      component: () => import('@/views/HomeView.vue'),
      meta: { requireAuth: true }
    },
    {
      path: '/service',
      name: 'ServiceView',
      component: () => import('@/views/ServiceView/ServiceView.vue'),
      redirect: { name: 'ServiceGroup' },
      meta: { requireAuth: true },
      children: [
        {
          path: 'groups',
          name: 'ServiceGroup',
          component: () => import('@/views/ServiceView/ServiceGroup.vue')
        },
        {
          path: 'group/:id',
          props: true,
          component: () => import('@/views/ServiceView/ServiceList.vue')
        }
      ]
    },
    {
      path: '/setting',
      name: 'SettingView',
      component: () => import('@/views/SettingView/SettingView.vue'),
      redirect: { name: 'SettingUser' },
      meta: { requireAuth: true },
      children: [
        {
          path: 'user-setting',
          name: 'SettingUser',
          component: () => import('@/views/SettingView/SettingUser.vue')
        },
        {
          path: 'service-setting',
          name: 'SettingService',
          component: () => import('@/views/SettingView/SettingService.vue')
        },
        {
          path: 'group-setting',
          name: 'SettingGroup',
          component: () => import('@/views/SettingView/SettingGroup.vue')
        },
        {
          path: 'authorization-setting',
          name: 'SettingAuthoriztion',
          component: () => import('@/views/SettingView/SettingAuth.vue')
        },
        {
          path: 'system-setting',
          name: 'SettingSystem',
          component: () => import('@/views/SettingView/SettingSystem.vue')
        }
      ]
    },
    {
      path: '/document',
      name: 'document',
      component: () => import('@/views/DocumentView.vue'),
      meta: { requireAuth: true }
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('@/views/HomeView.vue'),
      meta: { requireAuth: true }
    }
  ]
})

router.beforeEach(async (to) => {
  if (!to.meta.requireAuth) return

  const userStore = useUserStore()
  !userStore.IsUserLogin && (await userStore.GetUser())
  if (!userStore.IsUserLogin && to.name != 'LoginView') {
    return { name: 'LoginView' }
  }
})

export default router
