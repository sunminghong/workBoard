<template>
  <div id="app">
    <header-bar></header-bar>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
import HeaderBar from './components/Header.vue'

@Component({
  components: { HeaderBar },
})
export default class App extends Vue {
}
</script>

<style lang="scss">
@import './static/normalize';
@import './static/iview';
@import './static/index';

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
