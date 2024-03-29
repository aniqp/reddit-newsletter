/** @type {import('next').NextConfig} */
import dotenv from 'dotenv';

// Load environment variables from .env file
dotenv.config();

const nextConfig = {
    env: {
      CLIENT_ID: process.env.CLIENT_ID,
      CLIENT_SECRET: process.env.CLIENT_SECRET,
      REDIRECT_URI: process.env.REDIRECT_URI,
      USER_AGENT: process.env.USER_AGENT,
      STATE: process.env.STATE,
    }
  };
  
export default nextConfig;
