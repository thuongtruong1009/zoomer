import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
})

axiosInstance.interceptors.request.use(
    function (config: any) {
        const token = localStorage.getItem('token')
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
