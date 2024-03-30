import { headers } from 'next/headers'
import { NextResponse } from 'next/server'

export async function GET (req, res) {
    const headersList = headers()
    const accessToken = headersList.get('access-token')

    // Get list of user's subreddit from Reddit
    const userSubreddits = await fetch("https://oauth.reddit.com/subreddits/mine/subscriber", {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${accessToken}`
        }
    });

    const body = await userSubreddits.json();
    
    const subreddits = body.data.children.map(item => ({
        display_name: item.data.display_name,
        description: item.data.description
    }));

    return NextResponse.json(subreddits);
}
