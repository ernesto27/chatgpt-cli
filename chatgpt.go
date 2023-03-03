package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type chatGPT struct {
	gogptClient *gogpt.Client
}

func New() *chatGPT {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using system environment variables")
		fmt.Println()
	}

	return &chatGPT{
		gogptClient: gogpt.NewClient(os.Getenv("AUTH_TOKEN_OPEN_AI")),
	}
}

func (c *chatGPT) GetResponse(content string) (string, error) {
	resp, err := c.gogptClient.CreateChatCompletion(
		context.Background(),
		gogpt.ChatCompletionRequest{
			Model: gogpt.GPT3Dot5Turbo,
			Messages: []gogpt.ChatCompletionMessage{
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
