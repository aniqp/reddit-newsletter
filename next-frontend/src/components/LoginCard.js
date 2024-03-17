import React from 'react'
import Image from "next/image";
import Link from "next/link";

const LoginCard = () => {
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
            <Link href="/subreddits">
                <button className="btn email-login-btn">Log in</button>
            </Link>
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
