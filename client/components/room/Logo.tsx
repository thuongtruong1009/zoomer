import * as React from 'react'
import Typography from '@mui/material/Typography'
import Image from 'next/image'
import { Stack } from '@mui/material'

export function Logo() {
    return (
        <Stack direction="row" alignItems="center" sx={{px: 2, display: 'flex', gap: 2}}>
            <Image alt="Logo" src="/logo-64.png" title="logo_image" width={36} height={36} />
            <Image alt="Logo" src="/Brand.png" title="brand_image" width={96} height={22} />

            {/* <Typography
                variant="h6"
                component="div"
                sx={{
                    ml: 2,
                    background:
                        'linear-gradient(90deg, rgba(255, 0, 0, 1) 0%, rgba(255, 154, 0, 1) 10%, rgba(208, 222, 33, 1) 20%, rgba(79, 220, 74, 1) 30%, rgba(63, 218, 216, 1) 40%, rgba(47, 201, 226, 1) 50%, rgba(28, 127, 238, 1) 60%, rgba(95, 21, 242, 1) 70%, rgba(186, 12, 248, 1) 80%, rgba(251, 7, 217, 1) 90%, rgba(255, 0, 0, 1) 100%)',
                    backgroundClip: 'text',
                    WebkitBackgroundClip: 'text',
                }}
            >
                Zoomer
            </Typography> */}
        </Stack>
    )
}
