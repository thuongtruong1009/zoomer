import * as React from 'react'
import Box from '@mui/material/Box'
import { SearchField, ContactList, AccountMenu } from '@/components'
import { Head } from '@/components/room/RoomHeader'
import { Paper } from '@mui/material'

export function Panel() {
    return (
        <Box sx={{ width: '100%', maxWidth: '100%' }}>
            {/* <SearchField /> */}
            <Head />

            <ContactList />

            {/* <Paper sx={{ position: 'fixed', bottom: 0, left: 0, width: '25%' }} elevation={3}> */}
            <AccountMenu />
            {/* </Paper> */}
        </Box>
    )
}
