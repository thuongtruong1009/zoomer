import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { RoomLayout } from '@/layouts'

const RoomDefault: NextPageWithLayout = () => {
    return (
        <Box>
            <h1>Choose random a contect to start chating</h1>
        </Box>
    )
}

RoomDefault.Layout = RoomLayout

export default RoomDefault
