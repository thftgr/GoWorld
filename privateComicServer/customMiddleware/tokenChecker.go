package customMiddleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/privateComicServer/db"
)



func TokenChecker(c *fiber.Ctx) error {
	if db.Token[c.Cookies("token")].Id == 0 {
		return c.Redirect("/signIn")
	}
	return c.Next()
}
