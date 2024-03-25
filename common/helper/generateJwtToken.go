package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type GenerateJwtStruct struct {
	Id       uint
	Uid      string
	Username string
	jwt.RegisteredClaims
}

func GenerateJwtToken(g *GenerateJwtStruct, secretKey string, expire int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		GenerateJwtStruct{
			Id:       g.Id,
			Uid:      g.Uid,
			Username: g.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire * 1000 * 1000))),
			},
		},
	)

	return claims.SignedString([]byte(secretKey))
}

func ParseJwtToken(tokenStr string, secretKey string) (*GenerateJwtStruct, error) {
	jwtStruct := GenerateJwtStruct{}
	if _, err := jwt.ParseWithClaims(
		tokenStr,
		&jwtStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	); err != nil {
		return nil, err
	} else {
		return &jwtStruct, nil
	}
}
