import Rx from 'rxjs/Rx'

interface Task {
  createdAt: string | null,
  title: string | null,
  token: string | null,
}

class TaskRx {
  static data: Task[] = []

  static data$$: any = {}

  static init() {
    this.data$$ = new Rx.Subject()

    // 订阅
    this.data$$.map((task: Task[]) => task)
      // 推送到视图更新
      .subscribe({
        next: (task: Task[]) => Object.assign(this.data, task),
      })
  }

  static set(task: Task[]) {
    // 推送新的消息
    this.data$$.next(task)
  }
}

export { TaskRx, Task }
