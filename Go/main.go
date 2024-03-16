package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				Title         string      `json:"title"`
				Url           string      `json:"url"`
				Selftext      string      `json:"selftext"`
				Id            string      `json:"id"`
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
	Title string             	`json:"title"`
    SelfText string             `json:"selftext"`
    Comments SubredditComment   `json:"comments"`
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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/hot?limit=10", subreddit), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/comments/%s?sort=hot&limit=20&depth=3", subreddit, post_id), nil)
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
	accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzEwNjE1MTk5LjM3Mjc5LCJpYXQiOjE3MTA1Mjg3OTkuMzcyNzksImp0aSI6IkR1LWR0MmRoME5FdmRQNzdiNjdBSnVwNU5RaThZUSIsImNpZCI6IkJTcnctR3NfWjFJVVZzaTlRZFlVSmciLCJsaWQiOiJ0Ml92bnhiOThkbnciLCJhaWQiOiJ0Ml92bnhiOThkbnciLCJsY2EiOjE3MDk3ODYxMjg4MzIsInNjcCI6ImVKeUtWdEpTaWdVRUFBRF9fd056QVNjIiwiZmxvIjo5fQ.ThSF5Q8YCBPiapOz8sMCxL6od_Vxrt16GAgn54iadAYCrNPiirZikj6ZZNxFfxa9DaltGrSjLj-GlT3VGG9KHvLHVF7ioHWWDFUW3v5Ods2hW3Pp-1JsrWxFlzRLhkGCqlU797lg063VeEf4myZw0twtB5i-z50HiqUXY9eNS8yXKnN4uwrmBErXcEkpmFMzmR3eoZmZttn4STwmzuFjLzCdTP-ij8MyanREjamj7lHO5UMCj5_6mIOgzwREpBg9EHV1LK7Jwr96BBR7S8BlhL4GsIKFlBSw9URDa0Tyg_n6Nit5MsnO2D-Cdyu-GACUQdOLncqkSWS8rvnvS2TU-Q"
	
	body, err := GetSubreddits(accessToken)

	var subredditResponse SubredditResponse
	err = json.Unmarshal(body, &subredditResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	r := gin.Default()

	r.GET("/subreddits", func(c *gin.Context) {
		c.JSON(http.StatusOK, subredditResponse)
	})

	r.GET("/hotpostsandcomments", func(c *gin.Context) {
		subreddit := c.Query("subreddit")
		
		fmt.Println("subreddit:", subreddit)

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
				Title: child.Data.Title,
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
