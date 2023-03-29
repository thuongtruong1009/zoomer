import axios from 'axios'

export const axiosClient = axios.create({
    baseURL: '/api',
    headers: {
        'Content-Type': 'application/json',
    },
})

axiosClient.interceptors.response.use(
    function (response: any) {
        return response.data
    },
    function (error: any) {
        return Promise.reject(error)
    }
)
