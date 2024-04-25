import UserSubreddits from '@/components/UserSubreddits'
import AccountSettings from '@/components/AccountSettings';

export default function SubredditPage({ params }) { 
  return (
    <div className='flex'>
        <div className='w-full h-screen flex flex-col lg:flex-row pl-20 pt-10'>
          <div className='lg:w-1/3'>
            <AccountSettings />
          </div>
          <div className='mt-16 lg:mt-0 lg:w-2/3'>
            <UserSubreddits userId={params.id} />
          </div>
        </div>
    </div>
  )
}
