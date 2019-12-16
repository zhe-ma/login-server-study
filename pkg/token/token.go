package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Context struct {
	ID       uint64
	Username string
}

func Parse(tokenString string) (*Context, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		secret := viper.GetString("jwt_secret")
		return []byte(secret), nil

	})

	context := &Context{}
	if err != nil {
		return context, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		context.ID = uint64(claims["id"].(float64))
		context.Username = claims["username"].(string)
		return context, nil
	}

	return context, jwt.ErrSignatureInvalid
}

func Sign(c Context) (tokenString string, err error) {
	secret := viper.GetString("jwt_secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))

	return
}
