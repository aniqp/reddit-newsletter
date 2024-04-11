"use client"

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useDispatch, useSelector } from 'react-redux'
import { setEmail, setSubreddits } from '@/redux/slice';
import { updateUserEmail, addUserSubreddits, getUserSubreddits } from '@/db'

const EmailPage = () => {
  const router = useRouter()
  const dispatch = useDispatch()
  const user = useSelector(state => state.user)
  const [input, setInput] = useState('');

  const handleEmailChange = (e) => {
    setInput(e.target.value);
  };

  const handleEmailSubmit = async (e) => {
    dispatch(setEmail(input));
    await updateUserEmail(user.id, input);
    router.push('/subreddits');
  }

  const handleCancel = () => {

    router.push('/login');
  }

  if (!user.id) {
    router.push('/login');
  }

  return (
    <div className="login-background w-full h-screen flex justify-center items-center">
    <div className="xl:w-2/5 lg:w-2/5 md:w-3/5 w-4/5 login-card">
        <div className="text-3xl font-bold text-center mb-5">
            Welcome, {user.id}!
        </div>
        <div className='mb-10 text-sm opacity-50'>
          Thank you for signing up! To receive your daily newsletter, please enter your email address below.
        </div>
        <input
          type="email"
          id="email"
          value={input}
          onChange={handleEmailChange}
          required
          placeholder="Enter your email"
          style={{
              width: '100%',
              padding: '8px',
              borderRadius: '4px',
              border: '1px solid #ccc',
          }}
      />
      <div className='mt-5 mx-auto flex'>
          <button onClick={handleCancel} className="w-1/3 btn cancel-btn">
              <div className="flex items-center">
                <span>Cancel</span>
              </div>
          </button>
          <button onClick={handleEmailSubmit} className=" ml-5 w-2/3 btn email-login-btn">
              <div className="flex font-semibold items-center">
                <span>Confirm Email</span>
              </div>
          </button>
          
        </div>
    </div>
  </div>
  )
}

export default EmailPage
