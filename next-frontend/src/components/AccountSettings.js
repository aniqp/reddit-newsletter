import React from 'react'

export default function AccountSettings() {
  return (
    <div className='w-10/12 overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Account Settings</div>
        <div className='flex'>
            <div className='w-full'>
              <Profile />
              <LinkedAccounts />
              <DeleteAccount />
            </div>
        </div>
    </div>
  )
}

const Profile = () => {
    return(
        <div className='account-card shadow-md mb-10'>
          <div className='text-lg font-semibold account-card-header pb-2 mb-3'>
            Personal Information
          </div>
          <div className='flex justify-between'>
            <div className='flex flex-col'>
              <div className='mb-1'>
                John Doe
              </div>
              <div className='opacity-30'>
              johndoe@gmail.com
              </div>
            </div>
            <div className='flex items-center'>
              <button className='account mr-3'>Edit Profile</button>
              <button className='account'>Change Password</button>
            </div>
          </div>
        </div>
    )
}


const LinkedAccounts = () => {
  return(
      <div className='account-card shadow-md mb-10'>
        <div className='text-lg font-semibold account-card-header pb-2 mb-3'>
         Linked Accounts
        </div>
      </div>
  )
}

const DeleteAccount = () => {
  return(
      <div className='account-card shadow-md'>
        <div className='text-sm text-red-600 font-semibold account-card-header pb-2 mb-3'>
          Delete Account
        </div>
        <div className='flex justify-between'>
            <div className='w-9/12 opacity-75'>
            You are about to permanently delete your account. This action cannot be undone. Upon confirmation, all of your account data, including your profile, preferences, and associated content, will be irreversibly removed.
            </div>
            <div>
              <button className='account'>Delete Account</button>
            </div>
        </div>
      </div>
  )
}
