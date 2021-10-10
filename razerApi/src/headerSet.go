package src

import "github.com/labstack/echo"

func AddHeaderDEV(c echo.Context) {
	c.Response().Header().Set("Accept", "application/json")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}
