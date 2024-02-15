package openai_data

const GPT_REQUEST_URL string = "https://api.openai.com/v1/chat/completions"

const GPT_3_MODEL string = "gpt-3.5-turbo"
const GPT_4_MODEL string = "gpt-4-turbo-preview"

//GPT CHAT

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptRequest struct {
	Model            string       `json:"model"`
	Messages         []GptMessage `json:"messages"`
	Temperature      int          `json:"temperature"`
	MaxTokens        int          `json:"max_tokens"`
	TopP             int          `json:"top_p"`
	FrequencyPenalty int          `json:"frequency_penalty"`
	PresencePenalty  int          `json:"presence_penalty"`
}

type GptUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GptCompletionChoice struct {
	Index   int        `json:"index"`
	Message GptMessage `json:"message"`
	// FinishReason
	// stop: API returned complete message,
	// or a message terminated by one of the stop sequences provided via the stop parameter
	// length: Incomplete model output due to max_tokens parameter or token limit
	// function_call: The model decided to call a function
	// content_filter: Omitted content due to a flag from our content filters
	// null: API response still in progress or incomplete
	FinishReason string `json:"finish_reason"`
}

type GptResponse struct {
	ID      string                `json:"id"`
	Object  string                `json:"object"`
	Created int64                 `json:"created"`
	Model   string                `json:"model"`
	Choices []GptCompletionChoice `json:"choices"`
	Usage   GptUsage              `json:"usage"`
}

// GPT VISION
const GPT_VISION_MODEL string = "gpt-4-vision-preview"

type GptVisionMessageContent struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	ImageUrl string `json:"image_url"`
}

type GptVisionMessage struct {
	Role    string                    `json:"role"`
	Content []GptVisionMessageContent `json:"content"`
}

type GptVisionRequest struct {
	Model     string             `json:"model"`
	Messages  []GptVisionMessage `json:"messages"`
	MaxTokens int                `json:"max_tokens"`
}
