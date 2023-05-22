import { io } from 'socket.io-client';

export class SocketConnection {
  socket: any;

  constructor() {
    this.socket = io('http://localhost:8080');
  }

  connect = (cb: any) => {
    this.socket.on('connect', () => {
      console.log('Connected to WebSocket');
    });

    this.socket.on('message', (msg: any) => {
      cb(msg);
    });

    this.socket.on('disconnect', (event: any) => {
      console.log('Socket closed connection:', event);
    });

    this.socket.on('error', (error: any) => {
      console.log('Socket error:', error);
    });
  };

  sendMsg = (msg: any) => {
    console.log(msg);
    this.socket.send(JSON.stringify(msg));
  };

  connected = (user: any) => {
    this.socket.on('connect', () => {
      console.log('Successfully connected:', user);
      this.mapConnection(user);
    });
  };

  mapConnection = (user: any) => {
    console.log('Mapping:', user);
    this.socket.send(JSON.stringify({ type: 'bootup', user: user }));
  };
}
