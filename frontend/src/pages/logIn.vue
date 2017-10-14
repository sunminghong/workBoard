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

    <!-- 注册页面屏蔽自动填充账号密码 -->
    <template v-if="$route.name === 'SignUp'">
      <Input placeholder="请输入账号" size="small" v-show="false"></Input>
      <Input placeholder="请输入密码" size="small" type="password" v-show="false"></Input>
    </template>

    <div class="login-main">
      <!-- 错误提示信息 -->
      <Alert type="error" show-icon v-if="errorMessage" class="login-message">
        {{errorMessage}}
      </Alert>

      <div class="login-username">
        <span class="title">账号</span>
        <div>
          <Input v-model="username" placeholder="请输入账号" size="small" autocomplete="off" autofocus></Input>
        </div>
      </div>

      <div class="login-username">
        <span class="title">密码</span>
        <div>
          <Input
            placeholder="请输入密码"
            size="small"
            type="password"
            autocomplete="off"
            v-model="password"
            @click="submit"
            @keydown.enter.meta.native="submit"
            @keydown.enter.ctrl.native="submit"></Input>
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
  errorMessage: string = ''
  created() {
    if (this.$route.name === 'LogIn' || this.$route.name === 'SignIn') {
      this.buttonText = '登录'
    }
  }

  async submit() {
    try {
      if (this.$route.name === 'LogIn') {
        await UsersApi.login(this.username, this.password)
      } else {
        await UsersApi.create(this.username, this.password)
      }
      this.$router.push({ path: '/' })
    } catch (err) {
      this.errorMessage = err.response.data.message
    }
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

.login-message {
  position: absolute;
  width: 100%;
  left: 0;
  top: -40px;
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

