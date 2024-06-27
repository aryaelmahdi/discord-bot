package handler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

func (handler *DiscordHandlerImpl) RunBot(username, password string) error {
	cdPort, _ := strconv.Atoi(handler.SeleniumConfig.ChromeDriverPort)
	caps := selenium.Capabilities{"browserName": "chrome"}
	service, err := selenium.NewChromeDriverService(handler.SeleniumConfig.ChromeDriverPath, cdPort)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://%s:%s/wd/hub", handler.SeleniumConfig.ChromeDriverIP, handler.SeleniumConfig.ChromeDriverPort))
	if err != nil {
		return fmt.Errorf("error creating new WebDriver session: %v", err)
	}
	defer driver.Quit()

	if err := driver.Get(handler.SeleniumConfig.TargetURL); err != nil {
		return fmt.Errorf("error opening URL: %v", err)
	}
	time.Sleep(3000 * time.Millisecond)

	usernameField, err := driver.FindElement(selenium.ByCSSSelector, "input[placeholder='Username']")
	if err != nil {
		return fmt.Errorf("error finding username field: %v", err)
	}
	passwordField, err := driver.FindElement(selenium.ByCSSSelector, "input[placeholder='Password']")
	if err != nil {
		return fmt.Errorf("error finding password field: %v", err)
	}

	if err := usernameField.SendKeys(username); err != nil {
		return fmt.Errorf("error entering username: %v", err)
	}
	if err := passwordField.SendKeys(password); err != nil {
		return fmt.Errorf("error entering password: %v", err)
	}

	loginButton, err := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	if err != nil {
		return fmt.Errorf("error finding login button: %v", err)
	}

	if err := loginButton.Click(); err != nil {
		return fmt.Errorf("error clicking login button: %v", err)
	}
	err = driver.Wait(func(d selenium.WebDriver) (bool, error) {
		url, err := d.CurrentURL()
		if err != nil {
			return false, err
		}
		return url == "https://academia.pnj.ac.id/dashboard", nil // Adjust the URL as needed
	})
	if err != nil {
		return fmt.Errorf("error waiting for dashboard URL: %v", err)
	}

	dashboard, err := driver.FindElements(selenium.ByXPATH, "//*[contains(text(), 'SPIRIT Academia')]")
	if err != nil {
		return fmt.Errorf("error finding elements with text 'SPIRIT Academia': %v", err)
	}
	if len(dashboard) > 0 {
		fmt.Println("Found SPIRIT Academia on the page.")
	} else {
		fmt.Println("SPIRIT Academia not found on the page.")
		return fmt.Errorf("dashboard not found: %v", err)
	}

	time.Sleep(3000 * time.Millisecond)
	presentButton, err := driver.FindElements(selenium.ByXPATH, "//button[contains(text(), ' Hadir ')]")
	if err != nil {
		return fmt.Errorf("cannot find present button: %v", err)
	}

	fmt.Println("buttons found :", len(presentButton))
	if len(presentButton) == 0 {
		return fmt.Errorf("cannot find present button")
	}

	for _, button := range presentButton {
		if err := button.Click(); err != nil {
			return fmt.Errorf("error clicking 'Hadir' button: %v", err)
		}

		confirmationButton, err := driver.FindElement(selenium.ByXPATH, "//button[contains(text(), 'Konfirmasi')]")
		if err != nil {
			return fmt.Errorf("error finding Confirmation Button %v", err)
		}
		fmt.Println("found confirmation button", confirmationButton)

		if err := confirmationButton.Click(); err != nil {
			return fmt.Errorf("error clicking Confirmation Button %v", err)
		}
	}

	return nil
}
