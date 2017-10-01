<template>
  <main class="login" id="particles">
    <!-- 隐藏浏览器自动填充 -->
    <vue-particles class="login-bg"
      color="#dedede"
      :particleOpacity="1"
      :particlesNumber="40"
      shapeType="circle"
      :particleSize="5"
      linesColor="#dedede"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.6"
      :linesDistance="350"
      :moveSpeed="3"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"></vue-particles>

    <Input placeholder="请输入账号" size="small" v-show="false"></Input>
    <Input placeholder="请输入密码" size="small" type="password" v-show="false"></Input>

    <div class="login-main">
      <div class="login-username">
        <span class="title">账号</span>
        <div>
          <Input v-model="username" placeholder="请输入账号" size="small" autocomplete="off" autofocus></Input>
        </div>
      </div>

      <div class="login-username">
        <span class="title">密码</span>
        <div>
          <Input v-model="password" placeholder="请输入密码" size="small" type="password" autocomplete="off"></Input>
        </div>
      </div>

      <Button type="primary" class="login-submit" @click="submit">{{buttonText}}</Button>
    </div>

  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
import UsersApi from '../api/users'

@Component

export default class login extends Vue {
  username: string = ''
  password: string = ''
  vue: any = this
  buttonText: string = '注册'
  created() {
    console.log(112)
    if (this.$route.name === 'LogIn') {
      this.buttonText = '登录'
    }
  }
  submit() {
    if (this.$route.name === 'LogIn') {
      UsersApi.login(this.username, this.password)
      return
    }

    UsersApi.create(this.username, this.password)
  }
}

</script>


<style lang="scss">
.login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.login-main {
  z-index: 1;
  position: relative;
}

.login-username,
.login-password {
  display: flex;
  align-items: center;

  &:not(:last-of-type) {
    margin-bottom: 20px;
  }

  .title {
    display: flex;
    align-items: center;
    margin-right: 30px;
  }
}

.login-submit {
  margin-top: 20px;
}

#particles-js {
  position: fixed;
  z-index: 0;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>

