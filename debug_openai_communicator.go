package openai_api

type DebugOpenAiCommunicator struct{}

func (communicator DebugOpenAiCommunicator) GptCompletion(message string) (string, error) {
	return "DebugAnswer", nil
}

func (communicator DebugOpenAiCommunicator) GenerateImage(input string) (string, error) {
	return "", nil
}
