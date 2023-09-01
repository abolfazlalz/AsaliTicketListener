package env

import (
	"github.com/caarlos0/env/v9"
)

type TelegramType struct {
	Token  string `env:"TELEGRAM_TOKEN"`
	UserId string `env:"TELEGRAM_USER_ID"`
}

var Telegram *TelegramType

func loadTelegram() error {
	Telegram = &TelegramType{}
	if err := env.Parse(Telegram); err != nil {
		return err
	}
	return nil
}
