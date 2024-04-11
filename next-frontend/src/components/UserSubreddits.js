"use client"

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation';
import { useSelector, useDispatch } from 'react-redux';
import { updateStarred } from '@/db';

const UserSubreddits = () => {
    const router = useRouter()
    const dispatch = useDispatch()
    const user = useSelector((state) => state.user)

    useEffect(() => {
        if (!user.id) {
            router.push('/login')
        }
    }, [user.id])

    if (!user.id) {
        return null
    }

    return (
    <div className='w-full overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Subscribed Subreddits</div>
        <div className='flex'>
            <div className='w-8/12'>
                {user.subreddits.map((subreddit, index) => (
                    <div className={`${index === 0 ? "" : "my-3"} ${index === (user.subreddits.length - 1) ? "mb-10" : ""}`}>
                        <SubredditCard subreddit={subreddit} user={user} />
                    </div>
                ))}
            </div>
            <div className='flex flex-col mx-auto w-64'>
                <button className='btn mb-5 bg-blue-500 text-white rounded-md'>Update Subreddits</button>
                <button className='unsubscribe btn rounded-md'>Unsubscribe from all</button>
            </div>
        </div>
    </div>
  )
}

const SubredditCard = ({ subreddit, user }) => {
    const [checked, setChecked] = useState(subreddit.starred)
    const handleCheckmark = () => {
        updateStarred(user.id, subreddit.id, !checked)
        setChecked(!checked)
    }
    return (
    <a onClick={handleCheckmark}>
        <div className={`subreddit-card shadow-md ${checked ? '' : ''}`} onClick={console.log()}>
            <div>
                <div className='font-bold'>{subreddit.id}</div>
                <div className='opacity-40 mr-10'>
                    {subreddit.description}
                </div>
            </div>
            <div className='flex items-center'>
                <label className='container flex items-center cursor-pointer'>
                    <input 
                        className='mr-2'
                        checked={checked} 
                        type="checkbox"
                        // onChange={handleCheckmark}
                        disabled={true ? user.email === null : false} 
                        />
                    <div class="checkmark"></div>
                </label>
            </div>
        </div>
    </a>
    )
}

export default UserSubreddits
