package main

import (
	"fmt"
	"shopify-stock-monitor/pkg/config"
	"shopify-stock-monitor/pkg/shopify"
	"shopify-stock-monitor/pkg/util"
	"time"
)

func main() {
	cfg := config.NewConfig()

	if cfg.DiscordWebhookURL == "" {
        fmt.Println("Error: DISCORD_WEBHOOK_URL environment variable is not set.")
        return
    }

	if cfg.ShopifyBaseURL == "" {
		fmt.Println("Error: SHOPIFY_BASE_URL environment variable is not set.")
		return
	}
	
	previousStatus := make(map[int64]bool)

	for {
		fetchProductData(previousStatus, cfg)
		time.Sleep(15 * time.Minute)
	}
}

func fetchProductData(previousStatus map[int64]bool, cfg *config.Config) {
	products, err := shopify.FetchProductData(cfg.ShopifyBaseURL)
	if err != nil {
		util.Log(fmt.Sprintf("Error fetching product data: %v", err))
		return
	}

	currentStatus := make(map[int64]bool)
	for _, product := range products {
		for _, variant := range product.Variants {
			currentStatus[variant.ID] = variant.Available
		}
	}

	if len(previousStatus) == 0 {
		for _, product := range products {
			for _, variant := range product.Variants {
				message := fmt.Sprintf("Product: %s\nVariant: %d\nIn Stock: %t\n", product.Title, variant.ID, variant.Available)
				if !variant.Available {
					continue
				}
				imageURL := product.Images[0].Src
				productURL := fmt.Sprintf("%s/products/%s", cfg.ShopifyBaseURL, product.Handle)
				util.SendDiscordWebhook(cfg.DiscordWebhookURL, message, imageURL, productURL)
			}
		}
	} else {
		shopify.CompareStockStatus(previousStatus, currentStatus)

		for id, status := range currentStatus {
			if !previousStatus[id] && status {
				var product shopify.Product
				var variant shopify.Variant
				for _, p := range products {
					for _, v := range p.Variants {
						if v.ID == id {
							product = p
							variant = v
							imageURL := product.Images[0].Src
							productURL := fmt.Sprintf("%s/products/%s", cfg.ShopifyBaseURL, product.Handle)
							message := fmt.Sprintf("Product: %s\nVariant: %d\nIn Stock: %t\n", product.Title, variant.ID, variant.Available)
							util.SendDiscordWebhook(cfg.DiscordWebhookURL, message, imageURL, productURL)
							time.Sleep(1 * time.Second) 							// sleep for 1 second to avoid rate limiting
						}
					}
				}
			}
		}
	}

	for k, v := range currentStatus {
		previousStatus[k] = v
	}
}
