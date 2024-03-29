import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/e',
      name: 'e',
      component: () => import(/* webpackChunkName: "about" */ './pages/echartgl.vue')
    },
    {
      path: '/',
      name: 'ws',
      component: () => import(/* webpackChunkName: "about" */ './pages/WS.vue')
    },
    {
      path: '/l3d',
      name: 'l3d',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './pages/L3D.vue')
    }
  ]
})
