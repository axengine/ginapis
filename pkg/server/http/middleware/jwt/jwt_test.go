package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	var data = CustomClaims{
		Uid:              123,
		RegisteredClaims: nil,
	}
	tokenString := GenJWTTokenWithCustomClaims(&data, time.Now().Add(time.Second*2222))

	//time.Sleep(time.Second * 3)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		log.Println("Parse", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for k, v := range claims {
			log.Println("key=", k, " value=", v)
		}
		//custom := claims[CustomKey]
		//log.Println(custom)
		//log.Println(reflect.TypeOf(claims))
		//log.Println(reflect.TypeOf(custom))

		//log.Println("claims[CustomKey].(*customClaim):", claims[CustomKey].(*customClaim))
	} else {
		log.Println("invalid token")
	}
}
