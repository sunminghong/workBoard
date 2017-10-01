import axios from 'axios'

export default class {
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
