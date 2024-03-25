package reddit_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	reddit_models "reddit-newsletter/pkg/models"

	"github.com/gin-gonic/gin"
)

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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		fmt.Println("Response body:", string(body))
	}

	return body, nil
}

func getHotPosts(accessToken string, subreddit string, client *http.Client) ([]byte, error) {
	// Create a new GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/top?sort=hot&t=day&limit=10", subreddit), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add authorization header with OAuth2 access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response:", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		fmt.Println("Response body:", string(body))
	}

	return body, nil
}

func getComments(accessToken string, subreddit string, post_id string, client *http.Client) ([]byte, error) {
	// Create a new GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/comments/%s/top?sort=hot&t=day&limit=20&depth=3&truncate=25", subreddit, post_id), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add authorization header with OAuth2 access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response:", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		fmt.Println("Response body:", string(body))
	}

	return body, nil

}

func processComments(SubredditHotPosts reddit_models.SubredditHotPosts, rc RedditClient, subreddit string) []reddit_models.HotPostWithComments {
	var hotPostsWithComments []reddit_models.HotPostWithComments

	for _, child := range SubredditHotPosts.Data.Children {
		hotPost := reddit_models.HotPostWithComments{
			Title:    child.Data.Title,
			SelfText: child.Data.Selftext,
		}
		body, err := getComments(rc.accessToken, subreddit, child.Data.Id, rc.client)

		var subredditComments []reddit_models.SubredditComment
		err = json.Unmarshal(body, &subredditComments)
		if err != nil {
			fmt.Println("Error unmarshalling comments:", err)
			// Handle the error as needed, maybe skip this post or return an error response
			continue
		}

		hotPost.Comments = subredditComments[1]
		hotPostsWithComments = append(hotPostsWithComments, hotPost)

	}
	return hotPostsWithComments
}

func (rc *RedditClient) GetSubredditsHandler(c *gin.Context) {
	// Implement the logic to fetch subreddits
	body, err := getSubreddits(rc.accessToken, rc.client)
	var SubredditResponse reddit_models.SubredditResponse
	err = json.Unmarshal(body, &SubredditResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	c.JSON(http.StatusOK, SubredditResponse)
}

func (rc *RedditClient) GetHotPostsAndCommentsHandler(c *gin.Context) {
	// Implement the logic to fetch hot posts and comments
	subreddit := c.Query("subreddit")

	body, err := getHotPosts(rc.accessToken, subreddit, rc.client)

	var SubredditHotPosts reddit_models.SubredditHotPosts
	err = json.Unmarshal(body, &SubredditHotPosts)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		// If user with specified ID is not found, return a 404 error
		c.JSON(http.StatusNotFound, gin.H{"error": "subreddit hot posts failed"})
		return
	}

	hotPostsWithComments := processComments(SubredditHotPosts, *rc, subreddit)

	c.JSON(http.StatusOK, hotPostsWithComments)
}
