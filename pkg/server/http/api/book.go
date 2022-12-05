package api

import (
	"ginapis/pkg/server/http/bean"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/gin-gonic/gin"
)

// @Summary 书籍查询
// @Description 书籍查询
// @Tags API BOOK
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param Request body bean.ListBookReq true "request param"
// @Success 200 {array} bean.Book "success"
// @Router /v1/book/list [GET]
func listBook(c *gin.Context) {
	in := &bean.ListBookReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	if err := svc.ListBook(c.Request.Context(), in, out); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}
