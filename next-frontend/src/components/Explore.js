"use client"
import { useState } from 'react'

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

const Explore = () => {
    const [searchResults, setSearchResults] = useState([]);
    return (
        <div className='w-10/12 overflow-scroll'>
            <div className='text-3xl font-bold mb-5'>Explore Subreddits</div>
            <input className='search' />
            <div className='flex'>
                <div className='w-full'>
                </div>
            </div>
        </div>
    )
}

export default Explore
