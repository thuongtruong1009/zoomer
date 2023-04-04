import * as React from 'react'
import {
    AppBar,
    Button,
    ListItemText,
    Tooltip,
    Divider,
    ListItemIcon,
    Menu,
    MenuItem,
    Avatar,
} from '@mui/material'
import PersonAdd from '@mui/icons-material/PersonAdd'
import Settings from '@mui/icons-material/Settings'
import Logout from '@mui/icons-material/Logout'

export function AccountMenu() {
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null)
    const open = Boolean(anchorEl)
    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget)
    }
    const handleClose = () => {
        setAnchorEl(null)
    }
    return (
        <AppBar
            position="fixed"
            sx={{
                top: 'auto',
                left: 0,
                bottom: 0,
                maxWidth: '25%',
                py: 1,
                px: 2,
                borderTopRightRadius: '3rem',
                background: '#ADA2FF',
            }}
        >
            <Tooltip title="Account settings">
                <Button
                    onClick={handleClick}
                    size="small"
                    aria-controls={open ? 'account-menu' : undefined}
                    aria-haspopup="true"
                    aria-expanded={open ? 'true' : undefined}
                    sx={{
                        borderRadius: '0.5rem',
                        display: 'flex',
                        alignItems: 'center',
                        textAlign: 'left',
                        color: 'red',
                    }}
                >
                    <Avatar sx={{ width: 40, height: 40 }}>M</Avatar>
                    <ListItemText
                        primary="User 02"
                        sx={{
                            ml: 1.5,
                            span: { fontWeight: '500' },
                            '.MuiListItemText-secondary': { fontSize: '12px' },
                        }}
                    />
                </Button>
            </Tooltip>

            <Menu
                anchorEl={anchorEl}
                id="account-menu"
                open={open}
                onClose={handleClose}
                onClick={handleClose}
                PaperProps={{
                    elevation: 0,
                    sx: {
                        ml: 7,
                        borderRadius: '0.5rem',
                        overflow: 'visible',
                        filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
                        '& .MuiAvatar-root': {
                            width: 32,
                            height: 32,
                            ml: -0.5,
                            mr: 1,
                        },
                        '&:before': {
                            content: '""',
                            display: 'block',
                            position: 'absolute',
                            bottom: 20,
                            left: 0,
                            width: 10,
                            height: 10,
                            bgcolor: 'background.paper',
                            transform: 'translateX(-50%) rotate(45deg)',
                            zIndex: 0,
                        },
                    },
                }}
                transformOrigin={{ horizontal: 'right', vertical: 'bottom' }}
                anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
            >
                <MenuItem onClick={handleClose}>
                    <Avatar /> Profile
                </MenuItem>
                <MenuItem onClick={handleClose}>
                    <Avatar /> My account
                </MenuItem>
                <Divider />
                <MenuItem onClick={handleClose}>
                    <ListItemIcon>
                        <PersonAdd fontSize="small" />
                    </ListItemIcon>
                    Add another account
                </MenuItem>
                <MenuItem onClick={handleClose}>
                    <ListItemIcon>
                        <Settings fontSize="small" />
                    </ListItemIcon>
                    Settings
                </MenuItem>
                <MenuItem onClick={handleClose}>
                    <ListItemIcon>
                        <Logout fontSize="small" />
                    </ListItemIcon>
                    Logout
                </MenuItem>
            </Menu>
        </AppBar>
    )
}
