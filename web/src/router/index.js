import { createRouter, createWebHashHistory } from 'vue-router'
import indexView from '../views/indexView.vue'
import updatePassword from '../views/updatePassword.vue'
import deviceInfo from '@/components/device'
import group from '@/components/group'
import linkale from '@/components/linkage'
import types from '@/components/types'
const routes = [
  {
    path: '/',
    name: 'index',
    component: indexView,
  },
  {
    path: '/updatePassword',
    name: 'updatePassword',
    component: updatePassword,
  },
  {
    path: '/device',
    name: 'device',
    component: deviceInfo,
    details: "设备列表",
  },
  {
    path: '/group',
    name: 'group',
    component: group,
    details: "设备组",
  },
  {
    path: '/linkale',
    name: 'linkale',
    component: linkale,
    details: "设备联动",
  },
  {
    path: '/types',
    name: 'types',
    component: types,
    details: "设备型号",
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
    show:true,
  }
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes
})

export default router
