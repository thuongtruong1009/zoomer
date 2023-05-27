import Link from "@mui/material/Link";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { NextPageWithLayout } from '@/models'
import { AuthLayout } from '@/layouts'
import { TextField, Button } from '@mui/material'
import { useState } from 'react'
import { AuthServices } from '@/services'
import { useRouter } from 'next/router'

const SignUp: NextPageWithLayout = () => {
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

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (usernameValid && passwordValid) {
            try {
                const res = await AuthServices.signup({ username, password })
                if (res) {
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
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          boxShadow: 3,
          borderRadius: 2,
          px: 4,
          py: 6,
          color: "#4dabf5",
          background: "#fff",
        }}

      >
        <Typography component="h1" variant="h5">
          Sign up
        </Typography>
        <Box component="form" onSubmit={onSubmit} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            autoFocus
            helperText={username === "" ? "Your username must be filled" : ""}
            value={username}
            onChange={onChangeUsername}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            helperText={password === "" ? "Your password must be filled" : ""}
            value={password}
            onChange={onChangePassword}
          />

          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2, background: '#4dabf5', ":hover": { background: '#2196f3' } }}
          >
            Sign Up
          </Button>
        </Box>
      </Box>
    </Container>
  );
}

SignUp.Layout = AuthLayout

export default SignUp
