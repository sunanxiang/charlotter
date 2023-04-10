package server

import (
	"context"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/sunanxiang/charlotter/cache"
	"github.com/sunanxiang/charlotter/config"
)

func Completions(question string, msgChan chan string, errChan chan error) {
	client := openai.NewClient(config.OpenAiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v", err)
		errChan <- err
		return
	}

	err = cache.GlobalCache.Replace(question, resp.Choices[0].Message.Content, time.Minute*5)
	if err != nil {
		log.Printf("cache add error: %v", err)
	}

	msgChan <- resp.Choices[0].Message.Content
}
