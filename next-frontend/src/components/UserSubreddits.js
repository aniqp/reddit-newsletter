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
    <div className='w-full overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Subscribed Subreddits</div>
        <div className='flex'>
            <div className='w-8/12'>
                {subreddits.map((subreddit, index) => (
                    <div className={`${index === 0 ? "" : "my-3"} ${index === (subreddits.length - 1) ? "mb-10" : ""}`}>
                        <SubredditCard subreddit={subreddit} />
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

const SubredditCard = ({ subreddit }) => {
    return (
        <div className='subreddit-card' onClick={console.log()}>
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
