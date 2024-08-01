package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func MM(args []string) {
	if len(args) == 0 {
		fmt.Println("请输入一个变量名描述")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("请根据大多数知名源代码帮我取一个变量名，用它来表示【%s】", args[0]),
		},
	}
	out, err := Ai.OpenAiRequest(msg)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err.Error())
		return
	}
	fmt.Println(out)
}
