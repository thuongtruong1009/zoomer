import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { MainLayout } from '@/layouts'

const Home: NextPageWithLayout = () => {
    return (
        <Box>
            <h1>Home</h1>
        </Box>
    )
}

Home.Layout = MainLayout

export default Home
