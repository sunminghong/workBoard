<template>
  <header class="header">
    <div class="header-menu"></div>
    <div class="header-user" v-if="user.nickname">
      <span class="nickname">{{user.nickname}}</span>
      <Avatar icon="person" size="large" />
    </div>

    <div class="header-login" v-else>
      <router-link :to="{ name: 'LogIn'}" v-if="$route.name !== 'LogIn'">登录</router-link>
      <router-link :to="{ name: 'SignUp'}" v-if="$route.name !== 'SignUp'">注册</router-link>
    </div>
  </header>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import Component from 'vue-class-component'
import UsersRx from '../api/users'

const usersRx = new UsersRx()

@Component
export default class HeaderBar extends Vue {
  name: 'HeaderBar'
  user: any = usersRx.data
  created() {
    axios.get('/users/me', {
    }).then((res: any) => {
      usersRx.init()
      usersRx.set(res.data)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">
.header {
  z-index: 2;
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 50px;
  padding: 0 40px;
}
.header-user {
  .nickname {
    margin-right: 10px;
  }
}
</style>
