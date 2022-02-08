import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/consumeconfig',
    name: '消费者配置',
    meta: { title: '消费者配置', icon: 'form' },
    children: [
      {
        path: '/consumeconfig',
        name: '消费者配置',
        component: () => import('@/views/consumeconfig/index'),
        meta: { title: '消费者配置', icon: 'form' }
      }
    ]
  },
  {
    path: '/platform',
    component: Layout,
    children: [{
      path: '/platform',
      name: '平台管理',
      component: () => import('@/views/platform/index'),
      meta: { title: '平台管理', icon: 'table' }
    }]
  },
  {
    path: '/message',
    component: Layout,
    children: [{
      path: '/message',
      name: '消息列表',
      component: () => import('@/views/message/index'),
      meta: { title: '消息列表', icon: 'table' }
    }]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
