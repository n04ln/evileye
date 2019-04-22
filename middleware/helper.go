package middleware

import (
	"time"

	"github.com/NoahOrberg/evileye/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	ID        int64
	ExpiredAt int64 // UNIXTIME
}

type UserJWT struct {
	UserInfo
	jwt.StandardClaims
}

func CreateJWTToken(u entity.User, secret string) (string, error) {
	ui := UserInfo{
		ID:        u.ID,
		ExpiredAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtModel := UserJWT{
		UserInfo: ui,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtModel)
	return token.SignedString([]byte(secret))
}
