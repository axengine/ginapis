package main

import (
	"context"
	_ "ginapis/docs/api"
	"ginapis/pkg/config"
	"ginapis/pkg/dao"
	"ginapis/pkg/server/http/api"
	"ginapis/pkg/service"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/bkit/runner"
	"go.uber.org/zap"
	"log"
)

// @title api server
// @version 1.0.0
// @description

// @host :8080
// @BasePath /
func main() {
	if err := config.InitConf(); err != nil {
		panic(err)
	}
	logs.InitQezap(&config.Conf.Logger)
	defer logs.Qezap.Close()
	logs.Qezap.Info("init", zap.Any("config", config.Conf.EraseSensitive()))

	d := dao.New(config.Conf)
	defer d.Close()

	svc := service.New(config.Conf, d)
	defer svc.Close()

	ctx := context.Background()

	if err := runner.RunServer(api.NewApiHttpServer(config.Conf, svc),
		runner.WithListenAddr(config.Conf.Api.HttpListenAddr),
		runner.WithContext(ctx)); err != nil {
		log.Printf("runner exit: %v\n", err)
	}
}
