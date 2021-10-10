package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

const (
	InvalidToken = `invalid token`
	ExpiredToken = `expired token`
)

type Claims struct {
	Iss      string        `json:"iss"` //iss: 토큰 발급자 (issuer)
	Sub      string        `json:"sub"` //sub: 토큰 제목 (subject)
	Aud      string        `json:"aud"` //aud: 토큰 대상자 (audience) //UID
	Exp      int64         `json:"exp"` //exp: 토큰의 만료시간 (expiraton), millisecond 언제나 현재 시간보다 이후로 설정되어있어야합니다.
	Nbf      int64         `json:"nbf"` //nbf: 토큰 활성시간 (Not Before)
	Iat      int64         `json:"iat"` //iat: 토큰이 발급된 시간 (issued at), 이 값을 사용하여 토큰의 age 가 얼마나 되었는지 판단 할 수 있습니다.
	Jti      string        `json:"jti"` //jti: JWT의 고유 식별자로서, 주로 중복적인 처리를 방지하기 위하여 사용됩니다. 일회용 토큰에 사용하면 유용합니다.
	MapClaim jwt.MapClaims // 위 항목 외 추가 항목 추가
}

var lastUse = base64.RawStdEncoding.DecodeString

func GetClaims(token *string) (c *Claims, err error) {
	t := strings.ReplaceAll(*token, "Bearer ", "")
	t = strings.ReplaceAll(t, "bearer ", "")
	s := strings.Split(t, ".")
	if len(s) != 3 {
		return
	}
	b, err := lastUse(s[1])
	if err == nil {
		return unmashal(&b)
	}

	b, err = base64.RawStdEncoding.DecodeString(s[1])
	if err == nil {
		return unmashal(&b)
	}

	b, err = base64.StdEncoding.DecodeString(s[1])
	if err == nil {
		return unmashal(&b)
	}

	b, err = base64.RawURLEncoding.DecodeString(s[1])
	if err == nil {
		return unmashal(&b)
	}

	b, err = base64.URLEncoding.DecodeString(s[1])
	if err == nil {
		return unmashal(&b)
	}

	return
}

func unmashal(b *[]byte) (c *Claims, err error) {
	if err = json.Unmarshal(*b, &c); err != nil {
		return
	}

	err = json.Unmarshal(*b, &c.MapClaim)
	return
}

func (v *Claims) CheckJwt(method jwt.SigningMethod, token, key *string) (err error) {

	t, err := jwt.NewWithClaims(method, v.MapClaim).SignedString([]byte(*key))
	if err != nil {
		return
	}
	if t != *token {
		return errors.New(InvalidToken)
	}
	if v.Exp < time.Now().UnixMilli() {
		return errors.New(ExpiredToken)
	}

	return
}

func (v *Claims) CreateToken(method jwt.SigningMethod, SigningKey string) (token string, err error) {
	now := time.Now().UnixMilli()
	if v.Nbf <= now {
		v.Nbf = now
	} else {
		v.MapClaim["nbf"] = v.Nbf
	}

	if v.Iat <= now {
		v.Iat = now
	} else {
		v.MapClaim["iat"] = v.Iat
	}

	if v.Exp <= now {
		v.Exp = now + (time.Minute.Milliseconds() * 10)
	} else {
		v.MapClaim["exp"] = v.Exp
	}

	if v.Jti != "" {
		v.MapClaim["jti"] = v.Jti
	}

	if v.Iss != "" {
		v.MapClaim["iss"] = v.Iss
	}

	if v.Sub != "" {
		v.MapClaim["sub"] = v.Sub
	}

	v.MapClaim["aud"] = v.Aud

	return jwt.NewWithClaims(method, v.MapClaim).SignedString([]byte(SigningKey))

}
