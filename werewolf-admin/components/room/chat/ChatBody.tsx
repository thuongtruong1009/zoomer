import * as React from 'react'
import Box from '@mui/material/Box'
import CssBaseline from '@mui/material/CssBaseline'
import Paper from '@mui/material/Paper'
import List from '@mui/material/List'
import ListItem from '@mui/material/ListItem'
import ListItemAvatar from '@mui/material/ListItemAvatar'
import ListItemText from '@mui/material/ListItemText'
import Avatar from '@mui/material/Avatar'
import { useRef } from 'react'
import { ChatInput } from '@/components'

function refreshMessages(): MessageExample[] {
    const getRandomInt = (max: number) => Math.floor(Math.random() * Math.floor(max))

    return Array.from(new Array(50)).map(
        () => messageExamples[getRandomInt(messageExamples.length)]
    )
}

export function ChatBody() {
    const [value, setValue] = React.useState(0)
    const ref = React.useRef<HTMLDivElement>(null)
    const [messages, setMessages] = React.useState(() => refreshMessages())

    React.useEffect(() => {
        ;(ref.current as HTMLDivElement).ownerDocument.body.scrollTop = 0
        setMessages(refreshMessages())
    }, [value, setMessages])

    const bottomRef = useRef(null)

    return (
        <Box
            sx={{
                pb: 7,
                background:
                    'linear-gradient(45deg, #97DEFF 5%,  #E5D1FA 30%, #DFFFD8 60%, #FFC8C8 90%)',
            }}
            ref={ref}
        >
            <CssBaseline />
            <List sx={{ overflowY: 'scroll', maxHeight: 'calc(100vh - 8rem)' }}>
                {messages.map(({ secondary, person }, index) => (
                    <ListItem key={index + person}>
                        <ListItemAvatar sx={{ mr: -2 }}>
                            <Avatar
                                alt="Profile Picture"
                                src={person}
                                sx={{ width: 28, height: 28 }}
                            />
                        </ListItemAvatar>
                        <ListItemText
                            secondary={secondary}
                            sx={{
                                py: 1,
                                px: 1.5,
                                // mt: '0.5rem',
                                backgroundColor: 'white',
                                borderRadius: '2rem',
                                maxWidth: 'fit-content',
                            }}
                        />
                    </ListItem>
                ))}
            </List>

            <Paper
                sx={{ position: 'fixed', bottom: 0, right: 0, width: '75%', height: '4rem' }}
                elevation={3}
            >
                <ChatInput />
            </Paper>
        </Box>
    )
}

interface MessageExample {
    secondary: string
    person: string
}

const messageExamples: readonly MessageExample[] = [
    {
        secondary: "I'll be in the neighbourhood this week. Let's grab a bite to eat",
        person: '/static/images/avatar/5.jpg',
    },
    {
        secondary: `Do you have a suggestion for a good present for John on his work
      anniversary. I am really confused & would love your thoughts on it.`,
        person: '/static/images/avatar/1.jpg',
    },
    {
        secondary: 'I am try out this new BBQ recipe, I think this might be amazing',
        person: '/static/images/avatar/2.jpg',
    },
    {
        secondary: 'I have the tickets to the ReactConf for this year.',
        person: '/static/images/avatar/3.jpg',
    },
    {
        secondary: 'My appointment for the doctor was rescheduled for next Saturday.',
        person: '/static/images/avatar/4.jpg',
    },
    {
        secondary: `Menus that are generated by the bottom app bar (such as a bottom
      navigation drawer or overflow menu) open as bottom sheets at a higher elevation
      than the bar.`,
        person: '/static/images/avatar/5.jpg',
    },
    {
        secondary: `Who wants to have a cookout this weekend? I just got some furniture
      for my backyard and would love to fire up the grill.`,
        person: '/static/images/avatar/1.jpg',
    },
]
