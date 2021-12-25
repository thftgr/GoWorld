package route

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thftgr/GoWorld/privateComicServer/db"
	"github.com/thftgr/GoWorld/privateComicServer/entity"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)
type SignInStruct struct {
	Id string `json:"id" xml:"name" form:"name"`
	Password string `json:"password" xml:"pass" form:"pass"`
}

func SignIn (c *fiber.Ctx) (err error) {
	var req SignInStruct
	c.BodyParser(&req)


	row := db.Sqlite3.QueryRow(`
			SELECT ID,LOGIN_ID,PASSWORD,EMAIL,GROUP_ID
			FROM USER WHERE LOGIN_ID = ?
		`, req.Id)
	if row.Err() != nil {
		c.Response().SetStatusCode(http.StatusInternalServerError)
		c.Response().SetBodyString(row.Err().Error())
		return row.Err()
	}
	var r entity.User
	row.Scan(&r.Id, &r.LoginId, &r.Password, &r.Email, &r.GroupId)

	if bcrypt.CompareHashAndPassword([]byte(r.Password),[]byte(req.Password)) != nil{
		c.Response().SetStatusCode(http.StatusUnauthorized)
		return
	}


	cookie := fiber.Cookie{
		Name:     "token",
		Value:    uuid.New().String(),
		Path:     "/",
		Domain:   "127.0.0.1",
		MaxAge:   0,
		Expires:  time.Time{},
		Secure:   false,
		HTTPOnly: false,
		SameSite: "",
	}
	db.Token[cookie.Value] = struct {
		Id    int
		Admin bool
	}{Id: r.Id, Admin: r.GroupId == 1}


	c.Cookie(&cookie)
	res, err := json.Marshal(&r)
	c.Response().SetBodyString(string(res))
	return err
}
