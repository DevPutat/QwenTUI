package types

import "github.com/rivo/tview"

type Conf struct {
	ModelName string
	ApiURL    string
	ApiKey    string
}

type StreamResponse struct {
	ID       string         `json:"id"`
	Object   string         `json:"object"`
	Created  int64          `json:"created"`
	Model    string         `json:"model"`
	Choices  []StreamChoice `json:"choices"`
	Logprobs any            `json:"logprobs,omitempty"`
}

type StreamChoice struct {
	Index              int         `json:"index"`
	Delta              ChatMessage `json:"delta"`
	FinishReason       string      `json:"finish_reason,omitempty"`
	NativeFinishReason string      `json:"native_finish_reason,omitempty"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

type App struct {
	App       *tview.Application
	ChatField *tview.TextView
}

func (a *App) UpdateOutput(text string) {
	a.App.QueueUpdateDraw(func() {
		curText := a.ChatField.GetText(true)
		a.ChatField.SetText(curText + text)
	})
}
