import { ISignUpRequest, ISignInRequest } from '@/models'
import axiosInstance from './axios'

export const AuthServices = {
    signup(payload: ISignUpRequest) {
        return axiosInstance.post('/auth/signup', payload)
    },

    signin(payload: ISignInRequest) {
        return axiosInstance.post('/auth/signin', payload)
    },

    logout() {
        return axiosInstance.post('/logout')
    },
}
