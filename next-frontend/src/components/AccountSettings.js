"use client"

import { useState } from 'react'
import { useSelector } from 'react-redux';
import { deleteUser, updateUserEmail } from '@/db';
import { useRouter } from 'next/navigation';
import { Input, Avatar, Modal, message } from 'antd';
import Image from 'next/image';
import { useDispatch } from 'react-redux';
import { setEmail } from '@/redux/slice';


export default function AccountDetails() {
  const user = useSelector((state) => state.user)

  if (!user.id) {
    return null
  }

  return (
    <div className='w-10/12 overflow-scroll'>
        <div className='text-3xl font-semibold mb-5'>Account Details</div>
        <div className='flex'>
            <div className='w-full'>
              <Profile user={user} />
              <DeleteAccount user={user} />
            </div>
        </div>
    </div>
  )
}

const Profile = ({ user }) => {
  const dispatch = useDispatch();
  const [updatedEmail, setUpdatedEmail] = useState(user.email);
  const [disabled, setDisabled] = useState(true);
  const [messageApi, contextHolder] = message.useMessage();

  const success = () => {
    messageApi.success('Email was successfully changed!');
  };

  const handleEmailChange = (e) => {
    const email = e.target.value;
    setUpdatedEmail(email);
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (emailRegex.test(email)){
      setDisabled(false);
    } else {
      setDisabled(true);
    }
  };

  const handleUpdate = (e) => {
    console.log('changing now')
    updateUserEmail(user.id, updatedEmail)
    dispatch(setEmail(updatedEmail))
    success();
  }

  return(
      <div className=''>
        <div className='flex items-center'>
          <Avatar
            style={{
              backgroundColor: '#fde3cf',
              color: '#f56a00',
              height: '64px',
              width: '64px',
            }}
          >
            <span className='text-3xl'>{user.id.charAt(0).toUpperCase()}</span>
          </Avatar>
          <div className='ml-5'>
            <div className='font-semibold text-lg'>
              {user.id}
            </div>
            <div className='font-light text-sm'>
            Joined August 2021
            </div>
          </div>
        </div>
        <div className='flex-col mt-5 mb-20'>
          <div className='font-semibold mb-2'>
            Email
          </div>
          <div className='flex w-full items-center'>
            <div className='w-2/3'>
              <Input 
              size="middle" 
              value={updatedEmail} 
              type="email" 
              onChange={handleEmailChange}
              placeholder="Enter email"
              />
            </div>
            <div className='ml-3 h-full'>
              <button 
              className={`email-update-btn text-sm ${disabled ? 'pointer-events-none' : ''}`}
              style={{backgroundColor: disabled ? '#f0f0f0' : '#f56a00'}}
              onClick={handleUpdate} 
              disabled={disabled}
              >Update</button>
            </div>
          </div>
        </div>
      </div>
  )
}


const DeleteAccount = ({ user }) => {
  const router = useRouter();

  const [isModalOpen, setIsModalOpen] = useState(false);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    deleteUser(user.id)
    router.push('/login')
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  return(
      <div className='account-card'>
        <div className='text-sm font-semibold account-card-header pb-3 mb-3'>
          Delete Account
        </div>
        <div className='flex justify-between'>
            <div className='opacity-75 text-sm'>
            You are about to permanently delete your account. Upon confirmation, all of your account data will be irreversibly removed.
            </div>
        </div>
        <div className='w-full mt-5'>
            <button onClick={showModal}className='delete-account-btn text-red-600 text-sm'>Delete Account</button>
        </div>
        <Modal open={isModalOpen} footer={[]} width={400}>
          <div className='flex flex-col justify-center items-center my-5 mx-4'>
            <Image src="/bin.png" width={100} height={100} />
            <div className='text-2xl font-bold text-center my-5'>Are you sure you want to delete your account?</div>
            <button onClick={handleOk} className='delete-confirm-btn'>Yes, Delete </button>
            <button onClick={handleCancel} className='cancel-delete-btn mt-5'>Keep Account</button>
          </div>
        </Modal>
      </div>
  )
}
