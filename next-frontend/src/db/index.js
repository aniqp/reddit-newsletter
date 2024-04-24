import { db } from './firebaseConfig';
import { 
    query, 
    collection, 
    setDoc, 
    doc, 
    getDoc, 
    getDocs, 
    updateDoc, 
    deleteDoc 
} from 'firebase/firestore'

/**
 * Adds a new user to the Firestore database.
 * 
 * It sets the initial values for the user's access token, refresh token, and initializes
 * the email field as an empty string.
 * 
 * @param {Object} user - An object containing the user's information.
 * @param {string} user.id - The ID of the user, used as the document key.
 * @param {Object} user.tokens - An object containing the user's access and refresh tokens.
 * @param {string} user.tokens.accessToken - The user's access token.
 * @param {string} user.tokens.refreshToken - The user's refresh token.
 */
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

/**
 * Adds a user to all subreddits in the subreddits collection.
 * 
 * @param {string} userId - The ID of the user to add to all subreddits
 * @param {Array} subreddits - An array of subreddits to add the user to
 */
export async function addUserToAllSubreddits(userId, subreddits){
    try {
        for (const subreddit of subreddits) {
            const subredditId = subreddit.id;
            const subredditUserRef = doc(db, `subreddits/${subredditId}/users`, userId);
            await setDoc(subredditUserRef, {});
        }
    } catch (e) {
        console.error("Error adding user to all subreddits: ", e);
        throw new Error("Failed to add user to all subreddits");
    }
}

/**
 * Adds new subreddits to a user's profile and updates subreddits list for existing users.
 * 
 * @param {string} userId - The ID of the user.
 * @param {Array} subreddits - An array of a user's subreddits.
 */
export async function addUserSubreddits(userId, subreddits) {
    const subredditsRef = collection(db, `users/${userId}/subreddits`);
    
    try {
        // Fetch existing subreddits from Firestore
        const existingSubsSnapshot = await getDocs(subredditsRef);
        const existingSubs = existingSubsSnapshot.docs.map(doc => ({
            id: doc.id,
        }));

        // Create a map of new subreddit display names for quick lookup
        const newSubsMap = new Map(subreddits.map(sub => [sub.display_name, sub]));

        // Determine subreddits to add (not already in DB) and to remove (not in the new list)
        const subredditsToAdd = subreddits.filter(sub => !existingSubs.some(existing => existing.id === sub.display_name));

        const subredditsToRemove = existingSubs.filter(existing => !newSubsMap.has(existing.id));

        // Add new subreddits
        for (const subreddit of subredditsToAdd) {
            const newSubRef = doc(db, `users/${userId}/subreddits`, subreddit.display_name);
            await setDoc(newSubRef, { display_name_prefixed: subreddit.display_name_prefixed, description: subreddit.description, starred: false });
        }

        // Remove subreddits no longer in the list
        for (const subreddit of subredditsToRemove) {
            const subToDeleteRef = doc(db, `users/${userId}/subreddits`, subreddit.id);
            await deleteDoc(subToDeleteRef);
        }
    } catch (e) {
        console.error("Error updating user subreddits: ", e);
        throw new Error("Failed to update user subreddits");
    }
}

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

export async function checkSubredditUserList(subredditId) {
    try {
        const subredditRef = doc(db, `subreddits`, subredditId);
        const subredditUsersQuery = query(collection(db, `subreddits/${subredditId}/users`));
        const subredditUsersSnapshot = await getDocs(subredditUsersQuery);
            
        if (subredditUsersSnapshot.empty) {
            await deleteDoc(subredditRef);
        }
    } catch (e) {
        console.error("Error checking subreddit user list: ", e);
        throw new Error("Failed to check subreddit user list");
    }
}

/**
 * Deletes a user and all associated subreddits from Firestore.
 *
 * @param {string} userId - The ID of the user to be deleted.
 */
export async function deleteUser(userId) {
    try {
        const userRef = doc(db, 'users', userId);
        const subredditsRef = collection(db, `users/${userId}/subreddits`);
        const querySnapshot = await getDocs(subredditsRef);

        // First remove the user from all subreddits in the subreddits collection
        removeUserFromAllSubreddits(userId, querySnapshot.docs);

        // Then delete all subreddits associated with the user
        const deleteSubredditPromises = querySnapshot.docs.map(doc => deleteDoc(doc.ref));
        await Promise.all(deleteSubredditPromises);

        // Last delete the user document
        await deleteDoc(userRef);

        console.log(`Successfully deleted user ${userId} and all associated subreddits.`);
    } catch (e) {
        console.error(`Error deleting user ${userId}:`, e);
        throw new Error(`Failed to delete user ${userId}`);
    }
}

/**
 * Retrieves the email address of a specific user from the Firestore database.
 * 
 * @param {string} userId - The ID of the user whose email address is to be fetched.
 * @returns {Promise<Object|null>} - A promise that resolves to an object containing the email
 * address if the user exists, or null if the user does not exist.
 */
export async function fetchUserEmail(userId) {
    const userRef = doc(db, 'users', userId);
    const docSnapshot = await getDoc(userRef);

    if (docSnapshot.exists()) {
        const userData = docSnapshot.data();
        return { email: userData.email }; 
    } else {
        console.log("No such document!");
        return null;
    }
}

