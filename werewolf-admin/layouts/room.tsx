import { LayoutProps } from '@/models/common'
import { Stack } from '@mui/material'
import React from 'react'
import Grid from '@mui/material/Grid'
import { Panel } from '@/components'

export function RoomLayout({ children }: LayoutProps) {
    return (
        <Stack maxHeight="100vh" overflow="hidden">
            <Grid container>
                <Grid item xs={3} sx={{ borderRight: '1px solid #e9e9e9' }}>
                    <Panel />
                </Grid>
                <Grid item xs={9}>
                    <>{children}</>
                </Grid>
            </Grid>
        </Stack>
    )
}
