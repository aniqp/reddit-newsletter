package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

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

func GetComments(accessToken string, subreddit string, post_id string) ([]byte, error) {
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

func main() {
	accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzExMzMxODIxLjkwODkyLCJpYXQiOjE3MTEyNDU0MjEuOTA4OTIsImp0aSI6Ilk1Z1RmYXBRUHlrdUFOZkllMkowcWJBZG1GaW1XdyIsImNpZCI6IkJTcnctR3NfWjFJVVZzaTlRZFlVSmciLCJsaWQiOiJ0Ml92bnhiOThkbnciLCJhaWQiOiJ0Ml92bnhiOThkbnciLCJsY2EiOjE3MDk3ODYxMjg4MzIsInNjcCI6ImVKeUtWdEpTaWdVRUFBRF9fd056QVNjIiwiZmxvIjo5fQ.KVbUxz2bg1Mi0LIZn7B8r36pPBgptuzlGBERKHrXTOG3YmGu-2-x14RjGJ0N3jZSjsp6KKjlzW7k9z76ov6M494kCBrOrtVK2nNLkqZ1rZ2cUnAoj7389YAOu5rQzi-7-FCrotvZEHWzPDKcaPBAy1xNFONVDd9__VTrXmgFrDXBRPvSqsG0o3QQL22DuYZ17AuFgIKItVbb1S3fypwp05Bq5zMTGYSN-s4ORlXoETeYgoSCQlaxpaDZ3kKOaw3qvRM3vdpc3aMX9akDAnVz3W9RgAHA1Zz9oHk3-g2Umo-aFuhrZcjlnsq-Jm6dmwgWxflugzU-e9Ak1hSNs207lQ"

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
			}()
		}
		wg.Wait()

		c.JSON(http.StatusOK, hotPostsWithComments)
	})

	r.Run(":8080")
}
