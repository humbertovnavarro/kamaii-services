package tokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var TOKEN_TTL = time.Hour * 2

type UUID = string

type SessionTokenClaims = jwt.StandardClaims

var TOKEN_SECRET []byte

func init() {
	TOKEN_SECRET = []byte(os.Getenv("SESSION_TOKEN_SECRET"))
}

func GenerateUserToken(uuid UUID) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        uuid,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + int64(TOKEN_TTL),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SECRET)
}

func ParseUserToken(tokenString string) (*SessionTokenClaims, error) {
	claims := SessionTokenClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return TOKEN_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}
