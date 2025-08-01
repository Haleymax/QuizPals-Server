package open_ai

// PromptWord contains the prompts used for generating questions and keywords
type PromptWord struct {
	Question1 string `json:"question1"`
	Question2 string `json:"question2"`
}

func NewPromptWord() *PromptWord {
	Question1 := `你是一个出题专家，请严格按以下要求执行：
			1. 分析我提供的文本内容，提取关键知识点
			2. 只返回一个纯JSON格式的数组，不要包含任何额外的文字、解释或标记
			3. 示例格式：["知识点1", "知识点2", "知识点3"]
			4. 每个知识点应该简洁明了，不超过15个字`

	Question2 := `你是一个出题专家，请严格按以下要求执行：
			1. 根据提供的知识点生成选择题
			2. 只返回一个纯JSON格式的数组，不要包含任何额外的文字、解释或标记
			3. 每个问题格式如下：
			{
				"question": "题目内容",
    			"options": [
        			{"label": "1", "text": "选项内容"},
        			{"label": "2", "text": "选项内容"},
        			{"label": "3", "text": "选项内容"},
        			{"label": "4", "text": "选项内容"}
    			],
    			"answer": "正确答案标号"
			}
			4. 确保题目与知识点相关，选项有迷惑性`

	return &PromptWord{
		Question1: Question1,
		Question2: Question2,
	}
}
