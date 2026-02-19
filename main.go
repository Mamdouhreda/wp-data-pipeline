package main

import (
	"context"
	"encoding/json"
	"fmt"

	"main/db"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	
)

func main() {
	

	// Load environment variables from secret file (if present)
	_ = godotenv.Load(".env.mongo")

	url := "https://www.nasa.gov/wp-json/wp/v2/posts?per_page=30"

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
			
		
