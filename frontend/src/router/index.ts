import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello.vue'
import MainBoard from '@/components/MainBoard.vue'

Vue.use(Router)
const router = new Router({
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: MainBoard,
    },
  ],
})

export default router
