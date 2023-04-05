import * as React from 'react'
import Paper from '@mui/material/Paper'
import InputBase from '@mui/material/InputBase'
import Divider from '@mui/material/Divider'
import IconButton from '@mui/material/IconButton'
import MenuIcon from '@mui/icons-material/Menu'
import SearchIcon from '@mui/icons-material/Search'
import DirectionsIcon from '@mui/icons-material/Directions'
import EmojiEmotionsIcon from '@mui/icons-material/EmojiEmotions'
import ThumbUpIcon from '@mui/icons-material/ThumbUp'

export function ChatInput() {
    return (
        <Paper
            component="form"
            sx={{
                p: '2px 4px',
                display: 'flex',
                alignItems: 'center',
                width: '100%',
                borderRadius: '20px',
            }}
        >
            <IconButton sx={{ p: '10px' }} aria-label="menu">
                <MenuIcon />
            </IconButton>

            <InputBase
                sx={{ ml: 1, flex: 1 }}
                placeholder="Search Google Maps"
                inputProps={{ 'aria-label': 'search google maps' }}
            />

            <IconButton type="button" sx={{ p: '10px' }} aria-label="search">
                <EmojiEmotionsIcon />
            </IconButton>

            <IconButton color="primary" sx={{ p: '10px' }} aria-label="directions">
                <ThumbUpIcon />
            </IconButton>
        </Paper>
    )
}
