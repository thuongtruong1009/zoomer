import { NextPageWithLayout } from '@/models'
import { Box } from '@mui/system'
import { AuthLayout } from '@/layouts'
import { FormControl, TextField, Button } from '@mui/material'
import { useState } from 'react'
import { AuthServices } from '@/services'
import { useRouter } from 'next/router'

const Home: NextPageWithLayout = () => {
    const router = useRouter()

    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [usernameValid, setUsernameValid] = useState(false)
    const [passwordValid, setPasswordValid] = useState(false)

    const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        if (name !== '') {
            setUsername(value)
            setUsernameValid(true)
        } else if (password !== '') {
            setPassword(value)
            setPasswordValid(true)
        }
    }

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (usernameValid && passwordValid) {
            try {
                const res = await AuthServices.signup({ username, password })
                if (res.data) {
                    router.push('/auth/signin')
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
                    onChange={onChange}
                />
                <TextField
                    id="outlined-password-input"
                    label="Password"
                    type="password"
                    color="secondary"
                    autoComplete="current-password"
                    helperText="Your password must be at least 8 characters"
                    value={password}
                    onChange={onChange}
                />

                <Button variant="contained" onClick={onSubmit}>
                    Register
                </Button>
            </FormControl>
        </Box>
    )
}

Home.Layout = AuthLayout

export default Home
