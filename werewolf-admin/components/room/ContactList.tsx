import * as React from 'react'
import Box from '@mui/material/Box'
import List from '@mui/material/List'
import ListItemButton from '@mui/material/ListItemButton'
import ListItemText from '@mui/material/ListItemText'
import ListItemAvatar from '@mui/material/ListItemAvatar'
import Avatar from '@mui/material/Avatar'
import PersonIcon from '@mui/icons-material/Person'
import MoreHorizIcon from '@mui/icons-material/MoreHoriz'
import IconButton from '@mui/material/IconButton'
import { useRouter } from 'next/router'

export function ContactList() {
    const router = useRouter()
    const [selectedIndex, setSelectedIndex] = React.useState(1)

    const handleListItemClick = (
        event: React.MouseEvent<HTMLDivElement, MouseEvent>,
        index: number
    ) => {
        setSelectedIndex(index)
        router.push(`/room/${index}`)
    }

    return (
        <List
            component="nav"
            aria-label="main mailbox folders"
            sx={{
                position: 'relative',
                overflowY: 'auto',
                maxHeight: '80vh',
                px: 1,
            }}
        >
            {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((value, idx) => (
                <ListItemButton
                    key={idx}
                    selected={selectedIndex === idx}
                    onClick={(event) => handleListItemClick(event, idx)}
                    sx={{
                        borderRadius: '10px',
                        '.btn': {
                            visibility: 'hidden',
                        },
                        '&:hover, &.Mui-selected, &.Mui-mousedown': {
                            backgroundColor: '#EAF5FC',
                            '&:hover .btn': {
                                visibility: 'visible',
                            },
                        },
                    }}
                >
                    <ListItemAvatar>
                        <Avatar>
                            <PersonIcon />
                        </Avatar>
                    </ListItemAvatar>
                    <ListItemText
                        primary="User 02"
                        secondary="July 20, 2014"
                        sx={{
                            span: { fontWeight: '500' },
                            '.MuiListItemText-secondary': { fontSize: '12px' },
                        }}
                    />

                    <IconButton edge="end" aria-label="more" color="secondary" className="btn">
                        <MoreHorizIcon />
                    </IconButton>
                </ListItemButton>
            ))}
        </List>
    )
}
