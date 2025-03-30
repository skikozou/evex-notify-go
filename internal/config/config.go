package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../config/.env")
	if err != nil {
		return nil, err
	}

	return &Config{
		GITHUB_TOKEN:         os.Getenv("GITHUB_TOKEN"),
		TWITTER_BEARER_TOKEN: os.Getenv("TWITTER_BEARER_TOKEN"),
		TWITTER_AUTH_TOKEN:   os.Getenv("TWITTER_AUTH_TOKEN"),
		DISCORD_TOKEN:        os.Getenv("DISCORD_TOKEN"),
	}, nil
}
