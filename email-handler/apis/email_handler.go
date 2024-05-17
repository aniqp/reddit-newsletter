package email_handler

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	email_models "email-handler/pkg/models"
)

type EmailRecipient email_models.EmailRecipient
type RecipientsList []EmailRecipient

// type EmailData email_models.EmailData
// type SummaryData email_models.SummaryData
// type RedditSummaryEntry email_models.RedditSummaryEntry

// Function to retrieves the list of recipients and their data from the Firestore database.
func FetchRecipientsList(client *firestore.Client, ctx context.Context) (RecipientsList, error) {
	var recipients RecipientsList

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		// Use the document ID as the username
		username := doc.Ref.ID

		// Parse the Firestore document into an EmailRecipient struct
		var data EmailRecipient
		if err := doc.DataTo(&data); err != nil {
			return nil, err
		}

		// Set the username
		data.Username = username

		// Query the subreddits subcollection of the current user document
		subredditIter := doc.Ref.Collection("subreddits").Documents(ctx)
		var subredditIDs []string
		for {
			subredditDoc, err := subredditIter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}

			// Check if the 'starred' field is true before appending the subreddit ID
			var subredditData map[string]interface{}
			if err := subredditDoc.DataTo(&subredditData); err != nil {
				return nil, err
			}
			if starred, ok := subredditData["starred"].(bool); ok && starred {
				subredditIDs = append(subredditIDs, subredditDoc.Ref.ID)
			}
		}

		// Set the list of subreddit IDs in the EmailRecipient struct
		data.Subreddits = subredditIDs

		// Append the EmailRecipient to the recipients list
		recipients = append(recipients, data)
	}
	return recipients, nil
}

func FetchSubredditSummaries(client *firestore.Client, ctx context.Context, subredditIDs []string) (email_models.EmailData, error) {
	var result email_models.EmailData

	currentDate := time.Now().Format("2006-01-02")

	result.Date = currentDate

	result.Reddit = make([]email_models.RedditSummaryEntry, 0)
	// formattedData := time.Now().Format("2024/01/02")

	for _, subredditID := range subredditIDs {
		docs := client.Collection("daily_subreddit_summary").Doc(subredditID).Collection("dates").Doc(currentDate)

		doc, err := docs.Get(ctx)
		if err != nil {
			if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
				continue // No summary found for this subreddit and date, move to the next one
			}
			return result, fmt.Errorf("error retrieving summary for subreddit %s: %v", subredditID, err)
		}

		var summary email_models.SummaryData
		err = doc.DataTo(&summary)
		if err != nil {
			return result, fmt.Errorf("error parsing summary for subreddit %s: %v", subredditID, err)
		}

		result.Reddit = append(result.Reddit, email_models.RedditSummaryEntry{
			Name:    subredditID, // You may replace this with the actual name of the subreddit
			Summary: summary,
		})
	}
	return result, nil
}

func generateEmailTemplate(data email_models.EmailData) (string, error) {
	templateStr := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Your Daily Reddit Newsletter {{.Date}}</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					background-color: #f9f9f9;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					padding: 20px;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 4px 6px rgba(0,0,0,0.1);
				}
				.header {
					text-align: center;
					margin-bottom: 20px;
				}
				.header h1 {
					color: #333;
				}
				.reddit-item {
					margin-bottom: 30px;
					padding: 20px;
					background-color: #f3f3f3;
					border-radius: 6px;
				}
				.reddit-item h2 {
					color: #333;
					margin-top: 0;
				}
				.reddit-item p {
					color: #666;
					margin-bottom: 0;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Your Daily Reddit Newsletter {{.Date}}</h1>
				</div>
				{{range .Reddit}}
				<div class="reddit-item">
					<h2>{{.Name}}</h2>
					<h3>{{.Summary.Title}}</h3>
					<p>{{.Summary.Body}}</p>
				</div>
				{{end}}
			</div>
		</body>
	</html>`

	// Parse the template
	tmpl, err := template.New("emailTemplate").Parse(templateStr)
	if err != nil {
		return "", err
	}

	// Execute the template
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func SendEmail(data email_models.EmailData, recipientEmail string) error {
	// Generate the email HTML content
	htmlContent, err := generateEmailTemplate(data)
	if err != nil {
		return err
	}

	// Get email sender and password from environment variables
	from := os.Getenv("EMAIL_SENDER")
	pass := os.Getenv("EMAIL_APP_PASSWORD")
	to := recipientEmail

	// Compose the email message with MIME headers
	body := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Your Daily Reddit Newsletter\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" + htmlContent

	// Send the email using SMTP
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(body))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	fmt.Println("Email sent successfully to", recipientEmail)
	return nil
}
