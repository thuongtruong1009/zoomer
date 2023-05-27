import { axiosHttpInstance } from '@/services/http'
import { DefaultLayout } from '@/layouts'
import { AppPropsWithLayout } from '@/models'
import { createEmotionCache, theme } from '@/utils'
import { CacheProvider } from '@emotion/react'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider } from '@mui/material/styles'
import { SWRConfig } from 'swr'
import '../styles/globals.css'
import { Provider } from 'react-redux'
import { store } from '@/store/configureStore'

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

function MyApp({ Component, pageProps }: AppPropsWithLayout) {
    const Layout = Component.Layout ?? DefaultLayout

    return (
        <CacheProvider value={clientSideEmotionCache}>
            <ThemeProvider theme={theme}>
                <CssBaseline />

                <SWRConfig
                    value={{
                        fetcher: (url: string) => axiosHttpInstance.get(url),
                        shouldRetryOnError: false,
                    }}
                >
                    <Provider store={store}>
                        <Layout>
                            <Component {...pageProps} />
                        </Layout>
                    </Provider>
                </SWRConfig>
            </ThemeProvider>
        </CacheProvider>
    )
}
export default MyApp
