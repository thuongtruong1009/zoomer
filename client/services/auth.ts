import { ISignUpRequest, ISignInRequest } from '@/models'
import { axiosHttpInstance } from './axios'

export const AuthServices = {
    signup(payload: ISignUpRequest) {
        return axiosHttpInstance.post('/auth/signup', payload)
    },

    signin(payload: ISignInRequest) {
        return axiosHttpInstance.post('/auth/signin', payload)
    },

    logout() {
        return axiosHttpInstance.post('/logout')
    },
}
