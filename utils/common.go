package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}
