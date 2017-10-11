<template>
  <main class="main-board">
    <header class="main-board-nav">

    </header>
    <ul class="main-board-list">
      <li v-for="board in boardList" :key="board.name" class="main-board-body">
        <h6 class="main-board-title">{{board.title}}</h6>
        <ul class="board-task-list">
          <li v-for="task in board.tasks" :key="task.title" class="board-task-brief">
            {{task.title}}
          </li>
        </ul>
        <a href="javascript:" class="board-task-create">添加新的任务</a>
      </li>
    </ul>
  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'
import axios from 'axios'
import { Board } from '../service/mainBoard'

@Component

export default class MainBoard extends Vue {
  msg: string = 'this is a typescript project now'
  boardList: Board[] = [
    {
      title: '近期任务',
      tasks: [
        {
          title: '任务1',
          completeAt: null,
        },
        {
          title: '任务2',
          completeAt: '2017-1-1',
        },
      ],
    },
  ]
  mounted() {
    axios.get('/users/me', {
    }).then((res: any) => {
      console.log(res)
    })
  }
}
</script>


<style lang="scss">
.main-board {
  z-index: 1;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  top: 50px;
  padding-top: 60px;
  padding-bottom: 10px;
}

.main-board-nav {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 50px;
  background-color: #eee;
}

.main-board-list {
  height: 100%;
  padding-left: 10px;
  padding-right: 10px;
  display: flex;
}

.main-board-body {
  display: flex;
  flex-direction: column;
  width: 290px;
  background-color: red;
  border-radius: 3px;
  background-color: #eee;

  &:not(:last-of-type) {
    margin-right: 10px;
  }
}

.main-board-title {
  display: flex;
  align-items: center;
  flex: 0 0 auto;
  margin: 0 10px;
  height: 50px;
  font-size: 17px;
  font-weight: 600;
  line-height: 23px;
  cursor: move;
}

.board-task-list {
  display: flex;
  flex: 1 1 auto;
  flex-direction: column;
  overflow-y: auto;
  overflow-x: hidden;
}

.board-task-brief {
  display: flex;
  align-items: center;
  margin-left: 10px;
  margin-right: 10px;
  padding: 12px;
  height: 50px;
  font-size: 15px;
  line-height: 21px;
  font-weight: 500;
  background-color: #fff;
  border-radius: 3px;
  cursor: pointer;

  &:not(:last-of-type) {
    margin-bottom: 10px;
  }
}

.board-task-create {
  display: flex;
  flex: 0 0 auto;
  align-items: center;
  padding: 12px;
  margin: 10px;
  height: 50px;
  font-size: 16px;
  font-weight: 500;
  line-height: 22px;
  transition: background-color 0.2s linear, color 0.2s linear;

  &:hover {
    background-color: #e2e2e2;
  }
}
</style>

