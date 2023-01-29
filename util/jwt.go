package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtSecret string = "toy-tok"

type Claims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

func CreateAccessToken(userID int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * 60 * time.Minute)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "toytok-app",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tokenClaims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
