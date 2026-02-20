package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"main/agent"
	"main/db"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Load environment variables
	_ = godotenv.Load(".env.mongo")
	_ = godotenv.Load(".env")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	posts, err := fetchPosts()
	if err != nil {
		log.Fatalf("failed to fetch posts: %v", err)
	}

	// groq 
	aiAgent, err := agent.NewGroqAgent(os.Getenv("GROQ_API_KEY"))
	if err != nil {
		log.Fatalf("failed to initialize AI agent: %v", err)
	}

	enhancedDocs := enhancePosts(posts, aiAgent)

	if err := insertIntoMongo(ctx, enhancedDocs); err != nil {
		log.Fatalf("failed to insert into MongoDB: %v", err)
	}

	fmt.Println("Posts successfully enhanced and stored.")
}

func fetchPosts() ([]Post, error) {
	url := "https://www.nasa.gov/wp-json/wp/v2/posts?per_page=30"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var posts []Post
	if err := json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func enhancePosts(posts []Post, ai *agent.GroqAgent) []interface{} {
	docs := make([]interface{}, 0, len(posts))

	for _, post := range posts {

		enhancedTitle, err := ai.Enhance(
			"Rewrite this title to be more engaging. Return only one title: " + post.Title.Rendered,
		)
		if err != nil {
			log.Println("title enhancement failed:", err)
			enhancedTitle = post.Title.Rendered
		} else {
			// take first line if multiple returned
			lines := strings.Split(enhancedTitle, "\n")
			enhancedTitle = strings.TrimSpace(lines[0])
		}

		// enhancedContent, err := ai.Enhance(
		// 	"Summarize and improve this content: " + post.Content.Rendered,
		// )
		// if err != nil {
		// 	log.Println("content enhancement failed:", err)
		// 	enhancedContent = post.Content.Rendered
		// }

		doc := bson.M{
			"id":             post.ID,
			"link":           post.Link,
			"title":          enhancedTitle,
			"content":        post.Content.Rendered,
			"featured_image": string(post.Image),
		}

		docs = append(docs, doc)
	}

	return docs
}

func insertIntoMongo(ctx context.Context, docs []interface{}) error {
	client, err := db.ConnectToDB()
	if err != nil {
		return err
	}

	collection := client.Database(db.DatabaseName).Collection(db.CollectionName)

	if len(docs) == 0 {
		return fmt.Errorf("no documents to insert")
	}

	_, err = collection.InsertMany(ctx, docs)
	return err
}