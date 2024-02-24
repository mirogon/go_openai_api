package openai_api

import es "github.com/mirogon/go_error_system"

type DebugOpenAiCommunicator struct{}

func (communicator DebugOpenAiCommunicator) GptCompletion(message string) (string, es.Error) {
	return "DebugAnswer", nil
}

func (communicator DebugOpenAiCommunicator) GenerateImage(input string) (string, es.Error) {
	return "", nil
}
