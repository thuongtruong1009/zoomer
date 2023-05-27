import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { RoomLayout } from '@/layouts'
import Image from 'next/image'
import { Typography } from '@mui/material'

const RoomDefault: NextPageWithLayout = () => {
    const width = 300
    const height = width * 0.7
    return (
        <Box sx={{display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', minHeight: '100vh'}}>
            <Image src="/empty_chat.webp" alt="empty_room" width={width} height={height} />
            <Typography sx={{color: '#757575', mt: 2}}>
              Choose a room to start chating
            </Typography>
        </Box>
    )
}

RoomDefault.Layout = RoomLayout

export default RoomDefault
