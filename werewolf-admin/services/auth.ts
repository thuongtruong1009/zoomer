import { LoginPayload } from '@/models'
import { axiosClient } from './axios'

export const authApi = {
    login(payload: LoginPayload) {
        return axiosClient.post('/login', payload)
    },

    logout() {
        return axiosClient.post('/logout')
    },

    getProfile() {
        return axiosClient.get('/profile')
    },
}
