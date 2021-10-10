package Route

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/JWT"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
	"thftgr.com/GoWorld/debianWeb/src"
	"time"
)

//유저아이디, 유저네임, 메인모드, 벤 상태, 유저 퍼미션 => privileges, 후원기간 => donor_expire
//JWT 토큰
//type token struct {
//	TokenType    string `json:"token_type"`
//	AccessToken  string `json:"access_token"`
//	RefreshToken string `json:"refresh_token"`
//}
//type data struct {
//	UserID     int
//	Username   string
//	MainMode   int
//	IsBan      bool
//	Permission int
//	Support    int
//}
//

var loginDelay = map[string]struct {
	Try     int
	LastTry time.Time
}{}

func Login(c echo.Context) (err error) {
	cookie, err := c.Cookie("token")
	if err == nil {
		_, err = JWT.ParseToken(cookie.Value)
		if err == nil {
			return c.String(http.StatusBadRequest,"You're already logged in!")
		}
	}


	username := c.Request().FormValue("username")
	password := c.Request().FormValue("password")

	dd, err := mariadb.LoginData(username, password) //유저 데이터
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = dd
	claims["request_id"] = c.Response().Header().Get(echo.HeaderXRequestID)
	claims["created"] = time.Now().Format("2006-01-02T15:04:05-07:00")
	claims["expiration"] = time.Now().Add(time.Hour * 24 * 30).Format("2006-01-02T15:04:05-07:00")

	t, err := token.SignedString([]byte(src.Config.Jwt.Key))
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}

	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Domain = src.Config.Server.Domain
	//cookie.HttpOnly = true // 사용시 js 에서 기본으로는 읽지 못함
	cookie.Expires = time.Now().Add(time.Hour * 24) // 이유는 모르겠지만 24시간 이하면 셋이 안댐
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
