package selenium

import (
	"discord-bot/configs"
	"fmt"
	"strconv"

	"github.com/tebeka/selenium"
)

func InitSelenium(config *configs.SeleniumConfig) (*selenium.Service, selenium.Capabilities, error) {
	seleniumPort, _ := strconv.Atoi(config.SeleniumPort)
	caps := selenium.Capabilities{"browserName": "chrome"}
	service, err := selenium.NewChromeDriverService(config.ChromeDriverPath, seleniumPort)
	if err != nil {
		return nil, nil, fmt.Errorf("Error starting ChromeDriver server: ", err)
	}
	return service, caps, nil
}
