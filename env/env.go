package env

import (
	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := loadTelegram(); err != nil {
		return err
	}

	return nil
}
