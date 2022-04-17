package main

import (
	"fmt"
	"time"

	SE "github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

)

const (
	KChromeDriverLocation = "C:\\chromedriver.exe"
	//Selenium selectors
	ByID              = "id"
	ByXPATH           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name"
	ByTagName         = "tag name"
	ByClassName       = "class name"
	ByCSSSelector     = "css selector"
	//wait interval
	DefaultWaitInterval = 2000 * time.Millisecond
	// DefaultWaitTimeout is the default timeout for selenium.Wait function.
	DefaultWaitTimeout = 5 * time.Second

	//Xpath and css links
	LoginButtonXpath = "//button[@type='submit']"
	LeftPaneElements = "//mat-sidenav[@position='start']/div[1]/cog-app-nav[1]/mat-nav-list[1]/div[6]"
	AdvDiagXpath     = "//mat-sidenav[@position='start']/div[1]/cog-app-nav[1]/mat-nav-list[1]/div[10]"
)


type config struct {
	addr, browser, path string
	chromeoptions       []string
}

// func (sc *seleniumCohesity) SearchElement(By,value string) {

// }
//Inti selenium wendriver and return the object
func InitWebdriver() (sco SE.WebDriver, errn error) {
	var WebD SE.WebDriver
	seleniumCapabilities := SE.Capabilities{"browserName": "google-chrome",
		"idleTimeout": 1000000}
	port := 443
	chromeConfig := config{
		addr: fmt.Sprintf("http://%s:%d/wd/hub",
			"localhost", port),
		chromeoptions: []string{"-incognito", "--ignore-certificate-errors"},
	}
	// Including chrome options to open browser in incognito mode
	chromeCapabilities := chrome.Capabilities{
		Args: chromeConfig.chromeoptions,
	}
	fmt.Println("Creating new remote using Chrome")
	_, err1 := SE.NewChromeDriverService(KChromeDriverLocation, port)
	if err1 != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err1)
		return nil, err1
	}
	seleniumCapabilities.AddChrome(chromeCapabilities)

	wd, errWd := SE.NewRemote(seleniumCapabilities, chromeConfig.addr)
	if errWd != nil {
		fmt.Errorf("Error while creating new selenium remote connection %s",
			errWd)
		return nil, errWd
	}
	wd.Refresh()
	EndPoint := fmt.Sprintf("https://%s/login", "10.14.19.43")
	wd.Get(EndPoint)
	wd.MaximizeWindow("")

	searchElement, err := wd.FindElement("name", "username")
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}
	err = searchElement.SendKeys("admin")
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}
	searchElement, err = wd.FindElement("name", "password")
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}
	err = searchElement.SendKeys("admin")
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}

	searchElement, err = wd.FindElement("xpath", LoginButtonXpath)
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}
	err = searchElement.Click()
	if err != nil {
		fmt.Errorf("Error while starting New Chrome Driver Service %s", err)
		return nil, err
	}
	WebD = wd
	return WebD, nil
}


//main Entry point for the function
func main() {
	InitWebdriver()
}
