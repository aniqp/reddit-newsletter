package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	reddit_handler "reddit-newsletter/apis"

	firebase "firebase.google.com/go"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	c := cron.New()
	// Scheduled for 11 pm everyday.
	_, err := c.AddFunc("0 23 * * *", func() {
		runGetHotPostsAndComments(ctx)
	})
	if err != nil {
		log.Fatal("Error scheduling the cron job:", err)
	}

	c.Start()

	select {}
}

func runGetHotPostsAndComments(ctx context.Context) {

	accessClient := &http.Client{}
	accessToken, _ := reddit_handler.GetAccessToken(accessClient)

	var subreddits []string
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Use your Redis server address
		DB:   0,            // Use the default DB
	})
	sa := option.WithCredentialsFile("reddit-newsletter-firebase-key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("subreddits").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		name, ok := doc.Data()["name"].(string)
		if !ok {
			log.Fatalf("Failed to convert 'name' to string")
		}
		subreddits = append(subreddits, name)
	}

	fmt.Println("Subreddits: %", subreddits)
	defer client.Close()

	rc := reddit_handler.NewRedditClient(accessToken)

	for _, subreddit := range subreddits {
		hotPosts, err := rc.GetHotPostsAndCommentsResponse(subreddit)
		if err != nil {
			fmt.Println("Error getting data:", err)
			return
		}

		reddit_handler.AddToRedisQueue(hotPosts, rdb)
	}
	reddit_handler.AddToRedisQueue("done fetching", rdb)
}
