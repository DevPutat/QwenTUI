package config

import "github.com/DevPutat/QwenTUI/internal/types"

const (
	//defaultApiURL = "https://api.qwen.com/v1/chat"
	defaultApiURL = "https://openrouter.ai/api/v1/chat/completions"

	defaultApiKey    = "your-api-key"
	defaultModelName = "qwen/qwen3-235b-a22b:free"
)

func Config(key string) types.Conf {
	c := types.Conf{
		ApiKey:    defaultApiKey,
		ApiURL:    defaultApiURL,
		ModelName: defaultModelName,
	}
	if len(key) > 0 {
		c.ApiKey = key
	}
	return c
}
