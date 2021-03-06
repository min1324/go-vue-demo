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

// CreateToken create a user token
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, &claims, err

}
