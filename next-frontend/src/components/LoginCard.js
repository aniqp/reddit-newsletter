import React from 'react'
import Image from "next/image";

const LoginCard = () => {
    const handleLogin = async () => {
        window.open(`https://reddit.com/api/v1/authorize?client_id=${process.env.CLIENT_ID}&response_type=code&state=${process.env.STATE}&redirect_uri=${process.env.REDIRECT_URI}&duration=permanent&scope=identity,mysubreddits,subscribe,save`, "_self")
    }

  return (
    <div className="w-full flex justify-center">
        <div className="w-8/12 login-card">
            <div className="text-3xl font-semibold text-center mb-10">Account Login</div>
            <button onClick={handleLogin} className="btn email-login-btn">
                <div className="flex items-center">
                <Image className="mr-3" src="reddit-logo2.svg" width={35} height={35} />
                <span className="">Login with Reddit</span>
                </div>
            </button>
        </div>
    </div>
  )
}

export default LoginCard
