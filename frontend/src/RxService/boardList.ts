import Rx from 'rxjs/Rx'

interface Board {
  createdAt: string | null,
  title: string | null,
  token: string | null,
}

class BoardListRx {
  static data: Board[] = []

  static data$$: any = {}

  static init() {
    this.data$$ = new Rx.Subject()

    // 订阅
    this.data$$.map((boardList: Board[]) => boardList)
    // 推送到视图更新
    .subscribe({
      next: (boardList: Board[]) => this.data.push(...boardList),
    })
  }
  static set(boardList: Board[]) {
    // 推送新的消息
    this.data$$.next(boardList)
  }
}

export { BoardListRx, Board }
