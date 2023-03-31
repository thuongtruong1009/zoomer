import * as React from 'react'
import Box from '@mui/material/Box'
import Divider from '@mui/material/Divider'
import ListSubheader from '@mui/material/ListSubheader'
import { SearchField, ContactList, AccountMenu } from '@/components'
import { Head } from '@/components/room/head'

export function Panel() {
    return (
        <Box sx={{ width: '100%', maxWidth: '100%', bgcolor: 'background.paper' }}>
            {/* <SearchField /> */}
            <Head />

            <ContactList />

            <Divider />

            <AccountMenu />
        </Box>
    )
}
