package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func (handler *DiscordHandlerImpl) RunBot6A(username string, password string) error {
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		},
	}
	chromeDriverPort, _ := strconv.Atoi(handler.SeleniumConfig.ChromeDriverPort)
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)
	service, err := selenium.NewChromeDriverService(handler.SeleniumConfig.ChromeDriverPath, chromeDriverPort)
	if err != nil {
		return fmt.Errorf("Error starting ChromeDriver server: ", err)
	}
	defer service.Stop()

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://%s:%s/wd/hub", handler.SeleniumConfig.ChromeDriverIP, handler.SeleniumConfig.ChromeDriverPort))
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
	err = driver.WaitWithTimeout(func(d selenium.WebDriver) (bool, error) {
		url, err := d.CurrentURL()
		if err != nil {
			return false, err
		}
		return url == handler.SeleniumConfig.DashboardURL, nil
	}, 300*time.Second)
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
	time.Sleep(10 * time.Second)
	presentButton, err := driver.FindElements(selenium.ByXPATH, "//button[@type='button' and contains(@class, 'btn-primary') and text()=' Hadir ']")
	if err != nil {
		return fmt.Errorf("cannot find present button: %v", err)
	}

	fmt.Println("buttons found :", len(presentButton))
	if len(presentButton) == 0 {
		return fmt.Errorf("cannot find present button")
	}

	for _, button := range presentButton {
		if _, err := driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{button}); err != nil {
			fmt.Println("Error scrolling to element: ", err)
			return fmt.Errorf("error scrolling to element: %v", err)
		}
		if err := driver.WaitWithTimeout(selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
			isDisplayed, err := button.IsDisplayed()
			if err != nil {
				return false, err
			}
			return isDisplayed, nil
		}), 300*time.Second); err != nil {
			fmt.Println("Error waiting for element to be clickable: ", err)
		}
		if err := button.Click(); err != nil {
			return fmt.Errorf("error clicking 'Hadir' button: %v", err)
		}

		confirmationButton, err := driver.FindElement(selenium.ByXPATH, "//button[contains(text(), 'Konfirmasi')]")
		if err != nil {
			return fmt.Errorf("error finding Confirmation Button %v", err)
		}
		fmt.Println("found confirmation button", confirmationButton)
		if _, err := driver.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{confirmationButton}); err != nil {
			fmt.Println("Error scrolling to element: ", err)
			return fmt.Errorf("error scrolling to element: %v", err)
		}
		if err := driver.WaitWithTimeout(selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
			isDisplayed, err := confirmationButton.IsDisplayed()
			if err != nil {
				return false, err
			}
			return isDisplayed, nil
		}), 300*time.Second); err != nil {
			fmt.Println("Error waiting for element to be clickable: ", err)
		}
		time.Sleep(5 * time.Second)
		if err := confirmationButton.Click(); err != nil {
			return fmt.Errorf("error clicking Confirmation Button %v", err)
		}
		fmt.Println("clicked")
	}

	presentButton, err = driver.FindElements(selenium.ByXPATH, "//button[@type='button' and contains(@class, 'btn-primary') and text()=' Hadir ']")
	if err != nil {
		fmt.Println("no present button found")
	}
	if len(presentButton) != 0 {
		if err := handler.RunBot6A(username, password); err != nil {
			return fmt.Errorf("error ", err)
		}
	}
	fmt.Sprintf("username : %s success", username)

	return nil
}
