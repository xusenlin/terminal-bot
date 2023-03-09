package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"terminalBot/utils"
)

var ConfigFilePath = ""
var BeforeChatFilePath = ""

type Config struct {
	AuthToken      string                         `json:"authToken"`
	ProxyURL       string                         `json:"proxyURL"`
	Prompt         []openai.ChatCompletionMessage `json:"prompt"`
	QuestionPrefix string                         `json:"questionPrefix"`
}

func init() {
	currentUser, err := user.Current()
	configDir := currentUser.HomeDir + "/terminalBot"
	if err != nil {
		panic(err)
	}
	appName := filepath.Base(os.Args[0])
	BeforeChatFilePath = fmt.Sprintf("%v/%vChat.json", configDir, appName)
	ConfigFilePath = fmt.Sprintf("%v/%v.json", configDir, appName)
}

func Parse() (*Config, error) {
	var params Config

	if !utils.IsFile(ConfigFilePath) {
		return &params, errors.New(fmt.Sprintf("there is no corresponding '%v' configuration file in the same directory.", ConfigFilePath))
	}

	configData, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		return &params, err
	}

	err = json.Unmarshal(configData, &params)

	if err != nil {
		return &params, err
	}

	return &params, nil
}

func (c *Config) BeforeChat() ([]openai.ChatCompletionMessage, error) {
	var chat []openai.ChatCompletionMessage

	if !utils.IsFile(BeforeChatFilePath) {
		return chat, errors.New("previous chat file does not exist")
	}
	configData, err := ioutil.ReadFile(BeforeChatFilePath)
	if err != nil {
		return chat, err
	}

	err = json.Unmarshal(configData, &chat)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func (c *Config) RemoveBeforeChat() error {
	return utils.RemoveFile(BeforeChatFilePath)
}
func (c *Config) SaveChat(dialog []openai.ChatCompletionMessage) error {

	if err := utils.RemoveFile(BeforeChatFilePath); err != nil {
		return err
	}

	jsonData, err := json.Marshal(dialog)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(BeforeChatFilePath, jsonData, 0666)
	if err != nil {
		return err
	}
	return nil
}
