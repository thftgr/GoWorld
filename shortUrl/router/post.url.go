package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/shortUrl/response"
)

type body struct {
	Url string `json:"url"`
}

func PostUrl(c *fiber.Ctx) (err error) {
	var body body
	err = c.BodyParser(&body)
	if err != nil {
		return response.BadRequest(c,err.Error())
	}


	return c.SendString("Hello, World 👋!")
}

