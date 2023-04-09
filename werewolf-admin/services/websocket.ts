class SocketConnection {
    socket: WebSocket

    constructor() {
        this.socket = new WebSocket(process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8081/api')
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
