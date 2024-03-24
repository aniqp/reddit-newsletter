"use client"
import { useState } from 'react';
import UserSubreddits from '@/components/UserSubreddits'
import Sidebar from '@/components/Sidebar';
import Explore from '@/components/Explore';
import AccountSettings from '@/components/AccountSettings';
import React from 'react';
import { subreddits } from '@/app/api'


export default function SubredditPage() {
  const [activeTab, setActiveTab] = useState(1);

  React.useEffect(() => {
    const fetchData = async () => {
      const res = await subreddits()
      // const data = await res.json()
      console.log(res)
    }
    fetchData()
  }, [])

  return (
    <div className='flex'>
        <Sidebar activeTab={activeTab} setActiveTab={setActiveTab}/>
        <div className='w-full h-screen flex pl-20 pt-10'>
          {activeTab == 1 && <UserSubreddits />}
          {activeTab == 2 && <Explore />}
          {activeTab == 3 && <AccountSettings />}
        </div>
    </div>
  )
}
