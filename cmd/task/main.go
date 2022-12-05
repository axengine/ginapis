package main

import (
	"ginapis/pkg/config"
	"ginapis/pkg/server/task"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/bkit/runner"
	"go.uber.org/zap"
)

func main() {
	if err := config.InitConf(); err != nil {
		panic(err)
	}
	logs.InitQezap(&config.Conf.Logger)
	defer logs.Qezap.Close()

	logs.Qezap.Info("init", zap.Any("config", config.Conf.EraseSensitive()))

	taskSvc := task.NewExampleServer(config.Conf)
	// 开启任务
	if err := runner.RunServer(taskSvc); err != nil {
		logs.Qezap.Error("runner exit", zap.Error(err))
	}
}
