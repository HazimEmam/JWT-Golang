package utils

import (
	"time"

	"github.com/HazimEmam/JWTtutorial/models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Role string
	jwt.StandardClaims
}

var jwtKey = []byte("Secret_As_Fuk")

func GenerateToken(user models.User) (string, error) {
	expir := time.Now().Add(time.Minute * 15)
	claims := Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject: user.Email,
			ExpiresAt: expir.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "Faild to get the string token", err
	}
	return tokenStr, nil
}

func VerifyToken(token string) (*Claims, error) {
	cliams := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, cliams, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return cliams, nil
}
