import * as React from 'react'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import Image from 'next/image'
import { Logo } from './Logo'

export function Head() {
    return (
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <Logo />

            <Button color="inherit">Login</Button>
        </Toolbar>
    )
}
