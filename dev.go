package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Debug(args []string) {
	if len(args) == 0 {
		fmt.Println("请输入你要翻译的错误内容")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("开发程序时报错了，请根据我的错误提示，翻译成中文，并且提供一下解决问题的方案，错误是：【%s】", args[0]),
		},
	}
	out, err := Ai.OpenAiRequest(msg)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err.Error())
		return
	}
	fmt.Println(out)
}

func Q(args []string) {

	if len(args) == 0 {
		fmt.Println("请输入你的问题")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: args[0],
		},
	}
	Ai.OpenAiRequestStream(msg)
}
func Sql(args []string) {

	if len(args) == 0 {
		fmt.Println("请输入你的描述，方便我为你生成sql语句")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("你是一个精通mysql sql语句的助手，可以根据我的描述帮我生成sql语句，并且提供详细的注释，我的描述是：【%s】", args[0]),
		},
	}
	Ai.OpenAiRequestStream(msg)
}
