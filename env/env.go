package env

import (
	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	environments := []func() error{
		loadTelegram,
		loadService,
	}

	for _, environment := range environments {
		if err := environment(); err != nil {
			return err
		}
	}

	return nil
}
