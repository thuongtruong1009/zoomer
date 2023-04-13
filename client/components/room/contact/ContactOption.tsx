import * as React from 'react'
import List from '@mui/material/List'
import ListItem from '@mui/material/ListItem'
import ListItemButton from '@mui/material/ListItemButton'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemText from '@mui/material/ListItemText'
import StarIcon from '@mui/icons-material/Star'
import LockIcon from '@mui/icons-material/Lock'

export default function InsetList() {
    return (
        <List
            sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}
            aria-label="contacts"
        >
            <ListItem disablePadding>
                <ListItemButton>
                    <ListItemIcon>
                        <StarIcon />
                    </ListItemIcon>
                    <ListItemText primary="Add to favorite" />
                </ListItemButton>
            </ListItem>
            <ListItem disablePadding>
                <ListItemButton>
                    <ListItemIcon>
                        <LockIcon />
                    </ListItemIcon>
                    <ListItemText primary="Change to private" />
                </ListItemButton>
            </ListItem>
            <ListItem disablePadding>
                <ListItemButton>
                    <ListItemText inset primary="child inset" />
                </ListItemButton>
            </ListItem>
        </List>
    )
}
