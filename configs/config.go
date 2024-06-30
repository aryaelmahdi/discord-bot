package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type SeleniumConfig struct {
	SeleniumPort     string
	TargetURL        string
	ChromeDriverPath string
	ChromeDriverPort string
	ChromeDriverIP   string
}

type BotToken struct {
	Token  string
	Prefix string
}

type AppConfig struct {
	MySQL    MySQLConfig
	BotToken BotToken
	Selenium SeleniumConfig
}

func InitConfig() (*AppConfig, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	}
	return &AppConfig{
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
		BotToken: BotToken{
			Token:  os.Getenv("BOT_TOKEN"),
			Prefix: os.Getenv("BOT_PREFIX"),
		},
		Selenium: SeleniumConfig{
			SeleniumPort:     os.Getenv("SELENIUM_PORT"),
			TargetURL:        os.Getenv("TARGET_URL"),
			ChromeDriverPath: os.Getenv("CHROME_DRIVER_PATH"),
			ChromeDriverPort: os.Getenv("CHROME_DRIVER_PORT"),
			ChromeDriverIP:   os.Getenv("CHROME_DRIVER_IP"),
		},
	}, nil
}
