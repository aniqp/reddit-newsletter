import { initializeApp } from 'firebase/app';
import { getFirestore } from 'firebase/firestore';

const firebaseConfig = {
    apiKey: "AIzaSyCyibgQOTZot8DC_5XBJf8vgT2dkP21DoY",
    authDomain: "reddit-newsletter-5d86e.firebaseapp.com",
    projectId: "reddit-newsletter-5d86e",
    storageBucket: "reddit-newsletter-5d86e.appspot.com",
    messagingSenderId: "954987295318",
    appId: "1:954987295318:web:2f41338188b798acb341f7",
    measurementId: "G-2NZ1ZNYD2G"
};

const app = initializeApp(firebaseConfig);
const db = getFirestore(app);

export { db }
