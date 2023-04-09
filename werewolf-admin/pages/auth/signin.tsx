import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { AuthLayout } from '@/layouts'
import { FormControl, TextField, Button } from '@mui/material'
import { useState } from 'react'
import { AuthServices } from '@/services'
import { useRouter } from 'next/router'
import { localStore } from '@/utils'

const Home: NextPageWithLayout = () => {
    const router = useRouter()

    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [usernameValid, setUsernameValid] = useState(false)
    const [passwordValid, setPasswordValid] = useState(false)

    const onChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(e.target.value)
        if (username !== '') {
            setUsernameValid(true)
        }
    }

    const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value)
        if (password !== '') {
            setPasswordValid(true)
        }
    }

    const onSubmit = async () => {
        if (usernameValid && passwordValid) {
            try {
                const res = await AuthServices.signin({ username, password })
                if (res) {
                    localStore.set('user', res)
                    router.push('/room')
                } else {
                    console.log(res)
                }
            } catch (error) {
                console.log({ message: 'something went wrong' + error, isInvalid: true })
            }
        }
    }

    return (
        <Box>
            <FormControl
                component="form"
                sx={{
                    '& > :not(style)': { m: 1, width: '25ch' },
                }}
                noValidate
                autoComplete="off"
            >
                <h1>Signup</h1>
                <TextField
                    id="outlined-basic"
                    label="Username"
                    variant="outlined"
                    color="secondary"
                    helperText="Your username must be at least 3 characters"
                    value={username}
                    onChange={onChangeUsername}
                />
                <TextField
                    id="outlined-password-input"
                    label="Password"
                    type="password"
                    color="secondary"
                    autoComplete="current-password"
                    helperText="Your password must be at least 8 characters"
                    value={password}
                    onChange={onChangePassword}
                />

                <Button variant="contained" onClick={onSubmit}>
                    Login
                </Button>
            </FormControl>
        </Box>
    )
}

Home.Layout = AuthLayout

export default Home
