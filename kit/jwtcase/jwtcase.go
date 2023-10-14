package jwtcase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var SignedString = []byte("grapefruit")

type Token struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GetTokenStr ..
func GetTokenStr(ctx context.Context, uid int64) (string, error) {
	if uid == 0 {
		return "", errors.New("uid is empty")
	}
	now := time.Now()
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Token{
		UserID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			Issuer:    "grapefruit",
			Subject:   "grapefruit",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(SignedString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken  ..
func ParseToken(tokenStr string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return SignedString, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
