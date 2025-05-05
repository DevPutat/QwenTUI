package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DevPutat/QwenTUI/internal/types"
)

var timeOut = 60 * time.Second

func Send(query string, conf *types.Conf) (string, error) {
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

	client := &http.Client{
		Timeout: timeOut,
	}
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

func SendStream(query string, conf *types.Conf, app *types.App) {

	requestBody := types.ChatRequest{
		Model: conf.ModelName,
		Messages: []types.ChatMessage{
			{Role: "user", Content: query},
		},
		Stream: true,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", conf.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.ApiKey)

	client := &http.Client{
		Timeout: timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	reader := resp.Body
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			chunk := string(buf[:n])
			// app.UpdateOutput(chunk)
			processChunk(app, chunk)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
	return
}

func processChunk(app *types.App, chunk string) {
	var response types.StreamResponse
	err := json.Unmarshal([]byte(chunk), &response)
	if err != nil {
		return
	}
	if len(response.Choices) > 0 && response.Choices[0].Delta.Content != "" {
		app.UpdateOutput(response.Choices[0].Delta.Content)
	}
}
