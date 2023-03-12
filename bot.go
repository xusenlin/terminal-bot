package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"net/url"
	"terminalBot/config"
	"terminalBot/params"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := params.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	var dialog []openai.ChatCompletionMessage

	if p.ResetDialog {
		_ = c.RemoveBeforeChat()
	} else {
		dialog, _ = c.BeforeChat()
	}
	dialog = append(dialog, openai.ChatCompletionMessage{
		Role:    "user",
		Content: c.QuestionPrefix + p.Question + c.QuestionSuffix,
	})

	aiConfig := openai.DefaultConfig(c.AuthToken)

	if c.ProxyURL != "" {
		proxyUrl, err := url.Parse(c.ProxyURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		aiConfig.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	ai := openai.NewClientWithConfig(aiConfig)

	ctx := context.Background()
	message := append(c.Prompt, dialog...)
	stream, err := ai.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Stream:    true,
		Messages:  message,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	answer := ""
	defer stream.Close()
	defer func() {
		msg := append(dialog, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: answer,
		})
		if err := c.SaveChat(msg); err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n==============================================")
	}()
	fmt.Println("==============================================")
	for {
		response, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content := response.Choices[0].Delta.Content
		answer = answer + content
		fmt.Printf(content)
	}
}
