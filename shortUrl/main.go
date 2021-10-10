package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/shortUrl/router"
)

func main() {

	f := fiber.New()

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	f.Post("/url",router.PostUrl)
	f.Get("/:key",router.GetUrl)
	f.Listen(":80")
}

