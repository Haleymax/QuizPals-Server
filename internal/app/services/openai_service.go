package services

import (
	"QuizPals-Server/config"
	"QuizPals-Server/utils/open_ai"
)

type OpenAIService interface {
	GenerateQuestions(context string) ([]open_ai.Question, error)
}

type OpenAIServiceImpl struct {
	Client *open_ai.OpenAIClient
}

func NewOpenAIService() OpenAIService {
	cfg := config.GetConfig().OpenAI

	openAiClient := open_ai.NewOpenAIClient(cfg.APIKey, cfg.BaseURL)
	return OpenAIServiceImpl{
		Client: openAiClient,
	}
}

func (os OpenAIServiceImpl) GenerateQuestions(context string) ([]open_ai.Question, error) {
	quest, err := os.Client.StartSession(context)
	if err != nil {
		return []open_ai.Question{}, err
	}
	return quest, nil
}
