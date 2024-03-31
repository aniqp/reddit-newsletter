import UserSubreddits from '@/components/UserSubreddits'
import AccountSettings from '@/components/AccountSettings';
import React from 'react';


export default function SubredditPage({ params }) { 
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
