package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		After   string `json:"after"`
		Dist    int    `json:"dist"`
		Modhash interface{} `json:"modhash"`
		GeoFilter interface{} `json:"geo_filter"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc         interface{} `json:"approved_at_utc"`
				Subreddit             string      `json:"subreddit"`
				Selftext              string      `json:"selftext"`
				Title                 string      `json:"title"`
				Url                   string      `json:"url"`
				Id                    string      `json:"id"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func main() {
    // Set your OAuth2 access token
    accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlNIQTI1NjpzS3dsMnlsV0VtMjVmcXhwTU40cWY4MXE2OWFFdWFyMnpLMUdhVGxjdWNZIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ1c2VyIiwiZXhwIjoxNzEwNjE1MTk5LjM3Mjc5LCJpYXQiOjE3MTA1Mjg3OTkuMzcyNzksImp0aSI6IkR1LWR0MmRoME5FdmRQNzdiNjdBSnVwNU5RaThZUSIsImNpZCI6IkJTcnctR3NfWjFJVVZzaTlRZFlVSmciLCJsaWQiOiJ0Ml92bnhiOThkbnciLCJhaWQiOiJ0Ml92bnhiOThkbnciLCJsY2EiOjE3MDk3ODYxMjg4MzIsInNjcCI6ImVKeUtWdEpTaWdVRUFBRF9fd056QVNjIiwiZmxvIjo5fQ.ThSF5Q8YCBPiapOz8sMCxL6od_Vxrt16GAgn54iadAYCrNPiirZikj6ZZNxFfxa9DaltGrSjLj-GlT3VGG9KHvLHVF7ioHWWDFUW3v5Ods2hW3Pp-1JsrWxFlzRLhkGCqlU797lg063VeEf4myZw0twtB5i-z50HiqUXY9eNS8yXKnN4uwrmBErXcEkpmFMzmR3eoZmZttn4STwmzuFjLzCdTP-ij8MyanREjamj7lHO5UMCj5_6mIOgzwREpBg9EHV1LK7Jwr96BBR7S8BlhL4GsIKFlBSw9URDa0Tyg_n6Nit5MsnO2D-Cdyu-GACUQdOLncqkSWS8rvnvS2TU-Q"

    // Create a new HTTP client
    client := &http.Client{}

    // Create a new GET request
    req, err := http.NewRequest("GET", "https://oauth.reddit.com/subreddits/mine/subscriber", nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Add authorization header with OAuth2 access token
    req.Header.Set("Authorization", "Bearer " + accessToken)

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    // Check the status code
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error response:", resp.Status)
        fmt.Println("Response body:", string(body))
        return
    }

    // Unmarshal JSON response into struct
	var subredditResponse SubredditResponse
	err = json.Unmarshal(body, &subredditResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Subscribed Subreddits{")
	fmt.Println()
	// Print the filtered information
	for _, child := range subredditResponse.Data.Children {
		fmt.Println("Subreddit:", child.Data.DisplayNamePrefixed)
		fmt.Println("URL:", child.Data.URL)
		fmt.Println("Description:", child.Data.PublicDescription)
		fmt.Println()
	}
	fmt.Println("}")


	for _, child := range subredditResponse.Data.Children {
		// Create a new GET request
		req, err := http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/%s/hot?limit=10", child.Data.DisplayNamePrefixed), nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Add authorization header with OAuth2 access token
		req.Header.Set("Authorization", "Bearer " + accessToken)

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()
	
		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
	
		// Check the status code
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error response:", resp.Status)
			fmt.Println("Response body:", string(body))
			return
		}

		var responseBody SubredditHotPosts
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			fmt.Println("Error unmarshalling response body:", err)
			return
		}

		fmt.Println("EVERY SUBREDDITS HOT TOPIC{")
		fmt.Println()
		// Extract desired information and create an array of objects
		for _, child := range responseBody.Data.Children {
			fmt.Println("title:",  child.Data.Title)
			fmt.Println("url:", child.Data.Url)
			fmt.Println("selftext:", child.Data.Selftext)
			fmt.Println("postId:", child.Data.Id)
			fmt.Println()
		}
		fmt.Println("}")
	}

}
