package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DevPutat/QwenTUI/internal/types"
)

func Send(query string, conf types.Conf) (string, error) {
	requestBody := types.ChatRequest{
		Model: conf.ModelName,
		Messages: []types.ChatMessage{
			{Role: "user", Content: query},
		},
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", conf.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp types.ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return "", err
	}
	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("пустой ответ от API")
}
