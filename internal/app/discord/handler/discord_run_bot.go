package handler

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

func (handler *DiscordHandlerImpl) RunBot(username string, password string) error {
	driver, err := selenium.NewRemote(handler.Caps, fmt.Sprintf("http://%s:%s/wd/hub", handler.SeleniumConfig.ChromeDriverIP, handler.SeleniumConfig.ChromeDriverPort))
	if err != nil {
		return fmt.Errorf("error creating new WebDriver session: %v", err)
	}
	defer driver.Close()
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
		return url == handler.SeleniumConfig.DashboardURL, nil
	})
	if err != nil {
		return fmt.Errorf("error waiting for dashboard URL: %v", err)
	}

	dashboard, err := driver.FindElements(selenium.ByCSSSelector, "button.btn.btn-primary")
	if err != nil {
		return fmt.Errorf("error finding elements with text 'SPIRIT Academia': %v", err)
	}
	if len(dashboard) > 0 {
		fmt.Println("Found SPIRIT Academia on the page.")
	} else {
		fmt.Println("SPIRIT Academia not found on the page.")
		return fmt.Errorf("dashboard not found: %v", err)
	}

	time.Sleep(10000 * time.Millisecond)
	presentButton, err := driver.FindElements(selenium.ByXPATH, "//button[@type='button' and contains(@class, 'btn-primary') and text()=' Hadir ']")
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

		confirmationButtons, err := driver.FindElements(selenium.ByXPATH, "//button[contains(text(), 'Konfirmasi')]")
		if err != nil {
			return fmt.Errorf("error finding Confirmation Button %v", err)
		}
		fmt.Println("found confirmation button", confirmationButtons)

		for _, confirmationButton := range confirmationButtons {
			time.Sleep(5000 * time.Millisecond)
			if err := confirmationButton.Click(); err != nil {
				return fmt.Errorf("error clicking Confirmation Button %v", err)
			}
			fmt.Println("clicked")
		}
	}

	return nil
}
