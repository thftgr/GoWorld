package src

import (
	"github.com/labstack/echo/v4"
	"time"
)

var banIps = map[string]int{}

func Ban(ip string,sec int){
	banIps[ip] = time.Now().Second() + sec
}
func IsBan(ip string) bool {
	return time.Now().Second() < banIps[ip]
}
func UnBan(ip string) {
	delete(banIps,ip)
}
func CheckBan(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if IsBan(c.RealIP()) {
			return echo.ErrForbidden
		}
		return next(c)
	}
}
