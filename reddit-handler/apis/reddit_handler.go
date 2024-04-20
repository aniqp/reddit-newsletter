package reddit_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	reddit_models "reddit-newsletter/pkg/models"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedditClient defines a client for interacting with the Reddit API
type RedditClient struct {
	accessToken string
	client      *http.Client
}

// NewRedditClient creates a new RedditClient with the specified access token
func NewRedditClient(accessToken string) *RedditClient {
	return &RedditClient{
		accessToken: accessToken,
		client:      &http.Client{},
	}
}

func getSubreddits(accessToken string, client *http.Client) ([]byte, error) {
	// Create a new GET request
	req, err := http.NewRequest("GET", "https://oauth.reddit.com/subreddits/mine/subscriber", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add authorization header with OAuth2 access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Rate limit hit for GetSubreddits. Waiting before retrying...")
		resetTime := resp.Header.Get("Retry-After")
		waitTime, _ := strconv.Atoi(resetTime)
		time.Sleep(time.Duration(waitTime) * time.Second)
		return getSubreddits(accessToken, client)
	}

	// Read the response body
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error reading response:", err)
		}
		return body, nil
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Error response: %s\nResponse body: %s\n", resp.Status, string(body))
	return nil, fmt.Errorf("API request GetSubreddits failed with status: %s", resp.Status)
}

func getHotPosts(accessToken string, subreddit string, client *http.Client) ([]byte, error) {
	// Create a new GET request

	req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/top?sort=hot&t=day&limit=10", subreddit), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add authorization header with OAuth2 access token
	req.Header.Set("Authorization", "bearer "+accessToken)
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Rate limit hit for GetHotPosts. Waiting before retrying...")
		resetTime := resp.Header.Get("Retry-After")
		waitTime, _ := strconv.Atoi(resetTime)
		time.Sleep(time.Duration(waitTime) * time.Second)
		return getHotPosts(accessToken, subreddit, client)
	}

	// Read the response body
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error reading response:", err)
		}
		return body, nil
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Error response: %s\nResponse body: %s\n", resp.Status, string(body))
	return nil, fmt.Errorf("API request GetHotPosts failed with status: %s", resp.Status)
}

func getComments(accessToken string, subreddit string, post_id string, client *http.Client) ([]byte, error) {
	url := fmt.Sprintf("https://oauth.reddit.com/%s/comments/%s/top?sort=hot&t=day&limit=10&depth=1", subreddit, post_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Println("Rate limit hit for GetComments. Waiting before retrying...")
		resetTime := resp.Header.Get("Retry-After")
		waitTime, _ := strconv.Atoi(resetTime)
		time.Sleep(time.Duration(waitTime) * time.Second)
		return getComments(accessToken, subreddit, post_id, client)
	}

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response: %v", err)
		}
		fmt.Println("Got comments for 1 post")
		return body, nil
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Error response: %s\nResponse body: %s\n", resp.Status, string(body))
	return nil, fmt.Errorf("API request GetComments failed with status: %s", resp.Status)

}

func processComments(SubredditHotPosts reddit_models.SubredditHotPosts, rc RedditClient, subreddit string) []reddit_models.HotPostWithComments {
	var hotPostsWithComments []reddit_models.HotPostWithComments
	var wg sync.WaitGroup
	for _, child := range SubredditHotPosts.Data.Children {
		wg.Add(1)
		child := child
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			hotPost := reddit_models.HotPostWithComments{
				Title:    child.Data.Title,
				SelfText: child.Data.Selftext,
			}
			body, _ := getComments(rc.accessToken, subreddit, child.Data.Id, rc.client)

			var subredditComments []reddit_models.SubredditComment
			err := json.Unmarshal(body, &subredditComments)
			if err != nil {
				fmt.Println("Error unmarshalling comments:", err)
				return
			}

			hotPost.Comments = subredditComments[1]
			hotPostsWithComments = append(hotPostsWithComments, hotPost)
			time.Sleep(1 * time.Second)
		}()
	}
	wg.Wait()
	return hotPostsWithComments
}

func (rc *RedditClient) GetSubredditsResponse(accessToken string, client *http.Client) ([]byte, error) {
	// Implement the logic to fetch subreddits
	body, err := getSubreddits(accessToken, client)
	if err != nil {
		return nil, err
	}

	var subredditResponse reddit_models.SubredditResponse
	err = json.Unmarshal(body, &subredditResponse)
	if err != nil {
		return nil, err
	}

	formattedSubreddit, err := json.Marshal(subredditResponse)
	if err != nil {
		return nil, err
	}

	return formattedSubreddit, nil
}

func (rc *RedditClient) GetHotPostsAndCommentsResponse(subreddit string) (string, error) {
	// Implement the logic to fetch hot posts and comments
	body, err := getHotPosts(rc.accessToken, subreddit, rc.client)
	if err != nil {
		return "", err
	}

	var SubredditHotPosts reddit_models.SubredditHotPosts
	err = json.Unmarshal(body, &SubredditHotPosts)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response body: %v", err)
	}

	hotPostsWithComments := processComments(SubredditHotPosts, *rc, subreddit)
	subredditHotPostsWithComments := reddit_models.SubredditHotPostsWithComments{
		SubredditName: subreddit,
		HotPosts:      hotPostsWithComments,
	}
	subredditHotPosts, err := json.Marshal(subredditHotPostsWithComments)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON response: %v", err)
	}

	return string(subredditHotPosts), nil
}

func AddToRedisQueue(hotPostsJSON string, rdb *redis.Client) {
	if err := rdb.LPush(ctx, "hotPostsQueue", hotPostsJSON).Err(); err != nil {
		fmt.Println("Error pushing to Redis:", err)
		return
	}
}

// Temporary for debugging, this will be moved to the GPT summarizer
func ConsumeQueue(rdb *redis.Client) {
	for {
		// Using BRPop with a timeout of 0 blocks indefinitely until an item is available
		result, err := rdb.BRPop(ctx, 0*time.Second, "hotPostsQueue").Result()
		if err != nil {
			fmt.Println("Error popping from queue:", err)
			continue
		}

		// result[1] contains the popped item (the JSON string)
		jsonStr := result[1]

		var hotPosts reddit_models.SubredditHotPostsWithComments
		err = json.Unmarshal([]byte(jsonStr), &hotPosts)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		// Process the hot posts with comments
		fmt.Printf("Consumed and processing %d hot posts with comments\n", len(hotPosts.HotPosts))
		fmt.Printf("subredit: %s\n", hotPosts.SubredditName)
		// Here you can process the hotPosts as needed
		for _, post := range hotPosts.HotPosts {
			fmt.Printf("Post: %s, Comments: %d\n", post.Title, len(post.Comments.Data.Children))
		}
	}
}
