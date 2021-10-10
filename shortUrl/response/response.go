package response

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Response(c *fiber.Ctx,code int,body interface{})(err error){

	b, err := json.Marshal(body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		_, _ = c.WriteString(err.Error())
		return
	}
	c.Status(code)
	_, _ = c.WriteString(string(b))
	return

}
