"use client"

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation';
import { useSelector, useDispatch } from 'react-redux';
import { addUserSubreddits, getUserSubreddits, updateStarredSubreddit, updateAllSubredditSubscriptions } from '@/db';
import { setSubreddits } from '@/redux/slice';

const UserSubreddits = () => {
    const router = useRouter()
    const dispatch = useDispatch()
    const user = useSelector((state) => state.user)

    useEffect(() => {
        if (!user.id) {
            router.push('/login')
        }
        console.log(user.subreddits)
    }, [user.id, user.subreddits, router])

    const handleAllSubreddits = async (value) => {
        console.log(user.subreddits)
        await updateAllSubredditSubscriptions(user.id, value)
        dispatch({ type: 'user/updateStarredSubreddits', payload: value })
    }

    const handleUpdatingSubreddits = async () => {
        const subredditsApiResponse = await fetch(`/api/subreddits?accessToken=${user.accessToken}`);
        const subreddits = await subredditsApiResponse.json();
        await addUserSubreddits(user.id, subreddits);
        const subredditsFromDb = await getUserSubreddits(user.id);
        dispatch(setSubreddits(subredditsFromDb));
    }

    if (!user.id) {
        return null
    }

    return (
    <div className='w-full overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Subscribed Subreddits</div>
        <div className='flex'>
            <div className='w-8/12'>
                {user.subreddits.map((subreddit, index) => (
                    <div key={index} className={`${index === 0 ? "" : "my-3"} ${index === (user.subreddits.length - 1) ? "mb-10" : ""}`}>
                        <SubredditCard subreddit={subreddit} user={user} />
                    </div>
                ))}
            </div>
            <div className='flex flex-col mx-auto w-64'>
                <button onClick={handleUpdatingSubreddits} className='btn mb-5 bg-blue-500 text-white rounded-md'>Update Subreddits</button>
                <button className='unsubscribe btn rounded-md' onClick={()=> {handleAllSubreddits(false)}}>Unsubscribe from all</button>
            </div>
        </div>
    </div>
  )
}

const SubredditCard = ({ subreddit, user }) => {
    const dispatch = useDispatch()
    const [checked, setChecked] = useState(subreddit.starred)

    // Update the checked state when the starred value changes
    // This is needed when we unsubscribe or subscribe to all subreddits.
    useEffect(() => {
        setChecked(subreddit.starred)
    }, [subreddit.starred])

    const handleCheckmark = () => {
        updateStarredSubreddit(user.id, subreddit.id, !checked)
        dispatch({ type: 'user/updateStarredSubreddit', payload: { subredditId: subreddit.id, starred: !checked }})
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
                            onChange={handleCheckmark}
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
