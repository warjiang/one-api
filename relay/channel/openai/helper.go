package openai

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/songquanpeng/one-api/relay/model"
)

func ResponseText2Usage(responseText string, modeName string, promptTokens int) *model.Usage {
	usage := &model.Usage{}
	usage.PromptTokens = promptTokens
	usage.CompletionTokens = CountTokenText(responseText, modeName)
	usage.TotalTokens = usage.PromptTokens + usage.CompletionTokens
	return usage
}

func ExtractAzureDeploymentConfig(cfg string, modelName string) (*DeploymentConfig, error) {
	var channelConfig AzureChannelConfig
	err := json.Unmarshal([]byte(cfg), &channelConfig)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal deployment config failed")
	}
	idx := -1
	for i, item := range channelConfig.DeploymentConfigs {
		if item.ModelName == modelName {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, errors.New(fmt.Sprintf("deployment config for %s not found", modelName))
	}
	return &channelConfig.DeploymentConfigs[idx], nil
}
