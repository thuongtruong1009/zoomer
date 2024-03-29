import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { RoomLayout } from '@/layouts'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { ChatHeader, ChatBody } from '@/components'

const RoomSpecify: NextPageWithLayout = () => {
    const router = useRouter()

    const [roomId, setRoomId] = useState<string>('')
    useEffect(() => {
        if (router.query.roomId) {
            setRoomId(router.query.roomId as string)
        }
    }, [router.query.roomId])

    return (
        <Box>
            <ChatHeader />
            <ChatBody />
        </Box>
    )
}

RoomSpecify.Layout = RoomLayout

export default RoomSpecify
