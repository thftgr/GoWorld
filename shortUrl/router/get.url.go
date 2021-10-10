package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/shortUrl/db"
	"github.com/thftgr/GoWorld/shortUrl/response"
	"net/http"
)

func GetUrl(c *fiber.Ctx) (err error) {
	for {
		row := db.Sqlite3.QueryRow("SELECT URL FROM URLS WHERE SHORT = ?;", c.Params("key"))
		if err != nil {
			return response.BadRequest(c, err.Error())
		}
		var s string
		err = row.Scan(&s)
		if err != nil {
			return response.BadRequest(c, err.Error())
		}
		return c.Redirect(s, http.StatusTemporaryRedirect)

	}

}
