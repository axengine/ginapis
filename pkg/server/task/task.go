package task

import (
	"context"
	"ginapis/pkg/config"
	"github.com/bbdshow/bkit/runner"
	"time"
)

func NewExampleServer(cfg *config.Config) runner.Server {
	taskSvc := runner.NewTaskServer()
	_ = taskSvc.AddCronFunc("0 0 0 * * ? ", func() {
		// do something
	})
	_ = taskSvc.AddOnceTimeAfterFunc(time.Second, func(ctx context.Context) {
		// do something
	})

	return taskSvc
}
