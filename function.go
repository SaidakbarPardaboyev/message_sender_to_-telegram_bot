package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Config struct {
	BotToken   string
	AccountIds []string
}

type TelegramBot struct {
	Cfg Config
}

func (t *TelegramBot) SendTelegram(text string) error {
	client := &http.Client{}

	for _, e := range t.Cfg.AccountIds {
		// URL encode the text
		encodedText := url.QueryEscape(text)

		// Include parse_mode=Markdown for formatting
		botUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s&parse_mode=Markdown", t.Cfg.BotToken, e, encodedText)
		request, err := http.NewRequest("GET", botUrl, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close() // Use defer here to ensure the body is closed

		// Check the response status
		if resp.StatusCode != http.StatusOK {
			var errorResponse struct {
				Description string `json:"description"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
				return fmt.Errorf("failed to decode error response: %w", err)
			}
			return fmt.Errorf("failed to send message: %s", errorResponse.Description)
		}
	}

	return nil
}
