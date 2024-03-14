import React from 'react'
import UserSubreddits from '@/components/UserSubreddits'

const SubredditPage = () => {
  return (
    <div className='h-screen w-2/3 flex items-center'>
        <div className='w-1/2'>
            {/* Profile */}
        </div>
        <div className='w-1/2'>
            <UserSubreddits />
        </div>
    </div>
  )
}

export default SubredditPage
