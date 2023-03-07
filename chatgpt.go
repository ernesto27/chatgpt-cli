package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type chatGPT struct {
	gogptClient *openai.Client
}

func New() *chatGPT {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using system environment variables")
		fmt.Println()
	}

	return &chatGPT{
		gogptClient: openai.NewClient(os.Getenv("AUTH_TOKEN_OPEN_AI")),
	}
}

func (c *chatGPT) GetResponse(content string) (string, error) {
	resp, err := c.gogptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}
