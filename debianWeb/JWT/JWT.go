package JWT

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
	"thftgr.com/GoWorld/debianWeb/src"
)

func ParseToken(t string) (data map[string]interface{}, err error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(src.Config.Jwt.Key), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data = claims
	}
	return
}

func Jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return  func (c echo.Context) (err error){
		cookie, err := c.Cookie("token")
		if err != nil {
			return middlewareFunc.RequestIDToJson(http.StatusInternalServerError,c,"")
		}
		token, err := ParseToken(cookie.Value)
		if err != nil {
			return middlewareFunc.RequestIDToJson(http.StatusUnauthorized,c,"Invalid Token")
		}
		if err = mariadb.CheckToken(token["request_id"].(string));err != nil {
			middlewareFunc.RequestIDToJson(http.StatusUnauthorized,c,"Invalid Token")
			return
		}

		c.Set("JWT",token)
		return next(c)
	}
	
}

func AccessToken(){

}
func RefreshToken(){

}