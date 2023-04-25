import axios from 'axios'
import querystring from 'querystring'

export const axiosWsInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_WS_URL || 'http://localhost:8081/stream',
    headers: {
        'Content-Type': 'application/json',
    },
    paramsSerializer: (params: any) => querystring.stringify(params),
})

axiosWsInstance.interceptors.request.use(
    async (config: any) => {
        return config
    },
    async (error: any) => {
        return Promise.reject(error)
    }
)

axiosWsInstance.interceptors.response.use(
    (response: any) => {
        if (response && response.data) {
            return response.data
        }
        return response
    },
    (error: any) => {
        return Promise.reject(error)
    }
)
