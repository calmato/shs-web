import { initializeApp } from 'firebase/app'

const config = {
  apiKey: process.env.firebaseApiKey,
  authDomain: `${process.env.firebaseProjectId}.firebaseapp.com`,
  databaseURL: `https://${process.env.firebaseProjectId}.firebaseio.com`,
  projectId: process.env.firebaseProjectId,
  storageBucket: `${process.env.firebaseProjectId}.appspot.com`,
  messagingSenderId: process.env.firebaseMessagingSenderId,
}

const app = initializeApp(config)
export { app }
