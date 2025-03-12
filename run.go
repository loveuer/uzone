package uzone

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/tool"
)

func (u *uzone) startAPI(ctx context.Context) {
	_, cfg := u.api.GetUZone()

	fmt.Printf("Uzone | api listen at %s\n", cfg.Address)

	go func() {
		if err := u.api.Run(u.UseCtx()); err != nil {
			u.UseLogger().Panic("start api failed, err = %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		_ = u.api.Shutdown(tool.Timeout(2))
	}()
}

func (u *uzone) startTask(ctx context.Context) {
	fmt.Printf("Uzone | start task channel[%02d]\n", len(u.taskCh))
	for _, _ch := range u.taskCh {
		go func(ch <-chan func(interfaces.Uzone) error) {
			var err error
			for {
				select {
				case <-ctx.Done():
				case task, ok := <-ch:
					if !ok {
						return
					}

					if err = task(u); err != nil {
						u.UseLogger(ctx).Error(err.Error())
					}
				}
			}
		}(_ch)
	}
}

func (u *uzone) Run(ctx context.Context) {
	u.RunSignal(ctx)
}

func (u *uzone) runInitFns(ctx context.Context) {
	for _, fn := range u.initFns._sync {
		fn(u)
	}
}

func (u *uzone) startInitFns(ctx context.Context) {
	for _, fn := range u.initFns._async {
		go fn(u)
	}
}

func (u *uzone) RunSignal(ctxs ...context.Context) {
	c := context.Background()
	if len(ctxs) > 0 {
		c = ctxs[0]
	}

	ctx, cancel := signal.NotifyContext(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	u.ctx = ctx

	print(Banner)

	if len(u.initFns._sync) > 0 {
		u.runInitFns(ctx)
	}

	if len(u.initFns._async) > 0 {
		u.startInitFns(ctx)
	}

	if u.api != nil {
		u.startAPI(ctx)
	}

	if len(u.taskCh) > 0 {
		u.startTask(ctx)
	}

	<-ctx.Done()

	u.UseLogger().Warn(" Upp | quit by signal...")
	if u.cache != nil {
		u.cache.Close()
	}

	<-tool.Timeout(2).Done()
}
