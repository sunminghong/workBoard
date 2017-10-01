// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import iView from 'iview'
import lodash from 'lodash'
import axios from 'axios'
import 'iview/dist/styles/iview.css'
import App from './App.vue'
import router from './router/index'
import './axios'

Vue.use(iView)

// ts 动态添加属性写法
const vuePrototype: any = Vue.prototype
vuePrototype.$http = axios
vuePrototype._ = lodash

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App },
})
