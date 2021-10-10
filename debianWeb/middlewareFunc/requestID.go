package middlewareFunc

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RequestIDToJson(code int, c echo.Context,message string) (err error) {
	if code == http.StatusInternalServerError {
		message = ""
	}
	rid := struct {
		RequestId string `json:"request_id"`
		Message string `json:"message"`
	}{c.Response().Header().Get(echo.HeaderXRequestID),message}

	return c.JSON(code, rid)
}
