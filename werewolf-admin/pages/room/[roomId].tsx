import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { RoomLayout } from '@/layouts'

const Home: NextPageWithLayout = () => {
    return (
        <Box>
            <h1>contact 1</h1>
        </Box>
    )
}

Home.Layout = RoomLayout

export default Home
