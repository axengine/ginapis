package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CustomClaims 自定义Claims，使用时修改，增加字段
type CustomClaims struct {
	Uid int64 `json:"uid"`
	*jwt.RegisteredClaims
}

// hmacSecret HMAC加密密钥
var hmacSecret = []byte("hmacSampleSecret")

// NewMiddleware 创建中间件，secret为加密密钥
func NewMiddleware(secret string) gin.HandlerFunc {
	hmacSecret = []byte(secret)

	return func(c *gin.Context) {
		token, err := jwt.Parse(c.GetHeader("Authorization"), func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSecret, nil
		})
		if err != nil {
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			for k, v := range claims {
				c.Set(k, v)
			}
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func GenJWTTokenWithClaims(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}
	return signedToken
}

func GenJWTTokenWithCustomClaims(claims *CustomClaims, expire time.Time) string {
	claims.RegisteredClaims = &jwt.RegisteredClaims{
		Issuer:    "test",
		Subject:   "somebody",
		Audience:  []string{"somebody_else"},
		ExpiresAt: jwt.NewNumericDate(expire),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}
	return signedToken
}
