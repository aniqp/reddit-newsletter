package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type SubredditResponse struct {
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				DisplayNamePrefixed string `json:"display_name_prefixed"`
				URL                 string `json:"url"`
				PublicDescription   string `json:"public_description"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type SubredditHotPosts struct {
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				Title    string `json:"title"`
				Url      string `json:"url"`
				Selftext string `json:"selftext"`
				Id       string `json:"id"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type SubredditComment struct {
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Data struct {
				Body string `json:"body"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type HotPostWithComments struct {
	Title    string           `json:"title"`
	SelfText string           `json:"selftext"`
	Comments SubredditComment `json:"comments"`
}

// type Comment struct {
// 	SubredditID     string `json:"subreddit_id"`
// 	Author          string `json:"author"`
// 	Body            string `json:"body"`
// 	Replies         Replies `json:"replies"`
// 	// Add other fields you need from the comments
// }

// type Replies struct {
// 	Kind string `json:"kind"`
// 	Data struct {
// 		Children []struct {
// 			Kind string `json:"kind"`
// 			Data Comment `json:"data"`
// 		} `json:"children"`
// 	} `json:"data"`
// }

var client = &http.Client{}

func GetSubreddits(accessToken string) ([]byte, error) {
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
        return GetSubreddits(accessToken)
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

func GetHotPosts(accessToken string, subreddit string) ([]byte, error) {
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

	if resp.StatusCode == http.StatusTooManyRequests {
        fmt.Println("Rate limit hit for GetHotPosts. Waiting before retrying...")
        resetTime := resp.Header.Get("Retry-After")
        waitTime, _ := strconv.Atoi(resetTime)
        time.Sleep(time.Duration(waitTime) * time.Second)
        return GetHotPosts(accessToken, subreddit)
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

func GetComments(accessToken string, subreddit string, post_id string) ([]byte, error) {
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
        return GetComments(accessToken, subreddit, post_id)
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

func main() {
	accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzExOTA5MjMyLjg2NTM0MywiaWF0IjoxNzExODIyODMyLjg2NTM0MywianRpIjoiZWQtX0JabUN5TjFkZmpkY3hRZ2lvcDk3ajZvSjdBIiwiY2lkIjoicUk0bVgwQUctRGNaTFBEeWRZMHJtZyIsImxpZCI6InQyX2o0YnEwc3p4IiwiYWlkIjoidDJfajRicTBzengiLCJsY2EiOjE2NDMzOTc5MjcwMDAsInNjcCI6ImVKeUtWdEpTaWdVRUFBRF9fd056QVNjIiwiZmxvIjo5fQ.DYZw4Thr2aPVQ9I5IwMTDBgaar0_IfRrR6FXL2sL-ec_aKmdJs670NqUmyJOvsaDa7ubM22kR0TwW0sowkCNQtzZGbJxcn-7SI9XSawA0zA4DbQRXxJTCC_q5ZCVy-P-3xlhw2ElPWrj53hJDZ6E7QvwnmAcjfCBDWd_XbS7GXiQPK6j_ijBbeF2_EQK4LPlHK0QhE_SrjYHmQGYRmALrO6rVcPACIQ54LX9NKdCfx-9mUfnrB4nQKD3-DDSuPPg8ZD8MhUv_Cs8qSn_jBVkd89PanXbWzKv6kNEZFXEzkZ0iFiC7AClNUv1Cl0Z9wK1TaGEa8wLacBQJxJqtg1Vmg"

	r := gin.Default()

	r.GET("/subreddits", func(c *gin.Context) {
		body, err := GetSubreddits(accessToken)

		var subredditResponse SubredditResponse
		err = json.Unmarshal(body, &subredditResponse)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}
		c.JSON(http.StatusOK, subredditResponse)
	})

	r.GET("/hotpostsandcomments", func(c *gin.Context) {
		subreddit := c.Query("subreddit")

		body, err := GetHotPosts(accessToken, subreddit)

		var SubredditHotPosts SubredditHotPosts
		err = json.Unmarshal(body, &SubredditHotPosts)
		if err != nil {
			fmt.Println("Error unmarshalling response body:", err)
			// If user with specified ID is not found, return a 404 error
			c.JSON(http.StatusNotFound, gin.H{"error": "subreddit hot posts failed"})
			return
		}

		var hotPostsWithComments []HotPostWithComments

		var wg sync.WaitGroup
		for _, child := range SubredditHotPosts.Data.Children {
			wg.Add(1)
			child := child
			go func() {
				defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
				hotPost := HotPostWithComments{
					Title:    child.Data.Title,
					SelfText: child.Data.Selftext,
				}
				body, _ := GetComments(accessToken, subreddit, child.Data.Id)

				var subredditComments []SubredditComment
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

		c.JSON(http.StatusOK, hotPostsWithComments)
	})

	r.Run(":8080")
}
