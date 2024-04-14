package openai_data

// GPT VISION
const GPT_VISION_MODEL string = "gpt-4-vision-preview"

type GptVisionRequest struct {
	Model     string       `json:"model"`
	Messages  []GptMessage `json:"messages"`
	MaxTokens int          `json:"max_tokens"`
}

type GptMessage struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

type GptMessageTextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type GptMessageUrlContent struct {
	Type     string        `json:"type"`
	ImageUrl GptContentUrl `json:"image_url"`
}

type GptContentUrl struct {
	Url string `json:"url"`
}
