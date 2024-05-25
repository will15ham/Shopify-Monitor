package tests

import (
	"fmt"
	"shopify-stock-monitor/pkg/config"
	"shopify-stock-monitor/pkg/shopify"
	"shopify-stock-monitor/pkg/util"
	"testing"
)

func TestCompareStockStatus(t *testing.T) {
	cfg := config.NewConfig()
	cfg.DiscordWebhookURL = "https://discord.com/api/webhooks/1243952825453641738/Vbr_x20dZeH9lJzv-i_pnR8bKK0GdcX53p9RYURl3ccEd6T77u27hWx_KpSZPX80QiKG"
	cfg.ShopifyBaseURL = "https://your-shopify-store.com"

	// varaint 1 and 2 are product 1
	// varaint 3 and 4 are product 2

    previousStatus := map[int64]bool{
        1: true,
        2: false,
        3: true,
    }

    currentStatus := map[int64]bool{
        1: true,
        2: true,
        3: false,
        4: true,
    }

	products := []shopify.Product{
        {
            ID:    1,
            Title: "Product 1",
            Handle: "product-1",
            Variants: []shopify.Variant{
                {ID: 1, Title: "Variant 1", Available: true},
                {ID: 2, Title: "Variant 2", Available: true},
            },
            Images: []shopify.Image{
                {ID: 1, Src: "https://example.com/image1.jpg"},
            },
        },
        {
            ID:    2,
            Title: "Product 2",
            Handle: "product-2",
            Variants: []shopify.Variant{
                {ID: 3, Title: "Variant 3", Available: false},
                {ID: 4, Title: "Variant 4", Available: true},
            },
            Images: []shopify.Image{
                {ID: 2, Src: "https://example.com/image2.jpg"},
            },
        },
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
						}
					}
				}
			}
		}
	}
}