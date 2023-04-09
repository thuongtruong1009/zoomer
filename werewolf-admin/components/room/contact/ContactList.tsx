import * as React from 'react'
import List from '@mui/material/List'
import ListItemButton from '@mui/material/ListItemButton'
import ListItemText from '@mui/material/ListItemText'
import ListItemAvatar from '@mui/material/ListItemAvatar'
import Avatar from '@mui/material/Avatar'
import PersonIcon from '@mui/icons-material/Person'
import MoreHorizIcon from '@mui/icons-material/MoreHoriz'
import IconButton from '@mui/material/IconButton'
import { useRouter } from 'next/router'
import { ListSubheader, Paper, Typography } from '@mui/material'
import { BasicPopover } from '@/components'
import ContactOption from './ContactOption'
import { useEffect } from 'react'
import { RoomServices } from '@/services'
import { localStore } from '@/utils'
import moment from 'moment'

export const ContactList = () => {
    const router = useRouter()
    const [selectedIndex, setSelectedIndex] = React.useState(0)
    const [contacts, setContacts] = React.useState<any>([])

    useEffect(() => {
        const fetchData = async () => {
            const res = await RoomServices.getContactsList(localStore.get('user').username)
            setContacts(res.data)
        }

        fetchData().catch(console.error)
    }, [])

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
                maxHeight: 'calc(100vh - 8rem)',
                px: 1,
            }}
        >
            {contacts.map((contact: string, idx: number) => (
                <ListItemButton
                    key={idx}
                    selected={selectedIndex === idx}
                    onClick={(event) => handleListItemClick(event, idx)}
                    sx={{
                        borderRadius: '0.8rem',
                        position: 'relative',
                        bgcolor: 'white',
                        boxShadow:
                            '5px 5px 10px #EAF5FC, -2px -2px 5px rgba(9,148,255,0.05), inset -2px -2px 5px #EAF5FC',
                        border: '0.5px solid #e9e9e9',
                        my: 0.9,
                        '.btn': {
                            visibility: 'hidden',
                        },
                        '&.Mui-selected': {
                            backgroundColor: '#F5CA9D',
                            boxShadow:
                                '-10px -10px 20px rgba(255, 255, 255, 0.5), 10px 10px 20px rgba(0, 0, 0, 0.1)',
                            border: 'none',
                            '&:hover': {
                                backgroundColor: '#F5CA9D',
                            },
                        },
                        '&:hover, &.Mui-mousedown': {
                            backgroundColor: 'transparent',
                            '&:hover .btn': {
                                visibility: 'visible',
                            },
                        },
                    }}
                >
                    <ListItemAvatar>
                        <Avatar sx={{ boxShadow: '0 0 0 1px #fff' }}>
                            <PersonIcon />
                        </Avatar>
                    </ListItemAvatar>

                    <ListItemText
                        primary={
                            <React.Fragment>
                                <span
                                    style={{
                                        position: 'absolute',
                                        top: '2px',
                                        right: '9px',
                                        fontWeight: 400,
                                        color: '#0009',
                                        fontSize: '0.8em',
                                    }}
                                >
                                    {'· '}
                                    {moment().startOf('hour').fromNow()}
                                    <br />
                                </span>
                                <Typography
                                    sx={{
                                        fontWeight: '500',
                                    }}
                                    component="span"
                                    variant="body2"
                                    color="text.primary"
                                >
                                    {contact.username}
                                </Typography>
                            </React.Fragment>
                        }
                        secondary={
                            <React.Fragment>
                                {" — I'll be in your neighborhood doing errands this…"}
                            </React.Fragment>
                        }
                    />

                    <IconButton edge="end" aria-label="more" className="btn">
                        {/* <MoreHorizIcon /> */}
                        <BasicPopover title={<MoreHorizIcon />} content={<ContactOption />} />
                    </IconButton>
                </ListItemButton>
            ))}
        </List>
    )
}
