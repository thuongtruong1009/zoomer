export class SocketConnection {
  socket: WebSocket

  constructor() {
      this.socket = new WebSocket('ws://localhost:8081/ws')
  }

  connect = (cb: any) => {
      this.socket.onopen = () => {
          console.log('Connected to websocket')
      }

      this.socket.onmessage = (msg: any) => {
          cb(msg)
      }

      this.socket.onclose = (event: any) => {
          console.log('Socket closed connection: ', event)
      }

      this.socket.onerror = (error: any) => {
          console.log('Socket error: ', error)
      }
  }

  sendMsg = (msg: any) => {
      console.log(msg)
      this.socket.send(JSON.stringify(msg))
  }

  connected = (user: any) => {
      this.socket.onopen = () => {
          console.log('Successfully Connected', user)
          // initiate mapping
          this.mapConnection(user)
      }
  }

  mapConnection = (user: any) => {
      console.log('mapping', user)
      this.socket.send(JSON.stringify({ type: 'bootup', user: user }))
  }
}
