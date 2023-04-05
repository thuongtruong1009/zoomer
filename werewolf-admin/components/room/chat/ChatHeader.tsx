import * as React from 'react'
import { Avatar, IconButton, Paper, Stack, Typography } from '@mui/material'
import VideocamSharpIcon from '@mui/icons-material/VideocamSharp'
import QueryStatsIcon from '@mui/icons-material/QueryStats'
import WallpaperIcon from '@mui/icons-material/Wallpaper'
import WidgetsIcon from '@mui/icons-material/Widgets'

export const ChatHeader = () => {
    return (
        <Paper
            sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                background: 'white',
                padding: '0.5rem 1rem',
                boxShadow: '0 0 5px #e9e9',
            }}
        >
            <Stack direction="row" spacing={2} alignItems={'center'}>
                <Avatar sx={{ width: 34, height: 34 }}>M</Avatar>
                <Typography variant="subtitle1" sx={{ fontWeight: 500 }}>
                    Mai
                </Typography>
            </Stack>

            <Stack direction="row" spacing={2}>
                <IconButton aria-label="stats" component="label">
                    <QueryStatsIcon />
                </IconButton>

                <IconButton aria-label="capture screenshot" component="label">
                    <WallpaperIcon />
                </IconButton>

                <IconButton aria-label="call" component="label">
                    <VideocamSharpIcon />
                </IconButton>

                <IconButton aria-label="call" component="label">
                    <WidgetsIcon />
                </IconButton>
            </Stack>
        </Paper>
    )
}
