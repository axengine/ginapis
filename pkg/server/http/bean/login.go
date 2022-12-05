package bean

import "time"

type LoginReq struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResp struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type GetUserInfoReq struct {
	Uid int64 `form:"-"`
}

type GetUserInfoResp struct {
	Balance float64 `json:"balance"`
}
