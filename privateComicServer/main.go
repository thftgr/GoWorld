package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/thftgr/GoWorld/privateComicServer/route"
	"log"
)

func main() {
	f := fiber.New()
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	f.Use(logger.New())

	////파일들
	//f.Static("/signIn", "./signIn.html")
	//f.Static("/signUp/admin", "./signUp.html")
	//
	////service
	//f.Get("/", customMiddleware.TokenChecker, func(c *fiber.Ctx) error {
	//
	//	c.Response().SetBodyString(`{"token":"okokokokokoko"}`)
	//	return nil
	//})
	//
	//f.Get("/home", customMiddleware.TokenChecker, func(c *fiber.Ctx) error {
	//	c.Response().Header.Set("Content-Type", "text/html")
	//	return c.Response().SendFile("./index.html")
	//
	//})
	//
	////api
	//f.Post("/api/signIn", route.SignIn)
	//f.Post("/api/signUp/admin", route.AdminSignUp)

	f.Get("/api/file/list", route.GetFileList)
	f.Static("/api/file/", "./data")

	log.Fatalln(f.Listen(":80"))
}
