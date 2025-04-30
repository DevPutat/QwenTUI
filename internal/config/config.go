package config

import (
	"encoding/json"
	"os"

	"github.com/DevPutat/QwenTUI/internal/types"
)

const (
	//defaultApiURL = "https://api.qwen.com/v1/chat"
	defaultApiURL = "https://openrouter.ai/api/v1/chat/completions"

	defaultApiKey    = "your-api-key"
	defaultModelName = "qwen/qwen3-235b-a22b:free"
	configPath       = "tui-ai-chat.conf"
)

var conf *types.Conf

func Config() *types.Conf {
	if conf == nil {
		conf = &types.Conf{}
		conf = loadConf()
	}
	return conf
}

func loadConf() *types.Conf {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		saveConf("", "", "")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, conf); err != nil {
		panic(err)
	}
	return conf
}

func saveConf(key string, url string, model string) {
	if len(url) == 0 {
		url = defaultApiURL
	}
	if len(model) == 0 {
		model = defaultModelName
	}
	if len(key) == 0 {
		key = defaultApiKey
	}
	c := &types.Conf{
		ApiKey:    key,
		ApiURL:    url,
		ModelName: model,
	}
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		panic(err)
	}
}
