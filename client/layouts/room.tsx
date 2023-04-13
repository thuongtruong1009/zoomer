import { LayoutProps } from '@/models/common'
import { Box, Stack } from '@mui/material'
import React from 'react'
import Grid from '@mui/material/Grid'
import { AccountMenu, ContactList, Head, Panel } from '@/components'
import RoomSpecify from '@/pages/room/[roomId]'

export function RoomLayout({ children }: LayoutProps) {
    return (
        <Stack maxHeight="100vh" height="100vh" overflow="hidden" sx={{ bgcolor: 'primary.main' }}>
            <Grid container>
                <Grid item xs={3} sx={{ borderRight: '1px solid #e9e9e9' }}>
                    <Box sx={{ width: '100%', maxWidth: '100%' }}>
                        <Head />
                        <ContactList />
                        <AccountMenu data="ok" />
                    </Box>
                </Grid>
                <Grid item xs={9}>
                    <>{children}</>
                </Grid>
            </Grid>
        </Stack>
    )
}
