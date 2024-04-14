import { NextResponse } from 'next/server'

export async function GET (req, res) {
    const queryParams = Object.fromEntries(req.nextUrl.searchParams);
    const { accessToken } = queryParams;

    // Get list of user's subreddit from Reddit
    const userSubreddits = await fetch("https://oauth.reddit.com/subreddits/mine/subscriber", {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${accessToken}`
        }
    });

    const body = await userSubreddits.json();
    
    const subreddits = body.data.children.map(item => ({
        display_name_prefixed: item.data.display_name_prefixed,
        display_name: item.data.display_name,
        description: item.data.public_description
    }));

    return NextResponse.json(subreddits, {status: 200});
}
