package middleware

import (
	"ATIKAH-PERCA/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(adminID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 96).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
