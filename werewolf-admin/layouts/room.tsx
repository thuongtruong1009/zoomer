import { LayoutProps } from '@/models/common'
import { Stack } from '@mui/material'
import { Box } from '@mui/system'
import React from 'react'
import { Footer, Header } from '@/components'

export function RoomLayout({ children }: LayoutProps) {
    return (
        <Stack minHeight="100vh">
            <Header />

            <Box component="main" flexGrow={1}>
                {children}
            </Box>

            <Footer />
        </Stack>
    )
}
