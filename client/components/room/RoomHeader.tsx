import * as React from 'react'
import Toolbar from '@mui/material/Toolbar'
import Button from '@mui/material/Button'
import { Logo } from './Logo'
import { Paper, Stack, Tooltip } from '@mui/material'
import PersonAddIcon from '@mui/icons-material/PersonAdd'
import { SearchPopup } from '@/components'

export const RoomHeader: React.FC = () => {
    return (
        <Paper
            sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                background: 'white',
                    // 'linear-gradient(45deg, #97DEFF 5%,  #E5D1FA 30%, #DFFFD8 60%, #FFC8C8 90%)',
                height: '3.5rem',
            }}
            // elevation={0}
        >
            <Logo />

            <Stack direction="row" justifyContent="center" alignItems="center" spacing={-2}>
                <Tooltip title="Search user">
                    <SearchPopup />
                </Tooltip>
            </Stack>
        </Paper>
    )
}
