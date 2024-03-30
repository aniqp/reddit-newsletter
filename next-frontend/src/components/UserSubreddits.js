"use client"
import { useState, useEffect } from 'react'

const UserSubreddits = () => {
    const [subreddits, setSubreddits] = useState([])

    // useEffect(() => {
    //     const fetchSubreddits = async () => {
    //         try {
    //             const response = await fetch(`${process.env.REACT_APP_GO_ENDPOINT}/subreddits`);
    //             const res = await response.json();
    //             setSubreddits(res.data.children);
    //         } catch (error) {
    //             console.error("Failed to fetch subreddits:", error);
    //         }
    //     };
    //     fetchSubreddits();
    // }, []); 
    
    return (
    <div className='w-full overflow-scroll'>
        <div className='text-3xl font-bold mb-5'>Subscribed Subreddits</div>
        <div className='flex'>
            <div className='w-8/12'>
                {subreddits.map((subreddit, index) => (
                    <div className={`${index === 0 ? "" : "my-3"} ${index === (subreddits.length - 1) ? "mb-10" : ""}`}>
                        <SubredditCard subreddit={subreddit.data} />
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
                <div className='font-bold'>{subreddit.display_name_prefixed}</div>
                <div>
                    {subreddit.public_description}
                </div>
            </div>
            <div>
                <label htmlFor={`subscribe-${subreddit.display_name_prefixed}`} className='flex items-center cursor-pointer'>
                    <input
                        id={`subscribe-${subreddit.display_name_prefixed}`}
                        type="checkbox"
                        checked={true}
                        // onChange={(e) => onSubscriptionChange(e.target.checked, subreddit)}
                        className='mr-2'
                    />
                </label>
            </div>
        </div>
    )
}

export default UserSubreddits
