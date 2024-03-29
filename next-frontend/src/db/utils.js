import { db } from './firebaseConfig';
import { collection, addDoc, doc, getDoc, updateDoc } from 'firebase/firestore'

export async function checkUserExists(user) {
    const docRef = doc(db, 'users', user.id);
    const docSnapshot = await getDoc(docRef);
    
    if (docSnapshot.exists()) {
        await updateDoc(docRef, {
            accessToken: user.accessToken,
            refreshToken: user.refreshToken
          }).then(() => {
            console.log('User tokens updated successfully');
          }).catch((error) => {
            console.error('Error updating user tokens:', error);
          });
    } else {
        addNewUser(user);
    }
}

export async function addNewUser(user) {
    try {
        const docRef = await addDoc(collection(db, "users"), {
            id: user.id,
            name: user.name,
            oauth_client_id: user.oauth_client_id,
            accessToken: user.accessToken,
            refreshToken: user.refreshToken
        });

        console.log("User written with ID: ", docRef.id);
    } catch (e) {
        console.error("Error adding user: ", e);
    }
}
