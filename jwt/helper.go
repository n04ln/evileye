package jwt

import (
	"errors"
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

func GetUserInfoFromJWT(token string, secret string) (*UserInfo, error) {
	var u UserJWT
	_, err := jwt.ParseWithClaims(
		token, &u, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return nil, errors.New("failed to get username from jwt")
	}

	return &u.UserInfo, nil
}
