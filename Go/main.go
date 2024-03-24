package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzExMzI5MDU3Ljg2MzYyNCwiaWF0IjoxNzExMjQyNjU3Ljg2MzYyNCwianRpIjoiaFVSRFdJLUlpR2ZLSExFSTM4aS1ubm9xazdlclRRIiwiY2lkIjoicUk0bVgwQUctRGNaTFBEeWRZMHJtZyIsImxpZCI6InQyX2o0YnEwc3p4IiwiYWlkIjoidDJfajRicTBzengiLCJsY2EiOjE2NDMzOTc5MjcwMDAsInNjcCI6ImVKeUtWdEpTaWdVRUFBRF9fd056QVNjIiwiZmxvIjo5fQ.Snixsf61qG82G0Ykv6V8S1jcu_c_dQYkYtHXfOW438KA0YFOKY5ez8UIrx_7mB7lpvjxBCkYFX53_iKRlo9M4WtRzWhX93d-zXosOJiknH_PSnd3Eq-7D2MfNU-iax_aARyVGF3HeJJPzR3QwvayL5j74pZZ3dj4eW-s5JNGyI90P4X4ZY9aU1yGPP7ytHZrz6bJXdwmmYJQmIkt9IyWFcbvA7tNC8KiD02ZJN65ck3o3QawGkr7_nwpLsmslvd9nM9OyNel25OrKs5J8d9p4_sKZdN2Gv3Y_nAIyGszv_1Wz7Y5KvoBFo2Zys1zW9A17wq0GvKVhphUo8MTecFG9g"

	body, err := GetSubreddits(accessToken)

	var subredditResponse SubredditResponse
	err = json.Unmarshal(body, &subredditResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	r := gin.Default()

	// Configure CORS middleware options
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Or the specific origins you want to allow
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3005" // Use this to allow specific origins
		},
		MaxAge: 12 * time.Hour,
	}

	// Apply the middleware to the router (all routes)
	r.Use(cors.New(config))

	r.GET("/subreddits", func(c *gin.Context) {
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

		for _, child := range SubredditHotPosts.Data.Children {
			fmt.Println("Selftext:", child.Data.Selftext)
			hotPost := HotPostWithComments{
				Title:    child.Data.Title,
				SelfText: child.Data.Selftext,
			}
			body, err := GetComments(accessToken, subreddit, child.Data.Id)

			var subredditComments []SubredditComment
			err = json.Unmarshal(body, &subredditComments)
			if err != nil {
				fmt.Println("Error unmarshalling comments:", err)
				// Handle the error as needed, maybe skip this post or return an error response
				continue
			}

			hotPost.Comments = subredditComments[1]
			hotPostsWithComments = append(hotPostsWithComments, hotPost)
		}

		c.JSON(http.StatusOK, hotPostsWithComments)
	})

	r.Run(":8080")
}
