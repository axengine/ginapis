package api

import (
	"ginapis/pkg/server/http/middleware/jwt"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func registerApiRouter(e *gin.Engine) {
	store := persistence.NewInMemoryStore(time.Second)
	authHandle := jwt.NewMiddleware(cfg.Api.JWTSecret)
	v1 := e.Group("/v1")
	v1.POST("/login", login)

	v1.GET("/user/info", userInfo).Use(authHandle)

	bookG := v1.Group("/book/").Use(authHandle)
	bookG.GET("/list", cache.CachePageAtomic(store, time.Second*30, listBook))

}
