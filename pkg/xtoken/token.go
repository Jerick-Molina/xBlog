package xtoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Secret string

func (secret Secret) ReadToken(bearerToken string) error {

	if len(bearerToken) == 0 {

		bearer := bearerToken[:6]
		tokenParsed := bearerToken[7:]
		if bearer != "Bearer" {
			return fmt.Errorf("token is not of type bearer")
		}

		token, err := jwt.ParseWithClaims(tokenParsed, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil }, jwt.WithLeeway(5*time.Second))

		if err != nil {

			return err
		}

		if _, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("token empty")
}
