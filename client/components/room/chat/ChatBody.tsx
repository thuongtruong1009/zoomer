import * as React from 'react'
import Box from '@mui/material/Box'
import CssBaseline from '@mui/material/CssBaseline'
import Paper from '@mui/material/Paper'
import List from '@mui/material/List'
import ListItem from '@mui/material/ListItem'
import ListItemAvatar from '@mui/material/ListItemAvatar'
import ListItemText from '@mui/material/ListItemText'
import Avatar from '@mui/material/Avatar'
import { ChatInput } from '@/components'
import { RoomServices} from '@/services'
import { localStore } from '@/utils'
import { useRouter } from 'next/router'
import { SocketConnection } from '@/services/websocket'

export function ChatBody() {
    const router = useRouter()
    const isSelf = (authorId: string): boolean => authorId === localStore.get('user').data.username

    const [to, setTo] = React.useState<any>([])
    const [from, setFrom] = React.useState<any>([])
    const [msg, setMsg] = React.useState([])
    const [chats, setChats] = React.useState<any>([])
    const [chatHistory, setChatHistory] = React.useState<any>([])

    const conn = new SocketConnection()

    function handleWs() {
        try {
            conn.connect((message: any) => {
                const msg = JSON.parse(JSON.stringify(message))
                if (
                    router.query.roomId === msg.from ||
                    localStore.get('user').data.username === msg.from
                ) {
                    setChats([...chats, msg])
                    // setChatHistory([...chatHistory, msg])
                }
            })
            conn.connected(localStore.get('user').data.username)
        } catch (err) {
            console.log('Error: ', err)
        }
    }

    // const fetchChatHistory = async (u1: string, u2: string) => {
    //     // const res = await RoomServices.getChatHistory(u1, u2)
    //     const res = await RoomServices.getChatHistory({ params: { u1: u1, u2: u2 } })
    //     console.log(res)
    //     if (res.status && res['data'].length !== undefined) {
    //         setChats(res.data.data.reverse())
    //         console.log('1', res)
    //         setChatHistory(res.data.data.reverse())
    //     } else {
    //         setChatHistory([])
    //     }
    // }

    const fetchChatHistory = async (u1: string, u2: string) => {
      const res = await RoomServices.getChatHistory({ params: { u1: u1, u2: u2 } })
      console.log(res)

      if (res.status && res.data && Array.isArray(res.data.data)) {
        setChats(res.data.data.slice().reverse())
        console.log('1', res)
        setChatHistory(res.data.data.slice().reverse())
      } else {
        setChatHistory([])
      }
    }


    const sendMessage = (message: string) => {
        const msg = {
            type: 'message',
            user: localStore.get('user').data.username,
            chat: {
                from: localStore.get('user').data.username,
                to: String(router.query.roomId),
                msg: message,
                msg_type: 'text',
            },
        }
        conn.sendMsg(msg)
        fetchChatHistory(localStore.get('user').data.username, String(router.query.roomId))
    }

    // const sendMessageTo = (to: any) => {
    //     setTo(to)
    //     fetchChatHistory(localStore.get('user').data.username, to)
    // }

    React.useEffect(() => {
        handleWs()
        if (router.query.roomId) {
            console.log(router.query.roomId)
            // sendMessageTo(router.query.roomId)
            fetchChatHistory(localStore.get('user').data.username, String(router.query.roomId))
        }
    }, [router.query.roomId])

    const containerRef = React.useRef<HTMLUListElement>(null)
    React.useEffect(() => {
        const container = containerRef.current
        if (container) {
            container.scrollTop = container.scrollHeight
        }
    }, [containerRef.current?.innerHTML, sendMessage])
    const containerWidth = containerRef.current?.clientWidth

    return (
        <Box
            sx={{
                pb: 7,
            }}
        >
            <CssBaseline />

            <List sx={{ overflowY: 'scroll', maxHeight: 'calc(100vh - 8rem)' }} ref={containerRef}>
                {chats.map((m: any, index: number) => (
                    <ListItem key={m.id}>
                        <ListItemAvatar
                            sx={{
                                mr: -2,
                                display: isSelf(m.from) ? 'none' : 'block',
                            }}
                        >
                            <Avatar
                                alt="Profile Picture"
                                src={'/static/images/avatar/1.jpg'}
                                sx={{ width: 28, height: 28 }}
                            />
                        </ListItemAvatar>

                        <ListItemText
                            sx={{
                                display: 'flex',
                                justifyContent: isSelf(m.from) ? 'end' : 'start',
                                alignItems: 'center',
                                width: '100%',
                            }}
                        >
                            <div
                                style={{
                                    maxWidth: containerWidth && containerWidth * 0.75,
                                    color: isSelf(m.from) ? '#fff' : '#0009',
                                    boxShadow: isSelf(m.from)
                                        ? ''
                                        : '5px 5px 10px #EAF5FC, -2px -2px 5px rgba(9,148,255,0.05), inset -2px -2px 5px #EAF5FC',
                                    border: '0.5px solid #e9e9e9',
                                    fontSize: '0.875rem',
                                    padding: isSelf(m.from) ? '0.45rem 1.25rem' : '0.55rem 1.25rem',
                                    // mt: '0.5rem',
                                    background: isSelf(m.from)
                                        ? 'linear-gradient(-90deg, hsla(216, 96%, 56%, 1) 0%, hsla(178, 64%, 65%, 1) 100%)'
                                        : 'white',
                                    borderRadius: '1rem',
                                }}
                            >
                                <span>
                                    {m.msg}
                                    {/* {m.timestamp} */}
                                </span>
                            </div>
                        </ListItemText>
                    </ListItem>
                ))}
            </List>

            <Paper
                sx={{
                    position: 'fixed',
                    bottom: 0,
                    right: 0,
                    width: '75%',
                    minHeight: '4rem',
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'end',
                    p: 1,
                    background:
                        'linear-gradient(45deg, #97DEFF 5%,  #E5D1FA 30%, #DFFFD8 60%, #FFC8C8 90%)',
                }}
                elevation={3}
            >
                <ChatInput onData={sendMessage} />
            </Paper>
        </Box>
    )
}
