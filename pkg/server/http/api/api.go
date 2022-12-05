package api

import (
	"ginapis/pkg/config"
	"ginapis/pkg/service"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/runner"
)

var (
	svc *service.Service
	cfg *config.Config
)

func NewApiHttpServer(c *config.Config, s *service.Service) runner.Server {
	svc = s
	cfg = c

	midFlag := ginutil.MStd
	if cfg.Release() {
		midFlag = ginutil.MRelease | ginutil.MTraceId | ginutil.MRecoverLogger
	}

	httpHandler := ginutil.DefaultEngine(midFlag)

	registerApiRouter(httpHandler)

	return runner.NewHttpServer(httpHandler)
}
