import * as React from 'react'
import { Avatar, IconButton, Paper, Stack, Typography } from '@mui/material'
import VideocamSharpIcon from '@mui/icons-material/VideocamSharp'
import QueryStatsIcon from '@mui/icons-material/QueryStats'
import WallpaperIcon from '@mui/icons-material/Wallpaper'
import MenuIcon from '@mui/icons-material/Menu'

const WindowOptions = () => {
  const width = 1000;
  const height = 600;
  const left = (window.innerWidth / 2) - (width / 2);
  const top = (window.innerHeight / 2) - (height / 2);
  const size = 'width=' + width + ', height=' + height + ', left=' + left + ', top=' + top;
  const options = size + ', toolbar=yes, scrollbars=yes, resizable=yes';
  return options;
}

export const ChatHeader = () => {
    const createStream = async () => {
      const resp = await fetch("http://localhost:8081/create");
      const { data } = await resp.json();

      window.open(`/stream/${data.room_id}`, '_blank', WindowOptions())
    };

    return (
        <Paper
            sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                background:
                    'linear-gradient(45deg, #97DEFF 5%,  #E5D1FA 30%, #DFFFD8 60%, #FFC8C8 90%)',
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

                <IconButton aria-label="call" component="label" onClick={createStream}>
                    <VideocamSharpIcon />
                </IconButton>

                <IconButton aria-label="call" component="label">
                    <MenuIcon />
                </IconButton>
            </Stack>
        </Paper>
    )
}
