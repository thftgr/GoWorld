package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/privateComicServer/src"
)

func GetFileList(c *fiber.Ctx) (err error) {
	defer func() {
		err, _ = recover().(error)
		if err != nil {
			return
		}
	}()

	_ = c.JSON(src.RootDir.GetInfo(c.Query("path", "/")))

	return
}
