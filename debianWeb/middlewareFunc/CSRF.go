package middlewareFunc

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"thftgr.com/GoWorld/debianWeb/src"
)

func CSRF() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Request().Cookie("_csrf")
			if err == nil {
				c.Request().Header.Add("X-CSRF-Token", cookie.Value)
			}
			return next(c)
		}
	}
}
func CSRFConfig() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieDomain:   src.Config.Server.Domain,
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})
}
