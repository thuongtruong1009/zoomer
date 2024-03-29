import { localStore } from '@/utils'
import axios from 'axios'
import querystring from 'querystring'

export const axiosHttpInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json',
        // Credentials: 'include',
    },
    paramsSerializer: (params: querystring.ParsedUrlQueryInput | undefined) =>
        querystring.stringify(params),
})

axiosHttpInstance.interceptors.request.use(
    async (config: any) => {
        // config.paramsSerializer = (params: querystring.ParsedUrlQueryInput | undefined) =>
        //     querystring.stringify(params)
        // local token
        const token = localStore.get('user')?.data?.token

        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    async (error) => {
        return Promise.reject(error)
    }
)

axiosHttpInstance.interceptors.response.use(
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
