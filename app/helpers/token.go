package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("API_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, time.Time, error) {
	expiredTime := time.Now().Add(1440 * time.Minute)
	expiredAt := expiredTime.Unix()

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expiredTime, err
}

func DecodeToken(tokenStr string) (string, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
	}

	if claims.Username == "" {
		return "", err
	}

	return claims.Username, nil
}
