import firebase from 'firebase/compat/app';
import "firebase/compat/auth"

const firebaseConfig = {
  apiKey: "AIzaSyClk_Z6fi9sWDDDW6mA6NDzOQaAQAAdpSs",
  authDomain: "zoomer-45245.firebaseapp.com",
  projectId: "zoomer-45245",
  storageBucket: "zoomer-45245.appspot.com",
  messagingSenderId: "76109019772",
  appId: "1:76109019772:web:bf54b59878b271354827f9",
  measurementId: "G-61H8SXYTJW"
};

firebase.initializeApp(firebaseConfig)

export default firebase
