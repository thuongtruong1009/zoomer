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
import { RoomServices } from '@/services'

function refreshMessages(): MessageExample[] {
    const getRandomInt = (max: number) => Math.floor(Math.random() * Math.floor(max))

    return Array.from(new Array(50)).map(
        () => messageExamples[getRandomInt(messageExamples.length)]
    )
}

export function ChatBody() {
    const isSelf = (authorId: string): boolean => authorId === '0'

    //ws
    // const fetchChatHistory = async(u1: string, u2: string) => {
    //     const res = await RoomServices.getChatHistory(u1, u2)
    //     console.log(res)
    //     setChatHistory(res)
    // }

    const [to, setTo] = React.useState('1')
    const [from, setFrom] = React.useState('0')
    const [username, setUsername] = React.useState('0')
    const [msg, setMsg] = React.useState('')
    const [chats, setChats] = React.useState<any>([])
    const [chatHistory, setChatHistory] = React.useState<any>([])

    const [socketConn, setSocketConn] = React.useState<SocketConnection | null>(null)

    const handleWs = () => {
        const conn = new SocketConnection()
        setSocketConn(conn)

        conn.connect((message) => {
            console.log(message)
            const msg = JSON.parse(message.data)
            if (to === msg.from || username === msg.from) {
                setChats([...chats, msg])
                // setChatHistory([...chatHistory, msg])
            }
        })
    }

    // ------------------------------

    const [messages, setMessages] = React.useState(() => refreshMessages())

    const sendMessage = (message: string) => {
        setMessages((prevMessages) => [
            ...prevMessages,
            {
                secondary: message,
                person: '/static/images/avatar/1.jpg',
                authorId: '0',
            },
        ])
    }

    React.useEffect(() => {
        setMessages(refreshMessages())
    }, [setMessages])

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
                {messages.map(({ secondary, person, authorId }, index) => (
                    <ListItem key={index + person}>
                        <ListItemAvatar
                            sx={{
                                mr: -2,
                                display: isSelf(authorId) ? 'none' : 'block',
                            }}
                        >
                            <Avatar
                                alt="Profile Picture"
                                src={person}
                                sx={{ width: 28, height: 28 }}
                            />
                        </ListItemAvatar>

                        <ListItemText
                            sx={{
                                display: 'flex',
                                justifyContent: isSelf(authorId) ? 'end' : 'start',
                                alignItems: 'center',
                                width: '100%',
                            }}
                        >
                            <div
                                style={{
                                    maxWidth: containerWidth && containerWidth * 0.75,
                                    color: isSelf(authorId) ? '#fff' : '#0009',
                                    boxShadow: isSelf(authorId)
                                        ? ''
                                        : '5px 5px 10px #EAF5FC, -2px -2px 5px rgba(9,148,255,0.05), inset -2px -2px 5px #EAF5FC',
                                    border: '0.5px solid #e9e9e9',
                                    fontSize: '0.875rem',
                                    padding: isSelf(authorId)
                                        ? '0.45rem 1.25rem'
                                        : '0.55rem 1.25rem',
                                    // mt: '0.5rem',
                                    background: isSelf(authorId)
                                        ? 'linear-gradient(-90deg, hsla(216, 96%, 56%, 1) 0%, hsla(178, 64%, 65%, 1) 100%)'
                                        : 'white',
                                    borderRadius: '1rem',
                                }}
                            >
                                <span>{secondary}</span>
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

interface MessageExample {
    secondary: string
    person: string
    authorId: string
}

const messageExamples: readonly MessageExample[] = [
    {
        secondary: "I'll be in the neighbourhood this week. Let's grab a bite to eat",
        person: '/static/images/avatar/5.jpg',
        authorId: '1',
    },
    {
        secondary: `Do you have a suggestion for a good present for John on his work
      anniversary. I am really confused & would love your thoughts on it.`,
        person: '/static/images/avatar/1.jpg',
        authorId: '0',
    },
    {
        secondary: 'I am try out this new BBQ recipe, I think this might be amazing',
        person: '/static/images/avatar/2.jpg',
        authorId: '0',
    },
    {
        secondary: 'I have the tickets to the ReactConf for this year.',
        person: '/static/images/avatar/3.jpg',
        authorId: '1',
    },
    {
        secondary: 'My appointment for the doctor was rescheduled for next Saturday.',
        person: '/static/images/avatar/4.jpg',
        authorId: '0',
    },
    {
        secondary: `Menus that are generated by the bottom app bar (such as a bottom
      navigation drawer or overflow menu) open as bottom sheets at a higher elevation
      than the bar.`,
        person: '/static/images/avatar/5.jpg',
        authorId: '1',
    },
    {
        secondary: `Who wants to have a cookout this weekend? I just got some furniture
      for my backyard and would love to fire up the grill.`,
        person: '/static/images/avatar/1.jpg',
        authorId: '0',
    },
]
