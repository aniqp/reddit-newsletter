"use client"

import { useState, useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux';
import { setSubreddits } from '@/redux/slice'
import { useRouter } from 'next/navigation';
import { getUserSubreddits } from '@/db';

const UserSubreddits = ({ userId }) => {
    const router = useRouter()
    const dispatch = useDispatch()
    const user = useSelector((state) => state.user.account)
    const [userSubreddits, setUserSubreddits] = useState([])

    useEffect(() => {
        async function fetchData() {
            const res = await getUserSubreddits(userId); // Await the asynchronous call
            console.log(res);
            if (res) {
                dispatch(setSubreddits(res));
                setUserSubreddits(res);
            }
        }
        fetchData();
    }, [userId, dispatch])

    return (
    <div className='w-full overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Subscribed Subreddits</div>
        <div className='flex'>
            <div className='w-8/12'>
                {userSubreddits.map((subreddit, index) => (
                    <div className={`${index === 0 ? "" : "my-3"} ${index === (userSubreddits.length - 1) ? "mb-10" : ""}`}>
                        <SubredditCard subreddit={subreddit} user={user} />
                    </div>
                ))}
            </div>
            <div className='flex flex-col mx-auto w-64'>
                <button className='btn mb-5 bg-blue-500 text-white rounded-md'>Update Subreddits</button>
                <button className='btn mb-5 bg-blue-500 text-white rounded-md'>Save Changes</button>
                <button className='unsubscribe btn rounded-md'>Unsubscribe from all</button>
            </div>
        </div>
    </div>
  )
}

const SubredditCard = ({ subreddit, user }) => {
    return (
        <div className={`subreddit-card`} onClick={console.log()}>
            <div>
                <div className='font-bold'>{subreddit.id}</div>
                <div>
                    {subreddit.description}
                </div>
            </div>
            <div>
                <label className='flex items-center cursor-pointer'>
                    <input
                        type="checkbox"
                        disabled={true ? user.email === null : false}
                        checked={true ? (user.email !== null) : false}
                        // onChange={(e) => onSubscriptionChange(e.target.checked, subreddit)}
                        className='mr-2'
                    />
                </label>
            </div>
        </div>
    )
}

export default UserSubreddits
