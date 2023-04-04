import * as React from 'react'
import Toolbar from '@mui/material/Toolbar'
import Button from '@mui/material/Button'
import { Logo } from './Logo'
import { Tooltip } from '@mui/material'
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
                background: '#ADA2FF',
            }}
        >
            <Logo />

            <div>
                <Tooltip title="Search user">
                    <SearchPopup />
                </Tooltip>

                <Tooltip title="Add new">
                    <Button color="inherit">
                        <PersonAddIcon />
                    </Button>
                </Tooltip>
            </div>
        </Toolbar>
    )
}
