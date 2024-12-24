package function

import (
	"fmt"
	"net/http"
	"time"
)

type (
	Config struct {
		BotToken     string
		AccountIds   []string
		FunctionName string
	}

	TelegramBot struct {
		Cfg Config
	}
)

var (
	InfoSigniture  = " --- 🔵INFO --- "
	ErrorSigniture = " -- 🔴ERROR -- "
)

func (o *TelegramBot) SendTelegram(text string, status ...string) error {
	client := &http.Client{}

	var (
		message = o.Cfg.FunctionName
	)

	if len(status) > 0 && status[0] == "error" {
		message += " " + ErrorSigniture
	} else {
		message += " " + InfoSigniture
	}
	text = message + time.Now().Format("2006-01-02 15:04:05") + " " + text

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

// var (
// 	botToken = "7508206728:AAEDli0feSjSVqhbLs3CR4GCZssbbPy9LBY"
// 	chatId   = []string{"2066773119"}
// )

// func main() {

// 	var bot = TelegramBot{Cfg: Config{
// 		BotToken:     botToken,
// 		AccountIds:   chatId,
// 		FunctionName: "brrauf-function",
// 	}}

// 	bot.SendTelegram("hello", "error")
// }
