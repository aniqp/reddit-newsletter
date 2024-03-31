import { NextResponse } from 'next/server'

export async function GET (req, res) {
    const queryParams = Object.fromEntries(req.nextUrl.searchParams);
    const { code } = queryParams;

    // Basic HTTP Authorization Header 
    const encodedHeader = Buffer.from(`${process.env.CLIENT_ID}:${process.env.CLIENT_SECRET}`).toString("base64");

    // Retrieve access token from Reddit
    const responseToken = await fetch (`https://www.reddit.com/api/v1/access_token`, {
        method: 'POST',
        headers: {
            'Authorization': `Basic ${encodedHeader}`,
            'Content-Type': 'application/x-www-form-urlencoded',
            'User-Agent': process.env.USER_AGENT,
        },
        body: `grant_type=authorization_code&code=${code}&redirect_uri=${process.env.REDIRECT_URI}`
    });
    
    const tokenBody = await responseToken.json();

    // Get user information from Reddit
    const responseUser = await fetch(`https://oauth.reddit.com/api/v1/me`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${tokenBody.access_token}`
        }
    });

    if (!responseUser.ok) {
      throw new Error(`Failed to retrieve user information: ${responseUser.statusText}`);
    }

    const userBody = await responseUser.json();
    
    const user = {
      id: userBody.name,
      tokens: {
        accessToken: tokenBody.access_token,
        refreshToken: tokenBody.refresh_token
      }
    }

    return NextResponse.json(user, {status: 200});
  }
