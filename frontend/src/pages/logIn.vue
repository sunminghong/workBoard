<template>
  <main class="signup">
    <!-- 隐藏浏览器自动填充 -->
    <Input placeholder="请输入账号" size="small" v-show="false"></Input>
    <Input placeholder="请输入密码" size="small" type="password" v-show="false"></Input>

    <div class="signup-username">
      <span class="title">账号</span>
      <div>
        <Input v-model="username" placeholder="请输入账号" size="small" autocomplete="off" autofocus></Input>
      </div>
    </div>

    <div class="signup-username">
      <span class="title">密码</span>
      <div>
        <Input v-model="password" placeholder="请输入密码" size="small" type="password" autocomplete="off"></Input>
      </div>
    </div>

    <Button type="primary" class="signup-submit" @click="submit">登录</Button>
  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
import UsersApi from '../api/users'

@Component

export default class SignUp extends Vue {
  msg: string = 'this is a typescript project now'
  username: string = ''
  password: string = ''
  submit() {
    UsersApi.login(this.username, this.password)
  }
}

</script>


<style lang="scss">
.signup {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.signup-username,
.signup-password {
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

.signup-submit {
  margin-top: 20px;
}
</style>

