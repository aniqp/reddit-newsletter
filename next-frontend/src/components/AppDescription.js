"use client"
import Image from 'next/image'

const AppDescription = () => {
  const handleLogin = async () => {
    window.open(`https://reddit.com/api/v1/authorize?client_id=${process.env.CLIENT_ID}&response_type=code&state=${process.env.STATE}&redirect_uri=${process.env.REDIRECT_URI}&duration=permanent&scope=identity,mysubreddits,subscribe,save,read`, "_self")
}
  return (
    <div className='h-screen flex flex-col justify-center'>
        <div className='mx-auto flex text-center text-5xl mb-8 font-bold'>
            Stay in the loop with your subreddits.
        </div>
        <div className='text-md text-center opacity-50'>
        By connecting your Reddit account, we'll send you a personalized daily email summarizing the hot posts from your selected subreddits. Stay updated with the content that matters most to you, effortlessly!
        </div>
        <div className='flex justify-center'>
          <button onClick={handleLogin} className="mt-10 btn email-login-btn w-3/5">
              <div className="flex items-center">
              <Image className="mr-3" src="reddit-logo2.svg" width={35} height={35} />
              <span>Continue with Reddit</span>
              </div>
          </button>
        </div>
    </div>
  )
}

export default AppDescription
