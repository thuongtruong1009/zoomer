import firebase from 'firebase/compat'
import { useRouter } from 'next/router'
import { useEffect } from 'react'

export default function AuthProviderCallbackPage() {
  const router  = useRouter()
  const { provider } = router.query

  useEffect(()=> {
    if(provider) {
      firebase.auth().getRedirectResult().then((result: { user: any }) => {
        if(result.user) {
          router.push("/room");
        }else {
          router.push("/auth/signin")
        }
      }).catch((error: any)=> {
        console.error("Error during authentication: ", error)
        router.push("/login")
      })
    }
  }, [provider, router])

  return <div>Authenticating...</div>
}
