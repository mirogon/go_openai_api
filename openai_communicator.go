package openai_api

import (
	err "github.com/mirogon/go_error_system"
	es "github.com/mirogon/go_error_system"
	openai_data "github.com/mirogon/go_openai_api/data"
)

type GptCompletionRequestData struct {
	GptSessionId uint64
	TextInput    string
}

type GptCompletionResponse struct {
	ErrorCode    int    `json:"errorCode"`
	Result       string `json:"result"`
	GptSessionId uint64 `json:"gptSessionId"`
}

type OpenAiApiCommunicator interface {
	GptCompletion(messages []openai_data.GptMessage, maxToken int, gptModel string) (string, err.Error)
	GenerateImage(input string, numImages int, resolution string) ([]string, err.Error)
	TextToSpeech(input string, voice string) ([]byte, err.Error)
	GptVision(input string, imageUrl string) (string, es.Error)
}
