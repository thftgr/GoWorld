package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/privateComicServer/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AdminSignUpStruct struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Password string `json:"password"`
}

func AdminSignUp(c *fiber.Ctx) (err error) {
	row := db.Sqlite3.QueryRow("select count(*) > 0 as hasAdmin from USER where GROUP_ID = 1")

	var hasAdmin bool
	row.Scan(&hasAdmin)
	if hasAdmin {
		c.Response().SetStatusCode(http.StatusLocked)
		return
	}

	var req AdminSignUpStruct
	if c.BodyParser(&req) != nil {
		c.Response().SetStatusCode(http.StatusBadRequest)
		c.Response().SetBodyString("request parse fail.")
		log.Println(req)
		return
	}
	data, err := bcrypt.GenerateFromPassword([]byte(req.Password),10)
	_, err = db.Sqlite3.Exec(`
			INSERT INTO USER(LOGIN_ID,PASSWORD,EMAIL,GROUP_ID) VALUES (?,?,?,1)
			
		`, req.Id, string(data) , req.Email)
	if err != nil {
		c.Response().SetStatusCode(http.StatusInternalServerError)
		c.Response().SetBodyString(err.Error())
		return err
	}

	return c.JSON(req)
}
