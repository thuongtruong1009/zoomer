class SocketInstance {
    socket: WebSocket

    constructor() {
        this.socket = new WebSocket('ws://localhost:3000/api/chats')
    }

    connect = (cb: (arg0: any) => void) => {
        this.socket.onopen = () => {
            console.log('Connected to websocket')
        }

        this.socket.onmessage = (event: any) => {
            cb(event.data)
        }

        this.socket.onclose = (event: any) => {
            console.log('Socket closed connection: ', event)
        }

        this.socket.onerror = (error: any) => {
            console.log('Socket error: ', error)
        }
    }

    sendMsg = (msg: string | ArrayBufferLike | Blob | ArrayBufferView) => {
        this.socket.send(JSON.stringify(msg))
    }
}
