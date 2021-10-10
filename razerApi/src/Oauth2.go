package src

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type userAuth struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	BanDatetime int    `json:"ban_datetime"`
	Bcrypt      string `json:"bcrypt"`
}
type Bearer struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Authentication struct {
	Authentication string `json:"authentication"`
}
type FailAuthentication struct {
	Error string `json:"error"`
}

func CreateToken() {

}

func GetToken(db *sql.DB, c echo.Context) (code int, v interface{}) {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if !(c.FormValue("grant_type") == "password" &&
		c.FormValue("client_id") == "5" &&
		c.FormValue("client_secret") == "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk" &&
		c.FormValue("scope") == "*") {
		return 400, FailAuthentication{"Fail"}
	}
	fmt.Println("pass")

	jsonString, _ := json.Marshal(QueryGetJsonArray(db, "auth.sql", username))
	var s []userAuth
	_ = json.Unmarshal(jsonString, &s)
	if len(s) < 1 {
		return 401, FailAuthentication{"The provided authorization grant"}
	}
	md5Byte := md5.Sum([]byte(password))
	if bcrypt.CompareHashAndPassword([]byte(s[0].Bcrypt), []byte(hex.EncodeToString(md5Byte[:]))) == nil &&
		s[0].BanDatetime == 0 {
		fmt.Println("pass")
		b := Bearer{}
		b.TokenType = "Bearer"
		b.ExpiresIn = 86400 //1day
		b.AccessToken = GenAccessToken(username)
		b.RefreshToken = GenRefreshToken(username)
		_ = QueryOnly(db, "setToken.sql", s[0].Id, b.AccessToken, b.RefreshToken)

		return 200, b
	}
	//fmt.Println(s)

	return 401, FailAuthentication{"The provided authorization grant"}
}

type Oauth struct {
	Userid       int    `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	CreateTime   string `json:"cteateTime"`
	LifeTime     string `json:"lifeTime"`
	Active       int    `json:"active"`
}

func CheckTokenAlive(db *sql.DB, token string) ([]Oauth, bool) {
	jsonString, _ := json.Marshal(QueryGetJsonArray(db, "checkTokenAlive.sql", token))
	var s []Oauth
	fmt.Println(s)
	_ = json.Unmarshal(jsonString, &s)
	if len(s) < 1 {
		return s, false
	}
	return s, s[0].AccessToken == token
}

func RefreshToken(db *sql.DB, token string) {

}
