package refreshtoken

import (
	"github.com/golang-jwt/jwt/v4"
	"go-start/core/bizcode"
	"go-start/core/bizerror"
	"go-start/core/jwtutil"
	"time"
)

var hmacSampleSecret = []byte("FanOne")

type RefreshToken struct {
	Token string
	jwt.RegisteredClaims
}

func Generate(token string) (string, error) {
	claims := RefreshToken{
		token,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(30 * 24 * time.Hour)},
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(hmacSampleSecret)
	return token, err
}

func Parse(tokenStr string) (*RefreshToken, error) {
	refreshToken := RefreshToken{}
	token, err := jwt.ParseWithClaims(tokenStr, &refreshToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshToken); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}

func (t *RefreshToken) IsExpire() bool {
	return time.Now().Unix() > t.ExpiresAt.Unix()
}

func Refresh(refreshTokenStr string) string {
	refreshToken, err := Parse(refreshTokenStr)
	if err != nil || refreshToken == nil {
		panic(bizerror.Biz(bizcode.RefreshTokenInvalid))
	}

	if refreshToken.IsExpire() {
		panic(bizerror.Biz(bizcode.RefreshTokenExpire))
	}

	token, err := jwtutil.ParseToken(refreshToken.Token)
	if err != nil || token == nil {
		panic(err)
	}

	if !token.IsExpire() {
		return refreshToken.Token
	}

	generateToken, err := jwtutil.GenerateToken(token.Id, token.Name)
	if err != nil {
		panic(err)
	}

	return generateToken
}
