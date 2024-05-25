package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordEmbed struct {
    Content string `json:"content,omitempty"`
    Embeds  []Embed `json:"embeds,omitempty"`
}

type Embed struct {
    Title       string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    URL         string `json:"url,omitempty"`
    Image       EmbedImage `json:"image,omitempty"`
}

type EmbedImage struct {
    URL string `json:"url,omitempty"`
}

func SendDiscordWebhook(webhookURL, message, imageURL, productURL string) error {
    embed := Embed{
        Title: "Product Back in Stock",
        Description: message,
        URL: productURL,
        Image: EmbedImage{URL: imageURL},
    }

    discordEmbed := DiscordEmbed{
        Embeds: []Embed{embed},
    }

    jsonBody, err := json.Marshal(discordEmbed)
    if err != nil {
        return err
    }

    resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}