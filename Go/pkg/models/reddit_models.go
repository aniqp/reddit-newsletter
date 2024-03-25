package reddit_models

// SubredditResponse defines the structure of the subreddit response
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

// SubredditHotPosts defines the structure of the subreddit hot posts response
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

// SubredditComment defines the structure of the subreddit comment response
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

// HotPostWithComments defines the structure of a hot post with comments
type HotPostWithComments struct {
	Title    string           `json:"title"`
	SelfText string           `json:"selftext"`
	Comments SubredditComment `json:"comments"`
}
