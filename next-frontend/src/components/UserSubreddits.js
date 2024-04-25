"use client"

import { useEffect, useMemo } from 'react'
import { useRouter } from 'next/navigation';
import { useSelector, useDispatch } from 'react-redux';
import { addUserSubreddits, getUserSubreddits, updateStarredSubreddit, updateAllSubredditSubscriptions } from '@/db';
import { setSubreddits } from '@/redux/slice';
import { Table, Checkbox, Tooltip } from 'antd';
import Image from 'next/image';

const UserSubreddits = () => {
    const router = useRouter()
    const dispatch = useDispatch()
    const user = useSelector((state) => state.user)

    useEffect(() => {
        if (!user.id) {
            router.push('/login')
        }
    }, [user.id, user.subreddits, router])

    if (!user.id) {
        return null
    }

    const columns = useMemo(
        () => [
            {
            title: 'Subreddit',
            dataIndex: 'display_name_prefixed',
            },
            {
                title: 'Description',
                dataIndex: 'description',
            },
            {
                title: 'Subscribed',
                dataIndex: 'starred',
                render: (value, record) => (
                    <Checkbox
                      checked={value}
                      onChange={()=>{handleCheckbox(record)}}
                    />
                )
            }
        ],
    [],
    );

    const handleCheckbox = async (record) => {
        console.log(record)
        const checked = record.starred;
        updateStarredSubreddit(user.id, record.id, !checked);
        dispatch({ type: 'user/updateStarredSubreddit', payload: { subredditId: record.id, starred: !checked }});
    }

    const rowSelection = {
        onChange: (selectedRowKeys, selectedRows) => {
        // console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
        selectedRows.forEach(row => {
            console.log(row.id)
            const subredditId = row.id
            // updateStarredSubreddit(user.id, subredditId, !row.checked)
            // dispatch({ type: 'user/updateStarredSubreddit', payload: { subredditId: subredditId, starred: !row.checked }})
          });
        },
        getCheckboxProps: (record) => ({
            name: record.display_name_prefixed,
            checked: record.starred === 'true' ? true : false,
        }),
    };    

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

    return (
    <div className='w-full overflow-scroll'>
        <div className='flex'>
            <div className='w-10/12'>
                <div className='flex w-full'>
                    <div className='w-1/2 text-3xl font-semibold mb-5'>Your Subreddits</div>
                    <div className='w-1/2 flex justify-end mb-4'>
                        <Tooltip title="Refresh Subreddits">
                            <button className='circle-icon flex justify-center items-center' onClick={handleUpdatingSubreddits}>
                                <Image className="opacity-50 hover:opacity-55" src="/refresh.png" alt="Refresh Subreddits" width={25} height={25} />
                            </button>
                        </Tooltip>
                        <Tooltip title="Subscribe to all">
                            <button className='circle-icon flex justify-center items-center' onClick={()=> {handleAllSubreddits(true)}}>
                                <Image className="opacity-50 hover:opacity-55" src="/subscribe.png" alt="Subscribe to All" width={30} height={30} />
                            </button>
                        </Tooltip>
                        <Tooltip title="Unsubscribe from all">
                            <button className='circle-icon flex justify-center items-center' onClick={()=> {handleAllSubreddits(false)}}>
                                <Image className="opacity-50 hover:opacity-55" src="/unsubscribe.png" alt="Unsubscribe to All" width={35} height={35} />
                            </button>
                        </Tooltip>
                    </div>
                </div>
                <Table
                    bordered
                    // rowSelection={{
                    // type: 'checkbox',
                    // ...rowSelection,
                    // }}
                    columns={columns}
                    dataSource={user.subreddits.map(subreddit => ({ ...subreddit, key: subreddit.id }))}
                    pagination={{
                        position: ['none', 'bottomCenter'],
                    }}
                />
            </div>
        </div>
    </div>
  )
}

export default UserSubreddits
