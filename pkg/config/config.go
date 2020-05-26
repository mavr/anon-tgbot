package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Config containing configuration values
type Config struct {
	// Debug set logs level output
	App struct {
		Debug bool `toml:"debug"`
		Port  int  `toml:"port"`

		ExchangeNewChat string `toml:"new_chat_notification_exchange"`
	}

	// Telegram bot token
	Bot struct {
		TGBotToken string `toml:"telegram_bot_token"`
	}

	Rabbit struct {
		ConnectionString string `toml:"connection_string"`
	}
}

// New create new dealult config
func New(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := toml.Unmarshal(buf, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
