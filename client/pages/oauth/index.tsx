import { useEffect } from "react";
import firebase from "firebase/compat/app";

export default function LoginPage() {
  useEffect(() => {
    const uiConfig = {
      signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.GithubAuthProvider.PROVIDER_ID,
        firebase.auth.FacebookAuthProvider.PROVIDER_ID,
      ],
      signInFlow: "popup",
      callbacks: {
        signInSuccessWithAuthResult: () => false,
      },
    };

    const startFirebaseUI = () => {
      if (typeof window !== "undefined") {
        const firebaseui = require("firebaseui");
        const ui = new firebaseui.auth.AuthUI(firebase.auth());
        ui.start("#firebaseui-auth-container", uiConfig);
      }
    };

    startFirebaseUI();
  }, []);

  return (
    <div>
      <h1>Login</h1>
      <div id="firebaseui-auth-container"></div>
    </div>
  );
}
