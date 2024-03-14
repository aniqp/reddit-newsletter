"use client"
import React from 'react'

const subreddits = [
    { name: 'r/aww', description: 'A subreddit for cute and cuddly pictures', isSubscribed: true },
    { name: 'r/reactjs', description: 'A subreddit for ReactJS news and discussion', isSubscribed: false },
    { name: 'r/programming', description: 'A subreddit for news and discussion about programming', isSubscribed: true },
    { name: 'r/gaming', description: 'A subreddit for news and discussion about gaming', isSubscribed: false },
    { name: 'r/movies', description: 'A subreddit for news and discussion about movies', isSubscribed: false },
    { name: 'r/food', description: 'A subreddit for news and discussion about food', isSubscribed: false },
    { name: 'r/askreddit', description: 'A subreddit for news and discussion about anything', isSubscribed: true },
    { name: 'r/askscience', description: 'A subreddit for news and discussion about science', isSubscribed: true },
    { name: 'r/askhistorians', description: 'A subreddit for news and discussion about history', isSubscribed: true },
    { name: 'r/askmen', description: 'A subreddit for news and discussion', isSubscribed: false}
]


const UserSubreddits = () => {
  return (
    <div>
        <div className='text-2xl font-bold'>Your Subscribed Subreddits</div>
        <div className='subreddit-list bg-slate-200'>
            {subreddits.map(subreddit => (
                <div className='mx-5 my-5'>
                    <SubredditCard subreddit={subreddit} />
                </div>
            ))}
        </div>
        <div className='subreddit-list-btn-group flex justify-around py-5'>
            <button className='btn bg-gray-100 text-gray-400 rounded-md'>Unsubscribe from all</button>
            <button className='btn bg-blue-500 text-white rounded-md'>Save Changes</button>
        </div>
    </div>
  )
}

const SubredditCard = ({ subreddit }) => {
    return (
        <div onClick={console.log('switch')} className='hover:cursor-pointer flex justify-between bg-white rounded-md p-5'>
            <div>
                <div className='font-bold'>{subreddit.name}</div>
                <div>
                    {subreddit.description}
                </div>
            </div>
            <div>
                <label htmlFor={`subscribe-${subreddit.name}`} className='flex items-center cursor-pointer'>
                    <input
                        id={`subscribe-${subreddit.name}`}
                        type="checkbox"
                        checked={subreddit.isSubscribed}
                        // onChange={(e) => onSubscriptionChange(e.target.checked, subreddit)}
                        className='mr-2'
                    />
                </label>
            </div>
        </div>
    )
}

export default UserSubreddits
