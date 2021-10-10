package src

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strconv"
	"time"
)

func GenAccessToken(username string) (Token string) {
	ts := time.Now().UnixNano()
	tim := strconv.FormatInt(ts, 10)

	bs := sha3.Sum512([]byte(username + tim))
	Token = hex.EncodeToString(bs[:])
	bs = sha3.Sum512([]byte(tim))
	Token += hex.EncodeToString(bs[:])
	bs = sha3.Sum512(bs[:])
	Token += hex.EncodeToString(bs[:])
	bs = sha3.Sum512([]byte(tim + tim))
	Token += hex.EncodeToString(bs[:])
	fmt.Println(Token)
	return Token
}
func GenRefreshToken(username string) (Token string) {
	ts := time.Now().Unix()
	tim := strconv.FormatInt(ts, 10)
	tim += "Refresh"

	bs := sha3.Sum512([]byte(username + tim))
	Token = hex.EncodeToString(bs[:])
	bs = sha3.Sum512([]byte(tim))
	Token += hex.EncodeToString(bs[:])
	bs = sha3.Sum512(bs[:])
	Token += hex.EncodeToString(bs[:])
	bs = sha3.Sum512([]byte(tim + tim))
	Token += hex.EncodeToString(bs[:])
	fmt.Println(Token)
	return Token
}
