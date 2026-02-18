package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/db"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/genai"
)

func main() {
	//load the Gemini AI here 
	ctx := context.Background()
	_ = godotenv.Load(".env")
	clientgemini, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GENAI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

    result, err := clientgemini.Models.GenerateContent(
        ctx,
        "gemini-3-flash-preview",
        genai.Text("Explain how AI works in a few words"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())


	// Load environment variables from secret file (if present)
	_ = godotenv.Load(".env.mongo")

	url := "https://www.nasa.gov/wp-json/wp/v2/posts"

	// get back the response code
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// decode the data
	var posts []Post
	err = json.NewDecoder(res.Body).Decode(&posts)

	// Initialize MongoDB client
	client, err := db.ConnectToDB() // call the function from db.go
	collection := client.Database(db.DatabaseName).Collection(db.CollectionName)
	// Insert posts into the collection (assuming type conversion is compatible)
	docs := make([]interface{}, len(posts))
	for i, post := range posts{
		docs[i] = bson.M{
			"id" : post.ID,
			"link": post.Link,
			"title": post.Title.Rendered,
			"featured_image": string(post.Image),
		}
	}
	//insert the posts into the mango db database
	if len(docs) > 0 {
        if _, err := collection.InsertMany(context.Background(), docs); err != nil {
            fmt.Printf("failed to insert posts: %v\n", err)
            return
        }
        fmt.Println("posts inserted into MongoDB")
    } else {
        fmt.Println("no posts to insert")
    }
}	
			
		
