package openai_api

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAiApiCommunicatorImpl struct {
	Client *openai.Client
}

func CreateOpenAiApiCommunicator(client *openai.Client) OpenAiApiCommunicatorImpl {
	return OpenAiApiCommunicatorImpl{Client: client}
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
