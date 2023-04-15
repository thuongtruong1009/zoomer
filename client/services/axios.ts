import { localStore } from '@/utils'
import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json',
        Credentials: 'include',
    },
})

axiosInstance.interceptors.request.use(
    function (config: any) {
        const token = localStore.get('user').token
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    function (error: any) {
        return Promise.reject(error)
    }
)

axiosInstance.interceptors.response.use(
    function (response: any) {
        return response.data
    },
    function (error: any) {
        return Promise.reject(error)
    }
)

export default axiosInstance
