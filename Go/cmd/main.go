package main

import (
	"context"
	"fmt"
	"log"
	"os"

	reddit_handler "reddit-newsletter/apis"

	firebase "firebase.google.com/go"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	var subreddits []string
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Use your Redis server address
		DB:   0,                // Use the default DB
	})
	sa := option.WithCredentialsFile("../../reddit-newsletter-firebase-key.json")
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

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access the environment variable
	accessToken := os.Getenv("REDDIT_ACCESS_TOKEN")
	rc := reddit_handler.NewRedditClient(accessToken)

	for _, subreddit := range subreddits {
		hotPosts, err := rc.GetHotPostsAndCommentsResponse(subreddit)
		if err != nil {
			fmt.Println("Error getting data:", err)
			return
		}

		reddit_handler.AddToRedisQueue(hotPosts, rdb)
	}
}
