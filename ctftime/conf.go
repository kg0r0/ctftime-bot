package ctftime

import (
	"encoding/json"
	"io/ioutil"
)

// Config is ...
type Config struct {
	SlackConfig SlackConfig `json:"slack_config"`
}

// SlackConfig is ...
type SlackConfig struct {
	APIToken  string `json:"api_token"`
	ChannelID string `json:"channel_id"`
}

// NewConfig is ...
func NewConfig(configPath string) (*Config, error) {
	config := new(Config)
	err := config.readConfig(configPath)
	return config, err
}

func (config *Config) readConfig(configPath string) error {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	return nil
}
