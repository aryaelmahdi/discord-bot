package selenium

import (
	"discord-bot/configs"
	"fmt"
	"strconv"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func InitSelenium(config *configs.SeleniumConfig) (*selenium.Service, selenium.Capabilities, error) {
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless", // Run Chrome in headless mode
			"--disable-gpu",
			"--no-sandbox", // Required when running as root
		},
	}
	seleniumPort, _ := strconv.Atoi(config.SeleniumPort)
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)
	service, err := selenium.NewChromeDriverService(config.ChromeDriverPath, seleniumPort)
	if err != nil {
		return nil, nil, fmt.Errorf("Error starting ChromeDriver server: ", err)
	}
	return service, caps, nil
}
