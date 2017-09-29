import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello.vue'
import MainBoard from '@/components/MainBoard.vue'
import SignUp from '@/components/SignUp.vue'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/sign_up',
      name: 'SignUp',
      component: SignUp,
    },
    {
      path: '/',
      name: 'MainBoard',
      component: MainBoard,
    },
  ],
})

export default router
