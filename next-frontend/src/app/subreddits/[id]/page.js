"use client"
import { useEffect } from 'react';
import UserSubreddits from '@/components/UserSubreddits'
import AccountSettings from '@/components/AccountSettings';
import React from 'react';
import { useRouter } from 'next/navigation';


export default function SubredditPage({ params }) {
  // const router = u();

  // useEffect(() => {
  //   console.log('inninini')
  //   const { state, code } = router.query;
  //   console.log(state, code)

  //   let user = null;
  //   async function fetchCallbackCode(code) {
  //     const response = await fetch(`/api/login/callback?code=${code}`);
  //     const user = await response.json();
  //     console.log(user);
  //   }
  //   fetchCallbackCode();

  // }, [router.isReady, router.query]);
 
  return (
    <div className='flex'>
        <div className='w-full h-screen flex pl-20 pt-10'>
          <div className='w-1/3'>
            <AccountSettings />
          </div>
          <div className='w-2/3'>
            <UserSubreddits userId={params.id} />
          </div>
        </div>
    </div>
  )
}
