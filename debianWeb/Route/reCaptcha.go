package Route

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/src"
	"time"
)
type reCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
}
var errMsg = struct {
	message string
}{"ReCaptcha Fail"}

func ReCaptcha(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		if c.FormValue("g-recaptcha-response") == "" {
			return c.JSON(http.StatusInternalServerError,errMsg)
		}
		ServerKey := src.Config.ReCaptcha.ServerKey

		url := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s",ServerKey,c.FormValue("g-recaptcha-response"))
		method := "POST"

		client := &http.Client {}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError,errMsg)
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError,errMsg)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError,errMsg)
		}
		var cap reCaptchaResponse
		if json.Unmarshal(body, &cap) != nil {
			return c.JSON(http.StatusInternalServerError,errMsg)
		}
		if cap.Success {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized,errMsg)

	}





}
