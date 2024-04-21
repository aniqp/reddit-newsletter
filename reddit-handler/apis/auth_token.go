package reddit_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	reddit_models "reddit-newsletter/pkg/models"
	"strings"
)

func GetAccessToken(client *http.Client) (error, string) {
	accessTokenStruct := reddit_models.AccessToken{}

	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("username", "redditsummarydummy")
	form.Set("password", "Sdg.sjM49GCP#mP")

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err), ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("BSrw-Gs_Z1IUVsi9QdYUJg", "tDGnKTEmcfBweqaFlUspGM4DEJu8WQ")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err), ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err), ""
	}

	err = json.Unmarshal(body, &accessTokenStruct)
	accessToken := accessTokenStruct.AccessToken

	return err, accessToken
}
