package openai_data

// GPT VISION
const GPT_VISION_MODEL string = "gpt-4-vision-preview"

type GptVisionRequest struct {
	Model     string             `json:"model"`
	Messages  []GptVisionMessage `json:"messages"`
	MaxTokens int                `json:"max_tokens"`
}

type GptVisionMessage struct {
	Role    string        `json:"role"`
	Content []interface{} `json:"content"`
}

type GptVisionTextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type GptVisionUrlContent struct {
	Type     string              `json:"type"`
	ImageUrl GptVisionContentUrl `json:"image_url"`
}

type GptVisionContentUrl struct {
	Url string `json:"url"`
}
