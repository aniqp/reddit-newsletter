"use client"

import React from 'react'
import { useSelector } from 'react-redux';
import { deleteUser } from '@/db';
import { useRouter } from 'next/navigation';

export default function AccountSettings() {
  const user = useSelector((state) => state.user)

  if (!user.id) {
    return null
  }

  return (
    <div className='w-10/12 overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Account Settings</div>
        <div className='flex'>
            <div className='w-full'>
              <Profile user={user} />
              <DeleteAccount user={user} />
            </div>
        </div>
    </div>
  )
}

const Profile = ({ user }) => {
    return(
        <div className='account-card shadow-md mb-10'>
          <div className='text-lg font-semibold account-card-header pb-2 mb-3'>
            Reddit Account
          </div>
          <div className='flex justify-between'>
            <div className='flex flex-col'>
              <div className='mb-1 opacity-30'>
                {user.id}
              </div>
            </div>
          </div>
          <div className='text-lg mt-8 font-semibold account-card-header pb-2 mb-3'>
            Email
          </div>
          <div className='flex justify-between'>
            <div className='flex flex-col'>
              <div className='mb-1 opacity-30'>
                {user.email}
              </div>
            </div>
          </div>
        </div>
    )
}

const DeleteAccount = ({ user }) => {
  const router = useRouter();

  const handleDeleteAccount = () => {
    console.log('Deleting account')
    deleteUser(user.id)
    router.push('/login')
  }
  return(
      <div className='account-card shadow-md'>
        <div className='text-sm text-red-600 font-semibold account-card-header pb-2 mb-3'>
          Delete Account
        </div>
        <div className='flex justify-between'>
            <div className='opacity-75 text-sm'>
            You are about to permanently delete your account. This action cannot be undone. Upon confirmation, all of your account data, including your profile, preferences, and associated content, will be irreversibly removed.
            </div>
            
        </div>
        <div className='w-full mt-5'>
            <button onClick={handleDeleteAccount} className='delete-account'>Delete Account</button>
        </div>
      </div>
  )
}
