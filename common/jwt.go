package common

import (
	"demo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_jwt")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func CreateToken(u *model.User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour * 3).Unix()
	clains := &Claims{
		UserId: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "min",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, clains)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
