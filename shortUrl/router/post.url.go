package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"github.com/thftgr/GoWorld/shortUrl/db"
	_ "github.com/thftgr/GoWorld/shortUrl/db"
	"github.com/thftgr/GoWorld/shortUrl/response"
	"math/rand"
	"time"
)

type body struct {
	Url string `json:"url"`
}

func PostUrl(c *fiber.Ctx) (err error) {
	var body body
	err = c.BodyParser(&body)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	var key string
	for {
		key = genKey()
		_, err = db.Sqlite3.Exec("INSERT INTO URLS(SHORT, URL) VALUES(?,?);", key, body.Url)
		if err == sqlite3.ErrConstraintPrimaryKey {
			continue
		} else if err != sqlite3.ErrConstraintPrimaryKey {
			break
		} else {
			return err
		}

	}

	return c.SendString(key)
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var charset = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

func genKey() (s string) {
	for i := 0; i < 10; i++ {
		s += string(charset[random.Intn(61)])
	}
	return
}
