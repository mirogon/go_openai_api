package openai_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	http_ "github.com/mirogon/go_http"
	"github.com/sashabaranov/go-openai"
)

type OpenAiApiCommunicatorImpl struct {
	Client    *openai.Client
	OpenAiKey string
}

func CreateOpenAiApiCommunicator(client *openai.Client, openAiKey string) OpenAiApiCommunicatorImpl {
	return OpenAiApiCommunicatorImpl{Client: client, OpenAiKey: openAiKey}
}

func (c OpenAiApiCommunicatorImpl) GptCompletion(message string, maxToken int) (string, error) {
	messages := []openai.ChatCompletionMessage{{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}}

	var request openai.ChatCompletionRequest = openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Messages:  messages,
		MaxTokens: maxToken,
	}

	resp, err := c.Client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (communicator OpenAiApiCommunicatorImpl) GenerateImage(input string) (string, error) {
	imageReq := openai.ImageRequest{Prompt: input, Size: "1024x1024"}
	response, err := communicator.Client.CreateImage(context.TODO(), imageReq)
	if err != nil {
		return "", err
	}
	return response.Data[0].URL, nil
}

type AudioToSpeechRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

func (communicator OpenAiApiCommunicatorImpl) TextToSpeech(input string, voice string) ([]byte, error) {
	request := AudioToSpeechRequest{
		Model: "tts-1",
		Input: input,
		Voice: voice,
	}

	response, _ := sendRequest("POST", "https://api.openai.com/v1/audio/speech", request, communicator.OpenAiKey)

	buffer := make([]byte, 1000000)
	size, _ := response.Body.Read(buffer)
	buffer = buffer[:size]
	fmt.Println("Size: ", size)
	fmt.Println(string(buffer))

	return buffer, nil
}

func sendRequest(requestMethod string, requestUrl string, requestBody interface{}, openAiKey string) (*http.Response, error) {
	bodyRequestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	bodyStringReader := strings.NewReader(string(bodyRequestJson))
	request, err := http.NewRequest(requestMethod, requestUrl, bodyStringReader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+openAiKey)
	request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("Accept", "application/json")

	requestSender := http_.HttpRequestSenderImpl{}
	response, err := requestSender.SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getResponseBody[responseType any](response *http.Response) (responseType, error) {
	buffer := make([]byte, 100000)
	size, _ := response.Body.Read(buffer)
	buffer = buffer[:size]

	var responseBody responseType
	err := json.Unmarshal(buffer, &responseBody)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
