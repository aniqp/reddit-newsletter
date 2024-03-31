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

  return (
    <div className="w-full h-screen flex justify-center items-center">
    <div className="w-2/5 login-card">
        <div className="text-3xl font-semibold text-center mb-5">
            Welcome, {user.id}!
        </div>
        <div className='mb-10 text-sm'>
          Thank you for signing up. To ensure you receive our daily newsletters, please enter your email address below.
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
        <button onClick={handleEmailSubmit} className="btn email-login-btn">
            <div className="flex items-center">
              <span>Confirm Email</span>
            </div>
        </button>
        <button className="btn email-login-btn">
            <div className="flex items-center">
              <span>Cancel</span>
            </div>
        </button>
    </div>
  </div>
  )
}

export default EmailPage
