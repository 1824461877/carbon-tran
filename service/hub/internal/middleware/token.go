package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type AuthJwt struct {
	Id       string `json:"uid"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func NewJwt(key string) *JWT {
	return &JWT{
		[]byte(key),
	}
}

func (j *JWT) CreateToken(authJwt AuthJwt) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authJwt)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*AuthJwt, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthJwt{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*AuthJwt); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
