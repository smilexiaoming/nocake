package common

import (
	"nocake/global"
	"time"

	"github.com/golang-jwt/jwt"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToke 生成Token
func GenerateToke(username string) (string, error) {
	claims := Claims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}

// GenerateSkey 生成Skey
func GenerateSkey(open_id string) (string, error) {
	claims := Claims{open_id, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    open_id,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifySkey 验证Skey
func VerifySkey(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}
