package main

import (
	Logger "boardWeb/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Pre(middleware.Logger())
	Logger.Echo = e
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(":80"))
}
