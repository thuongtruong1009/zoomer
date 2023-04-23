import firebase from 'firebase'
import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { unwrapResult } from '@reduxjs/toolkit'
import { userApi } from '@/services'

const config = {
    apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
    authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
}
firebase.initializeApp(config)

function OAuth() {
    const dispatch = useDispatch()

    useEffect(() => {
        const unregisterAuthObserver = firebase.auth().onAuthStateChanged(async (user: any) => {
            if (!user) {
                console.log('user is not logged in')
                return
            }
            try {
                const currentUser = await userApi.getMe()
                // const actionResult = await dispatchEvent(getMe())
                // const currentUser = unwrapResult(actionResult)
                console.log('Logged in user: ', currentUser)
            } catch (error: any) {
                console.log('Failed to login: ', error.message)
            }
        })
        return () => unregisterAuthObserver()
    }, [])
    return
}
