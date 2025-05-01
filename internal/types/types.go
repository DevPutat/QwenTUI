package types

import "github.com/rivo/tview"

type Conf struct {
	ModelName string
	ApiURL    string
	ApiKey    string
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
