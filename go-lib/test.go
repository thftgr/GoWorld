package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thftgr/go-lib/jwt2"
)

var KEY = "server sign key"

func main() {

	fmt.Println("=================================토큰 생성========================================")
	claim1 := jwt2.Claims{
		Iss: "account.thftgr",
		Aud: "uid1",
	}
	token, rToken, err := claim1.CreateToken(jwt.SigningMethodHS256, KEY)
	fmt.Println(" token: ", token)
	fmt.Println("rToken: ", rToken)
	fmt.Println("   err: ", err)
	fmt.Println("=================================================================================")

	fmt.Println("===================================토큰 검증=======================================")
	claim2, err := jwt2.GetClaims(token)

	err = claim2.CheckJwt(jwt.SigningMethodHS256, token, KEY)
	bb := jwt2.CheckRefreshToken(token, rToken, KEY)
	fmt.Println(" token: ", err == nil, "err:", err)
	fmt.Println("rToken: ", bb)
	fmt.Println("=================================================================================")

	fmt.Println("===================================토큰 갱신=======================================")
	token, rToken, err = claim2.RefreshToken(jwt.SigningMethodHS256, KEY)
	fmt.Println(" token: ", token)
	fmt.Println("rToken: ", rToken)
	fmt.Println("   err: ", err)
	fmt.Println("=================================================================================")

}
