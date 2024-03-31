import { db } from './firebaseConfig';
import { query, collection, setDoc, doc, getDoc, getDocs, updateDoc } from 'firebase/firestore'

export async function checkUserExists(userId, userData) {
    // Documentation: https://modularfirebase.web.app/common-use-cases/firestore/#get-a-document
    const userRef = doc(db, 'users', userId);
    const userSnapshot = await getDoc(userRef);

    if (userSnapshot.exists()) {
        await updateDoc(userRef, {
            accessToken: userData.accessToken,
            refreshToken: userData.refreshToken
          });
    } else {
        addNewUser(userId, userData);
    }
    return getUserSubreddits(userId);
}

export async function addNewUser(userId, userData) {
    try {
        const userRef = doc(db, "users", userId);
        await setDoc(userRef, userData);

        const response = await fetch('http://localhost:3000/api/subreddits', { 
            headers: {
                'Access-Token': userData.accessToken
            },
        });
        const subreddits = await response.json();

        console.log(subreddits)
        subreddits.forEach(async (subreddit) => {
            const subredditRef = doc(db, `users/${userId}/subreddits`, subreddit.display_name);
            await setDoc(subredditRef, { description: subreddit.description });
        });
        console.log("User written with ID: ", userRef.id);
    } catch (e) {
        console.error("Error adding user: ", e);
    }
}

export async function getUserSubreddits(userId) {
    try {
        console.log(userId)
        const subredditsRef = collection(db, `users/${userId}/subreddits`);
        const q = query(subredditsRef);
        const querySnapshot = await getDocs(q);
  
        const subreddits = querySnapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
        }));
  
        // console.log(subreddits);
        return subreddits;
    } catch (e) {
        console.error("Error fetching user subreddits: ", e);
    }
}
