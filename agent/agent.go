package agent

import (
	"context"
	"fmt"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type GroqAgent struct {
	client openai.Client  // value, not pointer
}

func NewGroqAgent(apiKey string) (*GroqAgent, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY is required")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL("https://api.groq.com/openai/v1"),
	)

	return &GroqAgent{client: client}, nil  // now works
}

func (g *GroqAgent) Enhance(text string) (string, error) {
	ctx := context.Background()

	resp, err := g.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: "llama-3.3-70b-versatile",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(text),
		},
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from Groq")
	}

	return resp.Choices[0].Message.Content, nil
}