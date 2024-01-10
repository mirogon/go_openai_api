package openai_api

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
	GptCompletion(message string, maxToken int) (string, error)
	GenerateImage(input string) (string, error)
	TextToSpeech(input string, voice string) ([]byte, error)
}
