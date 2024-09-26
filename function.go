package function

import (
	"fmt"
	"net/http"
)

type Config struct {
	BotToken     string
	AccountIds   []string
	FunctionName string
}

type TelegramBot struct {
	Cfg Config
}

func (o *TelegramBot) SendTelegram(text string) error {
	client := &http.Client{}

	for _, e := range o.Cfg.AccountIds {
		botUrl := fmt.Sprintf("https://api.telegram.org/bot"+o.Cfg.BotToken+"/sendMessage?chat_id="+e+"&text=%s", text)
		request, err := http.NewRequest("GET", botUrl, nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(request)
		if err != nil {
			return err
		}
		resp.Body.Close()
	}

	return nil
}
