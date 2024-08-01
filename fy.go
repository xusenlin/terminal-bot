package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func FY(args []string) {

	if len(args) == 0 {
		fmt.Println("请输入你要翻译的中文内容")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("I want you to act as an English translator, spelling corrector and improver. I will speak to you in any language and you will detect the language, translate it and answer in the corrected and improved version of my text, in English. I want you to replace my simplified A0-level words and sentences with more beautiful and elegant, upper level English words and sentences. Keep the meaning same, but make them more literary. I want you to only reply the correction, the improvements and nothing else, do not write explanations. My first sentence is【%s】", args[0]),
		},
	}
	out, err := Ai.OpenAiRequest(msg)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err.Error())
		return
	}
	fmt.Println(out)

}

func FYY(args []string) {

	if len(args) == 0 {
		fmt.Println("请输入你要翻译的中文内容")
		return
	}
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("我想让你充当中文翻译助手。我会给你发送英文文本，请用中文翻译它，不要写解释。我的第一句英文是【%s】", args[0]),
		},
	}
	out, err := Ai.OpenAiRequest(msg)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err.Error())
		return
	}
	fmt.Println(out)

}
