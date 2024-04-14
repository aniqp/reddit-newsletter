package email_models

// EmailRecipient represents the data structure used to store the email recipient data.
type EmailRecipient struct {
	Username   string
	Email      string
	Subreddits []string
}

// EmailData represents the data structure used in the email template.
type EmailData struct {
	Date   string
	Reddit []RedditSummaryEntry
}

// RedditData represents the data for each Reddit entry in the email template.
type RedditSummaryEntry struct {
	Name    string
	Summary SummaryData
}

// SummaryData represents the data for each summary within a Reddit entry in the email template.
type SummaryData struct {
	Title string
	Body  string
}
