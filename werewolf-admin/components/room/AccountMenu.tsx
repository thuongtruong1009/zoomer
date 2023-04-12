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
    Stack,
    IconButton,
} from '@mui/material'
import PersonAdd from '@mui/icons-material/PersonAdd'
import Settings from '@mui/icons-material/Settings'
import Logout from '@mui/icons-material/Logout'
import { ToggleColorMode } from '../ToggleColorMode'

export function AccountMenu(props: any) {
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
                maxHeight: '4rem',
                py: 1,
                px: 2,
                background: 'white',
            }}
        >
            <Stack direction="row" justifyContent="space-between" alignItems="center">
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
                            '&:hover': { background: '#C0DEFF' },
                        }}
                    >
                        <Avatar sx={{ width: 34, height: 34 }}>M</Avatar> {props.data}
                        <ListItemText
                            primary="User 02"
                            sx={{
                                mx: 1.5,
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

                    <Divider />

                    <MenuItem onClick={handleClose}>
                        <ListItemIcon>
                            <PersonAdd fontSize="small" />
                        </ListItemIcon>
                        Add another account
                    </MenuItem>

                    {/* <MenuItem onClick={handleClose}>
                        <ListItemIcon>
                            <Tooltip title="Dark/Light mode">
                                <ToggleColorMode />
                            </Tooltip>
                        </ListItemIcon>
                        {`Dark/Light mode`}
                    </MenuItem> */}

                    <MenuItem onClick={handleClose}>
                        <ListItemIcon>
                            <Logout fontSize="small" />
                        </ListItemIcon>
                        Logout
                    </MenuItem>
                </Menu>

                <Tooltip title="Setting">
                    <IconButton>
                        <Settings fontSize="small" />
                    </IconButton>
                </Tooltip>
            </Stack>
        </AppBar>
    )
}
