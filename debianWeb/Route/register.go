package Route

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strings"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
	"thftgr.com/GoWorld/debianWeb/src"
)

func Register(c echo.Context) (err error) {
	defer func() {
		if err != nil{
			middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
			return
		}
	}()
	username := strings.TrimSpace(c.Request().FormValue("username"))
	password := strings.TrimSpace(c.Request().FormValue("password"))
	email := strings.TrimSpace(c.Request().FormValue("email"))
	var reg *regexp.Regexp
	//유저네임
	if reg, err = regexp.Compile("([a-zA-Z0-9-_\\s]){2,30}"); err != nil {
		return
	}
	if username != reg.FindString(username) {
	}

	alive, err := mariadb.AliveUsername(username)
	if err != nil {
		return
	}
	for _,v := range src.Config.InvalidUsername {
		if strings.ToLower(v) == strings.ToLower(username){
			alive = false
			break
		}
	}
	if !alive {
		return middlewareFunc.RequestIDToJson(http.StatusBadRequest, c, "An user with that username already exists!")
	}

	//이메일
	if reg, err = regexp.Compile("(.+)+@\\w([-_.]?\\w)*[.][a-zA-Z]{2,3}"); err != nil {
		return
	}
	if email != reg.FindString(email) {
		return middlewareFunc.RequestIDToJson(http.StatusBadRequest, c, "Invalid email")
	}
	alive, err = mariadb.AliveEmail(email)
	if err != nil {

		return
	}
	if !alive {
		return middlewareFunc.RequestIDToJson(http.StatusBadRequest, c, "An user with that email address already exists!")
	}

	//비밀번호

	if len(password) < 8 {
		return middlewareFunc.RequestIDToJson(http.StatusBadRequest, c, "Invalid password")
	}


	fmt.Println("PASS")

	return c.NoContent(http.StatusOK)
}

