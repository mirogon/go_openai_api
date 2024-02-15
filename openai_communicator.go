package openai_api

import openai_data "github.com/mirogon/go_openai_api/data"

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
	GptCompletion(messages []openai_data.GptMessage, maxToken int, gptModel string) (string, error)
	GenerateImage(input string) (string, error)
	TextToSpeech(input string, voice string) ([]byte, error)
}
