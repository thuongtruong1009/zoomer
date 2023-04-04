import * as React from 'react'
import Box from '@mui/material/Box'
import { SearchField, ContactList, AccountMenu } from '@/components'
import { Head } from '@/components/room/head'

export function Panel() {
    return (
        <Box sx={{ width: '100%', maxWidth: '100%' }}>
            {/* <SearchField /> */}
            <Head />

            <ContactList />

            <AccountMenu />
        </Box>
    )
}
