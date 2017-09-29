<template>
  <main class="main-board">
    <div>
      账号
      <Input v-model="userName" placeholder="请输入..." style="width: 300px"></Input>
    </div>
    <div>
      密码
      <Input v-model="userName" placeholder="请输入..." style="width: 300px"></Input>
    </div>
  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component

export default class SignUp extends Vue {
  msg: string = 'this is a typescript project now'
  userName: string = ''
  passwork: string = ''
}

</script>


<style lang="scss">
</style>

