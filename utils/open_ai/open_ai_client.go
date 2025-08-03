package open_ai

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// Question 表示单个选择题的结构
type Question struct {
	Question string           `json:"question"`
	Options  []QuestionOption `json:"options"`
	Answer   string           `json:"answer"`
}

// QuestionOption 表示选择题的选项
type QuestionOption struct {
	Label string `json:"label"`
	Text  string `json:"text"`
}

// OpenAIClient 封装与OpenAI的交互
type OpenAIClient struct {
	messages []openai.ChatCompletionMessage
	reader   *bufio.Reader
	client   *openai.Client
	prompts  *PromptWord
}

// NewOpenAIClient 创建新的OpenAI客户端实例
func NewOpenAIClient(apiKey, baseURL string) *OpenAIClient {
	config := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		config.BaseURL = baseURL
	}

	return &OpenAIClient{
		messages: make([]openai.ChatCompletionMessage, 0),
		reader:   bufio.NewReader(os.Stdin),
		client:   openai.NewClientWithConfig(config),
		prompts:  NewPromptWord(),
	}
}

// StartSession 开始一个新的会话
func (oc *OpenAIClient) StartSession(content string) ([]Question, error) {
	log.Println("Use AI question setting assistant")

	// 第一步：提取知识点
	knowledgePoints, err := oc.extractKnowledgePoints(content)
	if err != nil {
		return []Question{}, fmt.Errorf("failed to extract knowledge points: %v", err)
	}

	// 第二步：生成问题
	questions, err := oc.generateQuestions(knowledgePoints)
	if err != nil {
		return []Question{}, fmt.Errorf("generation problem failed: %v", err)
	}

	return questions, nil
}

// readUserInput 读取用户输入
func (oc *OpenAIClient) readUserInput() (string, error) {
	input, err := oc.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// extractKnowledgePoints 从文本中提取知识点
func (oc *OpenAIClient) extractKnowledgePoints(text string) ([]string, error) {
	oc.resetMessages()
	oc.addSystemMessage(oc.prompts.Question1)
	oc.addUserMessage(text)

	response, err := oc.getAIResponse()
	if err != nil {
		return nil, err
	}

	// 调试：打印原始响应
	log.Printf("Response: %s\n", response)

	// 尝试清理响应中的非JSON内容
	jsonStart := strings.Index(response, "[")
	jsonEnd := strings.LastIndex(response, "]")
	if jsonStart == -1 || jsonEnd == -1 {
		return nil, fmt.Errorf("no valid JSON response")
	}

	cleanResponse := response[jsonStart : jsonEnd+1]

	var knowledgePoints []string
	if err := json.Unmarshal([]byte(cleanResponse), &knowledgePoints); err != nil {
		return nil, fmt.Errorf("failed to parse knowledge points: %v\nresponse: %s", err, response)
	}

	return knowledgePoints, nil
}

// generateQuestions 根据知识点生成问题
func (oc *OpenAIClient) generateQuestions(knowledgePoints []string) ([]Question, error) {
	oc.resetMessages()
	oc.addSystemMessage(oc.prompts.Question2)

	knowledgeStr, err := json.Marshal(knowledgePoints)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize knowledge points: %v", err)
	}
	oc.addUserMessage(string(knowledgeStr))

	response, err := oc.getAIResponse()
	if err != nil {
		return nil, err
	}

	// 调试：打印原始响应
	log.Printf("Response: %s\n", response)

	// 尝试清理响应中的非JSON内容
	jsonStart := strings.Index(response, "[")
	jsonEnd := strings.LastIndex(response, "]")
	if jsonStart == -1 || jsonEnd == -1 {
		return nil, fmt.Errorf("no valid JSON response")
	}

	cleanResponse := response[jsonStart : jsonEnd+1]

	var questions []Question
	if err := json.Unmarshal([]byte(cleanResponse), &questions); err != nil {
		return nil, fmt.Errorf("failed to parse knowledge question : %v\nresponse: %s", err, response)
	}

	return questions, nil
}

// getAIResponse 获取AI响应
func (oc *OpenAIClient) getAIResponse() (string, error) {
	resp, err := oc.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    "qwen-plus",
			Messages: oc.messages,
		},
	)
	if err != nil {
		return "", fmt.Errorf("AI question failed: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no AI response received")
	}

	assistantReply := resp.Choices[0].Message.Content
	oc.addAssistantMessage(assistantReply)
	return assistantReply, nil
}

// 辅助方法
func (oc *OpenAIClient) resetMessages() {
	oc.messages = make([]openai.ChatCompletionMessage, 0)
}

func (oc *OpenAIClient) addSystemMessage(content string) {
	oc.messages = append(oc.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: content,
	})
}

func (oc *OpenAIClient) addUserMessage(content string) {
	oc.messages = append(oc.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}

func (oc *OpenAIClient) addAssistantMessage(content string) {
	oc.messages = append(oc.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
}
