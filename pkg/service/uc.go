package service

import (
	"context"
	"ginapis/pkg/server/http/bean"
	"ginapis/pkg/server/http/middleware/jwt"
	"ginapis/pkg/types"
	gojwt "github.com/golang-jwt/jwt/v4"
	"time"
)

func (svc *Service) Login(ctx context.Context, in *bean.LoginReq, out *bean.LoginResp) error {
	expire := time.Now().Add(time.Hour * 24 * 7)
	token := jwt.GenJWTTokenWithClaims(types.UserClaims{
		Uid: 123,
		RegisteredClaims: &gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(expire),
		},
	})
	out.Token = token
	out.Expire = expire
	return nil
}

func (svc *Service) UserInfo(ctx context.Context, in *bean.GetUserInfoReq, out *bean.GetUserInfoResp) error {
	out.Balance = 123
	return nil
}
