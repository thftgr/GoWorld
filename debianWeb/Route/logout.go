package Route

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/mariadb"
)

func Logout(c echo.Context) (err error){
	err = mariadb.Logout(c.Get("JWT").(map[string]interface{}))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Domain = "debian.thftgr.synology.me"
	cookie.HttpOnly = true // 사용시 js 에서 기본으로는 읽지 못함
	cookie.Path = "/"
	c.SetCookie(cookie)
	return nil

}
