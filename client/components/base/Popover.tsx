import * as React from 'react'
import Popover from '@mui/material/Popover'
import Typography from '@mui/material/Typography'
import Button from '@mui/material/Button'
import { OverridableComponent } from '@mui/material/OverridableComponent'
import { SvgIconTypeMap } from '@mui/material'

interface IPopoverProps {
    title: React.ReactNode
    content: React.ReactNode
}

export const BasicPopover = ({ title, content }: IPopoverProps) => {
    const [anchorEl, setAnchorEl] = React.useState<HTMLButtonElement | null>(null)

    const handleClick = (event: React.MouseEvent<HTMLDivElement | any>): void => {
        setAnchorEl(event.currentTarget)
    }

    const handleClose = () => {
        setAnchorEl(null)
    }

    const open = Boolean(anchorEl)
    const id = open ? 'simple-popover' : undefined

    return (
        <>
            <div aria-describedby={id} onClick={handleClick}>
                {title}
            </div>
            <Popover
                id={id}
                open={open}
                anchorEl={anchorEl}
                onClose={handleClose}
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'left',
                }}
            >
                {content}
            </Popover>
        </>
    )
}