/**
 * Retrieves all subreddits associated with a specific user.
 * 
 * @param {string} userId - The ID of the user whose subreddits are to be fetched.
 * @returns {Promise<Object[]>} A promise that resolves to an array of subreddit objects.
 */
export async function getUserSubreddits(userId) {
    try {
        const subredditsRef = collection(db, `users/${userId}/subreddits`);
        const q = query(subredditsRef);
        const querySnapshot = await getDocs(q);
  
        const subreddits = querySnapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
        }));
        return subreddits;
    } catch (e) {
        console.error("Error fetching user subreddits: ", e);
    }
}

/**
 * Removes a user from all subreddits in the subreddits collection.
 * 
 * @param {string} userId - The ID of the user to remove from all subreddits
 * @param {Array} subreddits - An array of subreddits to remove the user from
 */
export async function removeUserFromAllSubreddits(userId, subreddits){
    try {
        for (const subreddit of subreddits) {
        const subredditId = subreddit.id;
        const subredditUserRef = doc(db, `subreddits/${subredditId}/users`, userId);
        const subredditUserRefSnapshot = await getDoc(subredditUserRef);

        if (subredditUserRefSnapshot.exists()) {
            await deleteDoc(subredditUserRef);
        }

        // Check if that subreddit still has users
        checkSubredditUserList(subredditId) 
        }
    } catch (e) {
        console.error("Error removing user from all subreddits: ", e);
        throw new Error("Failed to remove user from all subreddits");
    }
}


/**
 * Toggles the subscription status of all subreddits for a specific user.
 * 
 * @param {string} userId - The ID of the user whose subreddits' subscription status is to be updated.
 * @param {boolean} subscribe - The new subscription status to apply to all user subreddits.
 * If `true`, the user will be subscribed to all their subreddits; if `false`, they will be unsubscribed.
 */
export async function updateAllSubredditSubscriptions(userId, subscribe) {
    try {
        const subredditsRef = collection(db, `users/${userId}/subreddits`);
        const q = query(subredditsRef);
        const querySnapshot = await getDocs(q);

        // Loop through each subreddit document and update the "starred" field
        querySnapshot.forEach(async (docSnapshot) => {
            const docRef = doc(db, `users/${userId}/subreddits`, docSnapshot.id);
            await updateDoc(docRef, { starred: subscribe });
        });

        // We then want to update the subreddits collection by adding the user
        // to all subreddits or removing them from all subreddits.
        if (subscribe) {
            addUserToAllSubreddits(userId, querySnapshot.docs);
        } else {
            removeUserFromAllSubreddits(userId, querySnapshot.docs);
        }

    } catch (e) {
        console.error("Error unsetting starred subreddits: ", e);
        throw new Error("Failed to update subreddits");
    }
}

/**
 * Updates the 'starred' status of a specific subreddit for a user.
 * 
 * @param {string} userId - The ID of the user whose subreddit's 'starred' status is to be updated.
 * @param {string} subredditId - The ID of the subreddit whose 'starred' status is to be updated.
 * @param {boolean} value - The new value to set for the 'starred' field (true or false).
 */
export async function updateStarredSubreddit(userId, subredditId, value) {
    const userSubredditsRef = doc(db, `users/${userId}/subreddits`, subredditId);
    const userSnapshot = await getDoc(userSubredditsRef);

    if (userSnapshot.exists()) {
        // First, update the starred status of the user's subreddit
        await updateDoc(userSubredditsRef, {
            starred: value,
        });
    }

    // We then want to check if the subreddit exists in the subreddits table
    const subredditRef = doc(db, `subreddits`, subredditId);
    const subredditSnap = await getDoc(subredditRef);
    if (!subredditSnap.exists()) {
        // If it doesn't exist, we create the subreddit
        await setDoc(subredditRef, {
            name: "r/" + subredditId,
        });
    }

    // We then want to update that subreddit's user list
    // This will help us determine if we need to continue summarizing the subreddit
    // Or delete the subreddit if no users are subscribed to it
    const subredditUsersRef = doc(db, `subreddits/${subredditId}/users`, userId);
    const subredditSnapshot = await getDoc(subredditUsersRef);

    if (subredditSnapshot.exists() && !value) {
        // If the user is unstarring the subreddit, remove them from the list
        await deleteDoc(subredditUsersRef);
        // Then check if the subreddit still has subscribed users
        checkSubredditUserList(subredditId);
    } else if (!subredditSnapshot.exists() && value) {
        // If the user is starring the subreddit, add them to the list
        await setDoc(subredditUsersRef, {});
    }
}

/**
 * Updates the email address of a specific user in the Firestore database.
 * 
 * @param {string} userId - The ID of the user whose email is to be updated.
 * @param {string} email - The new email address to set for the user.
 */
export async function updateUserEmail(userId, email) {
    const userRef = doc(db, 'users', userId);
    const userSnapshot = await getDoc(userRef);

    if (userSnapshot.exists()) {
        await updateDoc(userRef, {
            email: email,
          });
    }
}




