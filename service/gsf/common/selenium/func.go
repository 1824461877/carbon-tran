package selenium

import (
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"time"
)

func TrimHiddenCharacter(originStr string) string {
	srcRunes := []rune(originStr)
	dstRunes := make([]rune, 0, len(srcRunes))
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}
	return string(dstRunes)
}

func AuthData(wd selenium.WebDriver) (*Token, error) {
	time.Sleep(5 * time.Second)
	script, err := wd.ExecuteScript("return window.localStorage.getItem('persist:auth')", nil)
	if err != nil {
		return AuthData(wd)
	}

	if script == nil {
		return AuthData(wd)
	}

	var info = TrimHiddenCharacter(script.(string))
	var gt GsfToken
	err = json.Unmarshal([]byte(info), &gt)
	if err != nil {
		return nil, err
	}

	if gt.Data.(string) != "null" {
		var token Token
		err = json.Unmarshal([]byte(gt.Data.(string)), &token)
		if err != nil {
			return nil, err
		}
		return &token, nil
	} else {
		return AuthData(wd)
	}
}

func GetToken(c *Cache) {
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		fmt.Println(err.Error())
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://gsf-sustaincert.auth0.com/login?state=hKFo2SBCeFJUZ0dWRXF6a2dHT2NJTnNHaGZrVTlZUlJlNE1nRKFupWxvZ2luo3RpZNkgcWNUR1oybXNZemRlem9LUC1fOV96M0M4aGVwRGQ1THCjY2lk2SBJakNYOTgycUsxbHhEMHFJU3RYNFRtSk1ud3JlMjY2dQ&client=IjCX982qK1lxD0qIStX4TmJMnwre266u&protocol=oauth2&response_type=token%20id_token&redirect_uri=https%3A%2F%2Fregistry.goldstandard.org%2Flogin%2Fcallback&scope=openid%20email%20profile&_send_telemetry=true&_times_to_retry_failed_requests=0&token_issuer=https%3A%2F%2Fgsf-sustaincert.auth0.com%2F&root_url=https%3A%2F%2Fgsf-sustaincert.auth0.com&nonce=LbMR80cQxH6S9v36HG8KGkjzKeRHfnde&auth0Client=eyJuYW1lIjoiYXV0aDAuanMiLCJ2ZXJzaW9uIjoiOS44LjIifQ%3D%3D"); err != nil {
		panic(err)
		return
	}

	time.Sleep(10 * time.Second)
	// Get a reference to the text box containing code.
	if email, errs := wd.FindElement(selenium.ByName, "email"); err != nil {
		panic(errs)
		return
	} else {
		errs = email.SendKeys("zhangzhiheng@vanke.com")
		if errs != nil {
			return
		}
	}

	if password, errs := wd.FindElement(selenium.ByName, "password"); err != nil {
		panic(errs)
		return
	} else {
		errs = password.SendKeys("Zzz1824461877.")
		if errs != nil {
			return
		}
	}

	if element, errs := wd.FindElements(selenium.ByTagName, "button"); errs != nil {
		panic(errs)
		return
	} else {
		errs = element[len(element)-1].Click()
		if errs != nil {
			return
		}
	}

	if token, er := AuthData(wd); er == nil || token != nil {
		c.Token = token
		wd.Quit()
	}

}

func Api() {

}
