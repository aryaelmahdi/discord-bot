package selenium

import (
	"discord-bot/configs"
	"fmt"
	"strconv"

	"github.com/tebeka/selenium"
)

func InitSelenium(config *configs.SeleniumConfig) (*selenium.Service, selenium.Capabilities, error) {
	cdPort, _ := strconv.Atoi(config.ChromeDriverPort)
	caps := selenium.Capabilities{"browserName": "chrome"}
	service, err := selenium.NewChromeDriverService(config.ChromeDriverPath, cdPort)
	if err != nil {
		return nil, nil, fmt.Errorf("Error starting ChromeDriver server: ", err)
	}
	return service, caps, nil
}
