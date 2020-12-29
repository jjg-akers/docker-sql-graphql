package auth

import (
	"errors"
	"fmt"

	jwtv4 "github.com/dgrijalva/jwt-go/v4"
	"github.com/graphql-go/graphql/gqlerrors"
)

var jwtSecret []byte = []byte("i'm a secret")

func CreateToken(username, password string) (string, error) {
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, jwtv4.MapClaims{
		"username": username,
		"password": password,
	})

	if tokenStr, err := token.SignedString(jwtSecret); err != nil {
		return "", gqlerrors.FormatError(err)
	} else {
		return tokenStr, nil
	}

}

func ValidateToken(t string) (bool, error) {
	if t == "" {
		return false, gqlerrors.FormatError(errors.New("Authorization token must be present"))
	}

	token, _ := jwtv4.Parse(t, func(token *jwtv4.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv4.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	if _, ok := token.Claims.(jwtv4.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, gqlerrors.FormatError(errors.New("Invalid authorization token"))
	}
}
