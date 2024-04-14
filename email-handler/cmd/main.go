package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	email_handler "email-handler/apis"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("../reddit-newsletter-firebase-key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	recipients, err := email_handler.FetchRecipientsList(client, ctx)
	if err != nil {
		log.Fatalf("Error fetching recipients: %v", err)
	}

	// Loop through the recipients and send their email
	for _, recipient := range recipients {
		// Now you can access each EmailRecipient in the recipients list
		fmt.Println("Username:", recipient.Username)
		fmt.Println("Email:", recipient.Email)
		fmt.Println("Subreddits:", recipient.Subreddits)
		fmt.Println("-----------------------------")

		// Fetch subreddit summaries for the current recipient
		emailContent, err := email_handler.FetchSubredditSummaries(client, ctx, recipient.Subreddits)
		if err != nil {
			fmt.Printf("Error fetching subreddit summaries for recipient %s: %v\n", recipient.Username, err)
			continue
		}

		// If there are no subreddit summaries, continue to the next recipient
		if len(emailContent.Reddit) == 0 {
			continue
		}

		// Print the fetched email data
		fmt.Println("Date:", emailContent.Date)
		for _, redditEntry := range emailContent.Reddit {
			fmt.Println("Reddit Name:", redditEntry.Name)
			fmt.Println("Summary Title:", redditEntry.Summary.Title)
			fmt.Println("Summary Body:", redditEntry.Summary.Body)
			fmt.Println("------------------------------")
		}

		// Send email
		err = email_handler.SendEmail(emailContent, recipient.Email)
		if err != nil {
			fmt.Printf("Error sending email to recipient %s: %v\n", recipient.Username, err)
			continue
		}
	}
}
