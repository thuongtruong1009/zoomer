import { LayoutProps } from '@/models/common'
import { Box, Stack } from '@mui/material'
import React from 'react'
import Grid from '@mui/material/Grid'
import { AccountMenu, ContactList, RoomHeader } from '@/components'
import { localStore } from '@/utils'

export function RoomLayout({ children }: LayoutProps) {
  const username = localStore.get('user') ? localStore.get('user').data.username : 'unknown user'
    return (
        <Stack maxHeight="100vh" height="100vh" overflow="hidden" sx={{ bgcolor: 'primary.main' }}>
            <Grid container>
                <Grid item xs={3} sx={{ borderRight: '1px solid #e9e9e9', background: 'white' }}>
                    <Box sx={{ width: '100%', maxWidth: '100%'}}>
                        <RoomHeader />
                        <ContactList />
                        <AccountMenu username={username} />
                    </Box>
                </Grid>
                <Grid item xs={9}>
                    <>{children}</>
                </Grid>
            </Grid>
        </Stack>
    )
}
