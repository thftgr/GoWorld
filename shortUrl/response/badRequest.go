package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type badRequestBody struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

func BadRequest(c *fiber.Ctx,body string)(err error){
	b := badRequestBody{
		Timestamp:time.Now().Format(time.RFC3339),
		Status:http.StatusBadRequest,
		Error:http.StatusText(http.StatusBadRequest),
		Message:body,
		Path:c.Path(),
	}
	c.Status(b.Status)
	return c.JSON(b)

}
