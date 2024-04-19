import { NextResponse } from 'next/server';

export async function GET(req, res) {
    const queryParams = Object.fromEntries(req.nextUrl.searchParams);
    const { accessToken } = queryParams;

    async function fetchSubreddits(accessToken, retryCount = 0) {
        const response = await fetch("https://oauth.reddit.com/subreddits/mine/subscriber", {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${accessToken}`
            }
        });

        if (response.status === 429 && retryCount < 3) { // Check if rate limited and limit retry attempts
            const retryAfter = parseInt(response.headers.get('Retry-After'), 10);
            console.log(`Rate limited. Retrying after ${retryAfter} seconds.`);
            await new Promise(resolve => setTimeout(resolve, retryAfter * 1000));
            return fetchSubreddits(accessToken, retryCount + 1); // Retry the request
        }

        return response;
    }

    try {
        const userSubredditsResponse = await fetchSubreddits(accessToken);
        if (!userSubredditsResponse.ok) {
            throw new Error(`HTTP error! Status: ${userSubredditsResponse.status}`);
        }
        const body = await userSubredditsResponse.json();

        const subreddits = body.data.children.map(item => ({
            display_name_prefixed: item.data.display_name_prefixed,
            display_name: item.data.display_name,
            description: item.data.public_description
        }));

        return NextResponse.json(subreddits);
    } catch (error) {
        return res.status(500).json({ message: error.message });
    }
}
