package openai_api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	es "github.com/mirogon/go_error_system"
	http_ "github.com/mirogon/go_http"
	openai_data "github.com/mirogon/go_openai_api/data"
	"github.com/sashabaranov/go-openai"
)

type OpenAiApiCommunicatorImpl struct {
	Client    *openai.Client
	OpenAiKey string
}

func CreateOpenAiApiCommunicator(client *openai.Client, openAiKey string) OpenAiApiCommunicatorImpl {
	return OpenAiApiCommunicatorImpl{Client: client, OpenAiKey: openAiKey}
}

func (c OpenAiApiCommunicatorImpl) GptCompletion(messages []openai_data.GptMessage, maxToken int, gptModel string) (string, es.Error) {
	request := openai_data.GptRequest{
		Model:     gptModel,
		Messages:  messages,
		MaxTokens: maxToken,
	}

	resp, err := sendRequest("POST", openai_data.GPT_REQUEST_URL, request, c.OpenAiKey)
	if err != nil {
		return "", es.NewError("giOIG9", err.Error(), nil)
	}

	if resp.StatusCode != 200 {
		return "", es.NewError("3geOkH", resp.Status, nil)
	}

	body, err := getResponseBody[openai_data.GptResponse](resp)
	if err != nil {
		return "", es.NewError("sudBQX", err.Error(), nil)
	}

	if len(body.Choices) < 1 {
		return "", es.NewError("i9iEmg", "no response", nil)
	}

	return body.Choices[0].Message.Content, nil
}

func (communicator OpenAiApiCommunicatorImpl) GenerateImage(input string) (string, es.Error) {
	imageReq := openai.ImageRequest{Prompt: input, Size: "1024x1024"}
	response, err := communicator.Client.CreateImage(context.TODO(), imageReq)
	if err != nil {
		return "", es.NewError("BDO041", err.Error(), nil)
	}
	return response.Data[0].URL, nil
}

type AudioToSpeechRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

func (communicator OpenAiApiCommunicatorImpl) TextToSpeech(input string, voice string) ([]byte, es.Error) {
	request := AudioToSpeechRequest{
		Model: "tts-1",
		Input: input,
		Voice: voice,
	}

	response, err := sendRequest("POST", "https://api.openai.com/v1/audio/speech", request, communicator.OpenAiKey)
	if err != nil {
		return nil, es.NewError("SPEs76", "sendRequest failed", err)
	}

	if response.StatusCode != 200 {
		return nil, es.NewError("A0rUGW", response.Status, nil)
	}

	bytes, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return bytes, es.NewError("juv9RE", err2.Error(), nil)
	}
	return bytes, nil
}

func sendRequest(requestMethod string, requestUrl string, requestBody interface{}, openAiKey string) (*http.Response, es.Error) {
	bodyRequestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, es.NewError("XBkfcD", err.Error(), nil)
	}

	bodyStringReader := strings.NewReader(string(bodyRequestJson))
	request, err := http.NewRequest(requestMethod, requestUrl, bodyStringReader)
	if err != nil {
		return nil, es.NewError("XBkfcD", err.Error(), nil)
	}

	request.Header.Add("Authorization", "Bearer "+openAiKey)
	request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("Accept", "application/json")

	requestSender := http_.HttpRequestSenderImpl{}
	response, err := requestSender.SendRequest(request)
	if err != nil {
		return nil, es.NewError("zGfdwK", "HttpRequestSender_SendRequest", nil)
	}

	return response, nil
}

func getResponseBody[responseType any](response *http.Response) (responseType, es.Error) {
	buffer, _ := io.ReadAll(response.Body)

	var responseBody responseType
	err := json.Unmarshal(buffer, &responseBody)
	if err != nil {
		return responseBody, es.NewError("W4rNoZ", err.Error(), nil)
	}
	return responseBody, nil
}
