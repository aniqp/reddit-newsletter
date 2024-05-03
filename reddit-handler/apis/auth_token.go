package reddit_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	reddit_models "reddit-newsletter/pkg/models"
	"strings"
)

func GetAccessToken(client *http.Client) (string, error) {
	accessTokenStruct := reddit_models.AccessToken{}

	redditUsername := os.Getenv("REDDIT_USERNAME")
	redditPassword := os.Getenv("REDDIT_PASSWORD")
	redditAuthUsername := os.Getenv("REDDIT_AUTH_USERNAME")
	redditAuthPassword := os.Getenv("REDDIT_AUTH_PASSWORD")
	log.Println("Reddit Username " + redditUsername)
	log.Println("Reddit Password " + redditPassword)
	log.Println("Reddit Auth Username " + redditAuthUsername)
	log.Println("Reddit Auth Password " + redditAuthPassword)

	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("username", redditUsername)
	form.Set("password", redditPassword)

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(redditAuthUsername, redditAuthPassword)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &accessTokenStruct)
	accessToken := accessTokenStruct.AccessToken

	return accessToken, err
}
