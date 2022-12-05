package types

import (
	"fmt"
	"ginapis/pkg/server/http/middleware/jwt"
	gojwt "github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	var data = &UserClaims{
		Uid: 123,
		RegisteredClaims: &gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(time.Now().Add(time.Second)),
		},
	}
	tokenString := jwt.GenJWTTokenWithClaims(data)
	log.Println("token", tokenString)

	time.Sleep(time.Second * 2)
	token, err := gojwt.Parse(tokenString, func(token *gojwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hmacSampleSecret"), nil
	})

	if err != nil {
		log.Println("Parse", err)
		return
	}

	if claims, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
		for k, v := range claims {
			log.Println("key=", k, " value=", v)
		}
		//custom := claims[CustomKey]
	} else {
		log.Println("invalid token")
	}
}
