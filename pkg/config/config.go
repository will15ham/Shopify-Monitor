package config

import "os"

type Config struct {
    DiscordWebhookURL string
	ShopifyBaseURL     string
}

func NewConfig() *Config {
    return &Config{
        DiscordWebhookURL: os.Getenv("DISCORD_WEBHOOK_URL"),
		ShopifyBaseURL: os.Getenv("SHOPIFY_BASE_URL"),
	}
}