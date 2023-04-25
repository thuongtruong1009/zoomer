import { localStore } from '@/utils'
import axios from 'axios'
import querystring from 'querystring'
import firebase from 'firebase'

const getFirebaseToken = async () => {
    const currentUser = firebase.auth().currentUser
    if (currentUser) return currentUser.getIdToken()

    //Not logged in
    const hasRememberedAccount = localStorage.getItem('firebaseui::rememberedAccounts')
    if (!hasRememberedAccount) return null

    // Logged in but current user is not fetched -> wait 10s
    return new Promise((resolve, reject) => {
        const waitTimer = setTimeout(() => {
            reject(null)
            console.log('Reject timeout')
        }, 10000)

        const unregisterAuthObserver = firebase.auth().onAuthStateChanged(async (user: any) => {
            if (!user) {
                reject(null)
            }
            const token = await user.getIdToken()
            console.log('[AXIOS] Logged in user token: ', token)
            resolve(token)

            unregisterAuthObserver()
            clearTimeout(waitTimer)
        })
    })
}

export const axiosHttpInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json',
        // Credentials: 'include',
    },
    paramsSerializer: (params: any) => querystring.stringify(params),
})

axiosHttpInstance.interceptors.request.use(
    async (config: any) => {
        // local token
        // const token = localStore.get('user').token

        // firebase token
        const token = await getFirebaseToken()
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    async (error: any) => {
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
