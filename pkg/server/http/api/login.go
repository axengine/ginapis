package api

import (
	"ginapis/pkg/server/http/bean"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Description login
// @Tags API 登陆
// @Accept json
// @Produce json
// @Param Request body bean.LoginReq true "request param"
// @Success 200 {object} bean.LoginReq "success"
// @Router /v1/login [POST]
func login(c *gin.Context) {
	in := &bean.LoginReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &bean.LoginResp{}
	if err := svc.Login(c.Request.Context(), in, out); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary 用户信息查询
// @Description 用户信息查询
// @Tags API
// @Accept json
// @Produce json
// @Param Request body bean.GetUserInfoReq true "request param"
// @Success 200 {object} bean.GetUserInfoResp "success"
// @Router /v1/user/info [GET]
func userInfo(c *gin.Context) {
	in := &bean.GetUserInfoReq{Uid: c.GetInt64("uid")}
	out := &bean.GetUserInfoResp{}
	if err := svc.UserInfo(c.Request.Context(), in, out); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}
