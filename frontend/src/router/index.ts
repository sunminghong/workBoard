import Vue from 'vue'
import Router from 'vue-router'
import MainBoard from '@/pages/MainBoard.vue'
import logIn from '@/pages/logIn.vue'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/sign_up',
      name: 'SignUp',
      component: logIn,
    },
    {
      path: '/log_in',
      name: 'LogIn',
      component: logIn,
    },
    {
      path: '/',
      name: 'MainBoard',
      component: MainBoard,
    },
  ],
})

export default router
