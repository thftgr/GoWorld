package main

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/JWT"
	"thftgr.com/GoWorld/debianWeb/Logger"
	"thftgr.com/GoWorld/debianWeb/Route"
	"thftgr.com/GoWorld/debianWeb/mail"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/src"
)

var LogIO = bytes.Buffer{}

func init() {
	src.LoadSetting()

	mariadb.Connect()

	go Logger.LoadLogger(&LogIO)

}

func main() {

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{Output: &LogIO}),
		middleware.RequestID(),
		//middlewareFunc.CSRF(),
		//middlewareFunc.CSRFConfig(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{src.Config.Server.Domain},
			AllowMethods: []string{http.MethodGet, http.MethodPost},
		}),
		//middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(60)),
	)
	//TODO PUBLIC AREA
	e.Static("/", "web")
	e.GET("/api/v1/beatmap/scores", Route.BeatmapScores, src.CheckBan)
	e.POST("/api/v1/leaderboard", Route.LeaderBoard, src.CheckBan)
	//TODO PUBLIC AREA

	//TODO AUTH AREA
	//TODO JWT 토큰 타임아웃 처리,로그아웃시 램,db 넣기
	e.POST("/api/v1/register", Route.Register, src.CheckBan, Route.ReCaptcha)
	e.POST("/api/v1/resetPassword", Route.ResetPassword, src.CheckBan, Route.ReCaptcha)
	e.POST("/api/v1/login", Route.Login, src.CheckBan, Route.ReCaptcha)

	e.POST("/api/v1/logout", Route.Logout, JWT.Jwt) // 쿠키에서 토큰 만료시킴

	//e.POST("/api/v1/refreshToken",nil) //리프레시토큰 사용기간 < 1day 새로 발급
	//현재상태

	e.GET("/api/v1/account/verify/:userid", Route.AccountVerify, src.CheckBan)
	e.GET("/api/v1/account/disable/:userid", Route.AccountDisable, src.CheckBan)
	//TODO AUTH AREA

	//TODO TEST AREA
	e.GET("/api/v1/testToken", Route.TestTokenCheck, src.CheckBan, JWT.Jwt) // 토큰 유효 여부, 개발용 삭제예정
	e.GET("/api/v1/sendVerifyMail", mail.SendCertificationMail, src.CheckBan, JWT.Jwt)

	//TODO TEST AREA

	e.Logger.Fatal(e.Start(":" + src.Config.Server.Port)) // localhost:80
}
