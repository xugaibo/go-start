package jwtutil

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	Id   uint
	Name string
	jwt.RegisteredClaims
}

var hmacSampleSecret = []byte("FanOne")

// GenerateToken 签发用户Token
func GenerateToken(id uint, username string) (string, error) {
	// Create the Claims
	claims := JwtToken{
		id, username,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(24 * time.Hour)},
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(hmacSampleSecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(tokenStr string) (*JwtToken, error) {
	jwtToken := JwtToken{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtToken); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}

func (jwt *JwtToken) IsExpire() bool {
	return time.Now().Unix() > jwt.ExpiresAt.Unix()
}
