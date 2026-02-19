package agent

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func agent(text string){
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
	genai.Text(text),
	nil,
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result.Text())
}

