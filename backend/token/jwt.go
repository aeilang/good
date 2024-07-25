package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	accessKey  []byte
	refreshKey []byte
}

type Tokener interface {
	CreateAccessToken(data Data) (string, error)
	CreateRefreshToken(data Data) (string, error)
	VerifyAccessToken(tokenString string) (*Payload, error)
	VerifyRefreshToken(tokenString string) (*Payload, error)
}

func New(accessKey, refreshKey string) *Token {
	return &Token{
		accessKey:  []byte(accessKey),
		refreshKey: []byte(refreshKey),
	}
}

type Data struct {
	Email    string
	Role     string
	Duration time.Duration
}

func createToken(input Data, secretKey []byte) (string, error) {
	claims := Payload{
		Email: input.Email,
		Role:  input.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(input.Duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ts, err := token.SignedString(secretKey)
	return ts, err
}

func (t *Token) CreateAccessToken(input Data) (string, error) {
	return createToken(input, t.accessKey)
}

func (t *Token) CreateRefreshToken(input Data) (string, error) {
	return createToken(input, t.refreshKey)
}

func verifyToken(tokenString string, secretKey []byte) (*Payload, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(_ *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if payload, ok := token.Claims.(*Payload); ok {
		return payload, nil
	}

	return nil, fmt.Errorf("token is not correct")
}

func (t *Token) VerifyAccessToken(tokenString string) (*Payload, error) {
	return verifyToken(tokenString, t.accessKey)
}

func (t *Token) VerifyRefreshToken(tokenString string) (*Payload, error) {
	return verifyToken(tokenString, t.refreshKey)
}
