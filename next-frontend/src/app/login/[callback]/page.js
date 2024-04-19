"use client"

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSearchParams } from 'next/navigation'
import { useDispatch } from 'react-redux';
import { 
  checkUserExists, 
  addUserSubreddits, 
  getUserSubreddits, 
  addNewUser,
  fetchUserEmail } from '@/db';
import { setUserId, setEmail, setSubreddits, setUserData } from '@/redux/slice';


export default function LoginCallbackPage() {
    const dispatch = useDispatch();

    const searchParams = useSearchParams()
    const router = useRouter();

    const code = searchParams.get('code')
    const state = searchParams.get('state')
    const error = searchParams.get('error')

    if (error || state !== process.env.STATE) {
      console.log('Error:', error);
      router.push('/login')
    } 

    useEffect(() => {
        async function fetchCallbackData() {
          const userApiResponse = await fetch(`/api/login/callback?code=${code}`);
          const user = await userApiResponse.json();
          dispatch(setUserData({id: user.id, accessToken: user.tokens.accessToken}));

          const subredditsApiResponse = await fetch(`/api/subreddits?accessToken=${user.tokens.accessToken}`);
          const subreddits = await subredditsApiResponse.json();

          console.log(subreddits)
          
          const userExists = await checkUserExists(user);
          if (!userExists) {
            await addNewUser(user);
            await addUserSubreddits(user.id, subreddits);
            const subredditsFromDb = await getUserSubreddits(user.id);

            dispatch(setSubreddits(subredditsFromDb));
            router.push(`/email/${user.id}`)
          } else {            
            const userFromDb = await fetchUserEmail(user.id);
            await addUserSubreddits(user.id, subreddits);
            const subredditsFromDb = await getUserSubreddits(user.id);
            dispatch(setSubreddits(subredditsFromDb));
            
            // If existing user has not set email, redirect to email page
            if (userFromDb.email) {
                dispatch(setEmail(userFromDb.email));
                router.push('/subreddits')
                return
            }
            router.push(`/email/${user.id}`)
          }
        }
        fetchCallbackData();
    });

  return (
    <div className='h-screen flex justify-center items-center'>
      <div className="loader"></div>
    </div>
  )
}
