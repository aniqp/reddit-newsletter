import { db } from './firebaseConfig';
import { query, collection, setDoc, doc, getDoc, getDocs, updateDoc } from 'firebase/firestore'

/**
 * Checks if a user exists in the Firestore database. If the user exists,
 * their access and refresh tokens are updated.
 *
 * @param {Object} user - User object containing user details.
 * @param {string} user.id - The ID of the user (reddit username).
 * @param {Object} user.tokens - Object containing accessToken and refreshToken.
 * @param {string} user.tokens.accessToken - User's access token.
 * @param {string} user.tokens.refreshToken - User's refresh token.
 * @returns {Promise<boolean>} - Returns true if the user exists and their tokens get updated. Returns false if the user does not exist.
 */
export async function checkUserExists(user) {
    const userRef = doc(db, 'users', user.id);
    const userSnapshot = await getDoc(userRef);

    if (userSnapshot.exists()) {
        await updateDoc(userRef, {
            accessToken: user.tokens.accessToken,
            refreshToken: user.tokens.refreshToken
        });
        return true;
    }
    return false;
}

export async function addNewUser(user) {
    try {
        const userRef = doc(db, "users", user.id);
        await setDoc(userRef, {
            accessToken: user.tokens.accessToken, 
            refreshToken: user.tokens.refreshToken,
            email: '',
        });
    } catch (e) {
        console.error("Error adding user: ", e);
    }
}

export async function addUserSubreddits(userId, subreddits) {
    const userRef = doc(db, "users", userId);

    subreddits.forEach(async (subreddit) => {
        const subredditRef = doc(db, `users/${userId}/subreddits`, subreddit.display_name);
        await setDoc(subredditRef, { description: subreddit.description, starred: false });
    });
}

export async function updateUserEmail(userId, email) {
    const userRef = doc(db, 'users', userId);
    const userSnapshot = await getDoc(userRef);

    if (userSnapshot.exists()) {
        await updateDoc(userRef, {
            email: email,
          });
    }
}

export async function fetchUserData(userId) {
    const userRef = doc(db, 'users', userId);
    const docSnapshot = await getDoc(userRef);

    if (docSnapshot.exists()) {
        const userData = docSnapshot.data();
        console.log("Document data:", userData)
        return { email: userData.email }; // Assuming the document has an 'email' field
    } else {
        // Handle the case where the document does not exist
        console.log("No such document!");
        return null;
    }
}

export async function updateStarred(userId, subredditId, value) {
    const subredditsRef = doc(db, `users/${userId}/subreddits`, subredditId);
    const userSnapshot = await getDoc(subredditsRef);

    if (userSnapshot.exists()) {
        await updateDoc(subredditsRef, {
            starred: value,
          });
    }
}

export async function getUserSubreddits(userId) {
    try {
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
