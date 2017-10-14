<template>
  <main class="main-board">
    <header class="main-board-nav">
    </header>

    <draggable v-model="boardList$" element="ul" class="main-board-list" :options="{ draggable: 'li.draggable', filter: '' }">
      <li v-for="board in boardList$" :key="board.name" class="draggable">
        <h6 class="main-board-title">{{board.title}}</h6>
        <h6>xxx</h6>
        <a href="javascript:" class="board-task-create">添加新的任务</a>
      </li>

      <li>
        <div v-if="isShowCreateInput" class="main-board-create">
          <Input placeholder="请输入..."
            v-model="newBoardTitle"
            :autofocus="true"
             @click="createNewBoard"
             @keydown.enter.meta.native="createNewBoard"
             @keydown.enter.ctrl.native="createNewBoard"
            ></Input>

          <div class="main-board-create-buttons">
            <Button type="text" @click="isShowCreateInput = false">取消</Button>
            <Button type="primary" @click="createNewBoard">保存</Button>
          </div>
        </div>

        <a v-else href="javascript:" class="board-task-create" @click="isShowCreateInput = true">新建任务列表</a>
      </li>
    </draggable>
  </main>
</template>

<script lang="ts">
import Vue from 'vue'
import draggable from 'vuedraggable'
import Component from 'vue-class-component'
import { BoardListRx, Board } from '../RxService/boardList'

@Component({
  components: {
    draggable,
  },
})

export default class MainBoard extends Vue {
  boardList$: Board[] = BoardListRx.data
  newBoardTitle: string = ''
  isShowCreateInput: boolean = false
  created() {
    BoardListRx.init()

    BoardListRx.set([
      {
        createdAt: null,
        title: '第一个列表',
        token: null,
      },
    ])
  }
  // 创建新的 Board
  createNewBoard() {
    BoardListRx.set([
      {
        createdAt: null,
        title: this.newBoardTitle,
        token: null,
      },
    ])

    this.isShowCreateInput = false
    this.newBoardTitle = ''
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
  padding: 0 5px;
  text-align: left;
  white-space: nowrap;
  overflow-x: auto;
  overflow-y: hidden;
  transform: translate3d(0, 0, 0);
  // 去间距
  font-size: 0;
  -webkit-text-size-adjust: none;

  li {
    display: inline-flex;
    flex-direction: column;
    vertical-align: top;
    margin-left: 5px;
    margin-right: 5px;
    width: 290px;
    height: 100%;
    background-color: red;
    border-radius: 3px;
    background-color: #eee;
    overflow-y: auto;
    overflow-x: hidden;
  }
}

.main-board-title {
  display: flex;
  align-items: center;
  flex: 0 0 auto;
  margin: 0 10px;
  height: 50px;
  font-size: 17px;
  color: #1c2438;
  font-weight: 500;
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

.main-board-create {
  margin: 10px 10px 0;

  input {
    height: 40px;
    font-size: 16px;
    font-weight: 500;
    color: #1c2438;
  }
}

.main-board-create-buttons {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-top: 10px;

  button {
    margin-left: 10px;
  }
}
</style>

