import axios from 'axios'
import Rx from 'rxjs/Rx'

interface User {
  id: number;
  username: string;
  nickname: string;
  password?: string;
}

export default class {
  data: User = {
    id: 0,
    username: '',
    nickname: '',
    password: '',
  }
  data$$: any = {}
  init() {
    this.data$$ = new Rx.Subject()

    // 订阅
    this.data$$.map((user: User) => user)
    // 推送到视图更新
    .subscribe({
      next: (user: User) => Object.assign(this.data, user),
    })
  }
  unsubscribe() {
    this.data$$.unsubscribe()
  }
  set(user: User) {
    // 推送新的消息
    this.data$$.next(user)
  }
  static create(username: string, password: string) {
    return new Promise((resolve, reject) => {
      axios.post('/users/create', {
        username,
        password,
      }).then((res: any) => {
        resolve(res)
      }).catch((err: any) => {
        reject(err)
      })
    })
  }
  static login(username: string, password: string) {
    return new Promise((resolve, reject) => {
      axios.post('/users/login', {
        username,
        password,
      }).then((res: any) => {
        resolve(res)
      }).catch((err: any) => {
        reject(err)
      })
    })
  }
}
