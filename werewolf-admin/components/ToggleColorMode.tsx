import * as React from 'react'
import IconButton from '@mui/material/IconButton'
import Box from '@mui/material/Box'
import { useTheme, ThemeProvider, createTheme } from '@mui/material/styles'
import Brightness4Icon from '@mui/icons-material/Brightness4'
import Brightness7Icon from '@mui/icons-material/Brightness7'

const ColorModeContext = React.createContext({ toggleColorMode: () => {} })

function MyApp() {
    const theme = useTheme()
    const colorMode = React.useContext(ColorModeContext)
    return (
        <Box
            sx={{
                // bgcolor: 'background.default',
                bgcolor: '#C0DEFF',
                color: 'text.primary',
                borderRadius: 1,
            }}
        >
            <IconButton onClick={colorMode.toggleColorMode} color="inherit">
                {theme.palette.mode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
            </IconButton>
        </Box>
    )
}

export function ToggleColorMode() {
    const [mode, setMode] = React.useState<'light' | 'dark'>('light')
    const colorMode = React.useMemo(
        () => ({
            toggleColorMode: () => {
                setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'))
            },
        }),
        []
    )

    const theme = React.useMemo(
        () =>
            createTheme({
                palette: {
                    mode,
                },
            }),
        [mode]
    )

    return (
        <ColorModeContext.Provider value={colorMode}>
            <ThemeProvider theme={theme}>
                <MyApp />
            </ThemeProvider>
        </ColorModeContext.Provider>
    )
}
