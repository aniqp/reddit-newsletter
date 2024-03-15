"use client"
import React from 'react'
import Image from 'next/image'
import Link from 'next/link'

const tabs = [
    {id: 1, label: 'Subreddit Preferences', icon: '/envelope.png'},
    {id: 2, label: 'Explore', icon: '/search.png'},
    {id: 3, label: 'Account Settings', icon: '/user.png'},
]

const Sidebar = ({ activeTab, setActiveTab }) => {
    return (
        <div className='sidebar'>
            <div className='mt-10'>
                {tabs.map((tab, index) => (
                    <Tab
                    key={index} 
                    tab={tab} 
                    isActive={tab.id === activeTab} 
                    onClick={() => setActiveTab(tab.id)} />
                ))}
            </div>
            <div className='mb-2'>
                <Link href="/">
                    <Tab 
                    tab={{label: 'Logout', icon: '/sign-out.png'}} 
                    isActive={false}
                    onClick={()=>{}}/>
                </Link>
            </div>
        </div>
    )
}

const Tab = ({ tab, isActive, onClick }) => {
    const activeStyle = isActive ? 'active-tab' : '';

    return(
        <div className={`flex items-center mb-3 py-5 px-4 cursor-pointer ${activeStyle}`} onClick={onClick}>
            <Image src={tab.icon} height={28} width={28} />
            <span className='ml-4 text-lg'>{tab.label}</span>
        </div>
    )
}

export default Sidebar
