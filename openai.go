package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
	"strings"
)

type AiClient struct {
	client *openai.Client
}

func NewClient() (*AiClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY is not set")
	}
	apiUrl := os.Getenv("OPENAI_API_URL")
	if apiUrl == "" || !strings.HasPrefix(apiUrl, "http") {
		return nil, errors.New("OPENAI_API_URL is not set")
	}
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = apiUrl + "/v1"

	ai := &AiClient{
		client: openai.NewClientWithConfig(config),
	}
	return ai, nil
}

func (c *AiClient) OpenAiRequest(msg []openai.ChatCompletionMessage) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 2048,
		Messages:  msg,
		Stream:    false,
	}
	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func (c *AiClient) OpenAiRequestStream(msg []openai.ChatCompletionMessage) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 2048,
		Messages:  msg,
		Stream:    true,
	}
	stream, err := c.client.CreateChatCompletionStream(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			break
		}
		fmt.Print(response.Choices[0].Delta.Content)
	}
}
