"use client"
import React from 'react'
import Image from "next/image";
import { redirect } from 'next/navigation'

const LoginCard = () => {
    const handleLogin = async () => {
        window.open(`https://reddit.com/api/v1/authorize?client_id=${process.env.CLIENT_ID}&response_type=code&state=${process.env.STATE}&redirect_uri=${process.env.REDIRECT_URI}&duration=permanent&scope=identity,subscribe,save`, "_self")
    }


  return (
    <div className="w-full flex justify-center">
        <div className="w-8/12 login-card">
            <div className="text-3xl font-semibold text-center mb-8">Account Login</div>
            <div className="mb-5">
                <input className="login" type="text" placeholder="Email" />
            </div>
            <div className="mb-1">
                <input className="login" type="text" placeholder={"Password"} />
            </div>
            <a className="flex justify-end text-sm b-50">
                <span>Forgot password?</span>
            </a>
            <a onClick={handleLogin}>
                <button className="btn email-login-btn">Log in</button>
            </a>
            <div className="line-or-line">
                <span className="line"></span>
                <span className="or-text">OR</span>
                <span className="line"></span>
            </div>
            <button className="btn reddit-login-btn">
                <div className="flex items-center">
                <Image className="mr-3" src="reddit-logo-orange.svg" width={35} height={35} />
                <span className="">Continue with Reddit</span>
                </div>
            </button>
            <div className="mt-7 text-md text-center">
                <span className='b-50'>Don't have an account?&nbsp;</span>
                <a style={{color:"#ff6730"}}>Sign up!</a>
            </div>
        </div>
    </div>
  )
}

export default LoginCard
