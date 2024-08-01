package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"sync"
)

type Repo struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	Username      string `json:"username"`
	RepoName      string `json:"reponame"`
	Description   string `json:"description"`
	Lang          string `json:"lang"`
	LangColor     string `json:"langColor"`
	DetailPageURL string `json:"detailPageUrl"`
	StarCount     int    `json:"starCount"`
	ForkCount     int    `json:"forkCount"`
}

type Response struct {
	Code int    `json:"code"`
	Data []Repo `json:"data"`
}

var mutex sync.Mutex

func Top(args []string, lang string) {
	url := "https://e.juejin.cn/resources/github"
	p := fmt.Sprintf(`{"category":"trending","period":"month","lang":"%v","offset":0,"limit":20}`, lang)

	fmt.Println("Top start ...")
	requestBody := bytes.NewBuffer([]byte(p))

	// 发起 POST 请求
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Post request failed:", err)
		return
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Failed to decode JSON:", err)
		return
	}

	waitGroup := sync.WaitGroup{}
	for _, repo := range response.Data {
		waitGroup.Add(1)
		go func(repo Repo) {
			defer waitGroup.Done()
			TranslateOutput(repo)
		}(repo)
	}
	waitGroup.Wait()

}

func TranslateOutput(repo Repo) {
	msg := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("请根据我提供的内容，如果是英文，请翻译成中文，如果是中文则保持原样。：我提供的内容是：【%s】。输出翻译内容即可，不需要多余的内容。", repo.Description),
		},
	}
	out, err := Ai.OpenAiRequest(msg)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err.Error())
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("仓库:", repo.Username, "/", repo.RepoName, "\t Star", repo.StarCount, "\tFork", repo.ForkCount)
	fmt.Println("翻译：", out)
	fmt.Println("描述：", repo.Description)
	fmt.Println("===================================================")
}
