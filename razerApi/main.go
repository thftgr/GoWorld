package main

import (
	"./src"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"strings"
)

var DB *sql.DB

//var MongoDB *mongo.Client

type Authentication struct {
	Authentication string `json:"authentication"`
}

func init() {
	b, err := ioutil.ReadFile("./databaseSetting.json")
	errCheck(err)
	// json 데이터 읽어서 저장
	var data struct {
		Sql struct {
			Id     string
			Pw     string
			Url    string
			Driver string
		}
		Mongo struct {
			Id     string
			Pw     string
			Url    string
			Driver string
		}
	} // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언
	_ = json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	db, err := sql.Open(data.Sql.Driver, data.Sql.Id+":"+data.Sql.Pw+"@tcp("+data.Sql.Url+")/")
	if !errCheck(err) {
		DB = db
	}
	//fmt.Println(data)

	//MongoDB = src.ConnectMongo(data.Mongo.Id, data.Mongo.Pw, data.Mongo.Driver, data.Mongo.Url)
	//type dat	struct {
	//	ID int
	//	Data []int
	//}
	//var datt = dat{}
	//time.Sleep(time.Second * 2)
	//_ = MongoDB.Database("Debian").Collection("Oauth2").FindOne(context.TODO(), bson.M{"id": 1000}).Decode(&datt)

	//rr.Elements()

}

func main() {
	//2018-07-06T06:33:34+00:00
	//파일 read

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("oauth/token", func(c echo.Context) error {
		return c.JSON(src.GetToken(DB, c))
	})

	e.GET("/api/v2/users/:id/:mode", func(c echo.Context) error {
		id := c.Param("id")
		//id := 7986
		var mode string
		if c.QueryParam("rx") == "true" {
			mode = c.Param("mode") + ".relax"
		} else {
			mode = c.Param("mode")
		}

		sqlFile := src.ParseStringModeSQLFile(mode)
		b, err := json.Marshal(src.QueryGetJsonObject(DB, sqlFile, id, id, id, id, id, id, id, id, id, id))
		fmt.Println(string(b))
		errCheck(err)
		var dat src.User
		_ = json.Unmarshal(b, &dat)
		src.AddHeaderDEV(c)
		return c.JSON(200, dat)
	})

	e.GET("/api/v2/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		//id := 7986
		b, err := json.Marshal(src.QueryGetJsonObject(DB, "users.osu.sql", id, id, id, id, id, id, id, id, id, id))
		if err != nil {
			return c.JSON(401, "error")
		}
		var dat src.User
		_ = json.Unmarshal(b, &dat)
		src.AddHeaderDEV(c)

		return c.JSON(200, dat)
	})

	e.GET("/api/v2/me", func(c echo.Context) error {

		Token := strings.TrimLeft(c.Request().Header.Get("Authorization"), "Bearer ")

		if len(Token) != 512 { //토큰 길이 체크
			return c.JSONPretty(401, Authentication{"Basic"}, "  ")
		}

		//토큰이 유효한지 체크
		aut, pass := src.CheckTokenAlive(DB, Token)
		if !pass {
			return c.JSONPretty(401, Authentication{"Basic"}, "  ")
		}
		id := aut[0].Userid

		b, err := json.Marshal(src.QueryGetJsonObject(DB, "user.mode.favourite.sql", id))
		mod := map[string]int{}
		_ = json.Unmarshal(b, &mod)

		var mode int
		if c.QueryParam("rx") == "true" {
			mode = mod["favourite_mode"] + 4
		} else {
			mode = mod["favourite_mode"]
		}

		sqlFile := src.ParseIntModeSQLFile(mode)

		b, err = json.Marshal(src.QueryGetJsonObject(DB, sqlFile, id, id, id, id, id, id, id, id, id, id))

		if err != nil {
			return c.JSON(401, "")
		}
		var dat src.User
		_ = json.Unmarshal(b, &dat)

		src.AddHeaderDEV(c)
		return c.JSON(200, dat)

	})

	e.GET("/api/v2/users/:id/beatmapsets/most_played", func(c echo.Context) error {
		var ofs = "0"
		var lim = "5"

		if c.QueryParam("offset") != "" {
			ofs = c.QueryParam("offset")
		}
		if c.QueryParam("limit") != "" {
			lim = c.QueryParam("limit")
		}

		b, err := json.Marshal(src.QueryGetJsonArray(DB, "users.id.beatmapsets.most_played.sql", c.Param("id"), ofs, lim))
		if err != nil {
			return c.JSON(401, "")
		}
		fmt.Println(string(b))
		var dat src.MostPlayed
		_ = json.Unmarshal(b, &dat)

		src.AddHeaderDEV(c)
		return c.JSON(200, dat)

	})
	e.GET("/api/v2/users/:id/scores/best", func(c echo.Context) error {
		var ofs = "0"
		var lim = "5"

		if c.QueryParam("offset") != "" {
			ofs = c.QueryParam("offset")
		}
		if c.QueryParam("limit") != "" {
			lim = c.QueryParam("limit")
		}
		b, err := json.Marshal(src.QueryGetJsonArray(DB, "users.id.scores.best.sql", c.Param("id"), ofs, lim, c.Param("id"), c.Param("id")))
		//b, err := json.Marshal(src.QueryGetJsonArray(DB, "users.id.scores.best.sql", c.Param("id"), ofs, lim))
		if err != nil {
			return c.JSON(401, "")
		}
		fmt.Println(string(b))
		var dat src.BestPlay
		_ = json.Unmarshal(b, &dat)

		src.AddHeaderDEV(c)
		return c.JSON(200, dat)
	})

	e.Logger.Fatal(e.Start(":8002")) // localhost:8002

}

func errCheck(err error) (e bool) {
	e = err != nil
	if e {
		log.Println(err)
	}
	return
}
