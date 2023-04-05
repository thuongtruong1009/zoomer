import * as React from 'react'
import Toolbar from '@mui/material/Toolbar'
import Button from '@mui/material/Button'
import { Logo } from './Logo'
import { Stack, Tooltip } from '@mui/material'
import PersonAddIcon from '@mui/icons-material/PersonAdd'
import { SearchPopup } from '@/components'

export function Head() {
    return (
        <Toolbar
            sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                borderBottomRightRadius: '2rem',
                background:
                    'linear-gradient(45deg, #97DEFF 5%,  #E5D1FA 30%, #DFFFD8 60%, #FFC8C8 90%)',
            }}
        >
            <Logo />

            <Stack direction="row" justifyContent="center" alignItems="center" spacing={-2}>
                <Tooltip title="Search user">
                    <SearchPopup />
                </Tooltip>

                <Tooltip title="Add new">
                    <Button color="inherit">
                        <PersonAddIcon />
                    </Button>
                </Tooltip>
            </Stack>
        </Toolbar>
    )
}
