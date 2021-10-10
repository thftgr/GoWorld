package Route

import (
	"github.com/labstack/echo/v4"
)


func TestTokenCheck(c echo.Context) (err error) {
	return c.String(200, "OK")
}

