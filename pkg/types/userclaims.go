package types

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Uid int64 `json:"uid"`
	*jwt.RegisteredClaims
}
