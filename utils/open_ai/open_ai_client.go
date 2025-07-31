package open_ai

import "github.com/sashabaranov/go-openai"

type OpenAIClient struct {
	message      []openai.ChatCompletionMessage
	questionList []string
	keywordsList []string
	ApiKey       string
	BaseURL      string
}

func NewOpenAIClient(apiKey, baseURL string) *OpenAIClient {
	return &OpenAIClient{
		ApiKey:  apiKey,
		BaseURL: baseURL,
	}
}
