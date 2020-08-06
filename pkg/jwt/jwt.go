package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"gin-base/pkg/e"
)

var (
	secret = viper.GetString("app.secret")
)

type JwtContext struct {
	UserID uint64
	jwt.StandardClaims
}

func IssueToken(j JwtContext) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": j.UserID,
	})
	tokenString, err = token.SignedString([]byte(secret))

	return
}

func ParseToken(tokenString string) (*JwtContext, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtContext{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, e.TokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, e.TokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, e.TokenNotValidYet
			} else {
				return nil, e.TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*JwtContext); ok && token.Valid {
			return claims, nil
		}
		return nil, e.TokenInvalid
	} else {
		return nil, e.TokenInvalid
	}
}
