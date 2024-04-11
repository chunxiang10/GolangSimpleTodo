package handler

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var MyKey = []byte("123123123123")

type MyClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

func SendToken(uid string) (string, error) {
	// MyKey := []byte("123123123123")
	t := MyClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "jack",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	sss, e := token.SignedString(MyKey)
	return sss, e
}

func ParseToken(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if err != nil {
		return "0", err
	}
	claims := t.Claims.(*MyClaims).RegisteredClaims
	if claims.Issuer != "jack" {
		return "0", nil
	}
	return t.Claims.(*MyClaims).Uid, nil

}
