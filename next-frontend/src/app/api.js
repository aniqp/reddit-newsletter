const BASEURL = 'http://localhost:8080'
const SUBREDDIT_ENDPOINT = `/subreddits`

export const subreddits = async () => {
    const endpoint = `${BASEURL}${SUBREDDIT_ENDPOINT}`

    return await fetch(endpoint)
    .then(response => response.json())
    .catch(error => console.log(error));
}
