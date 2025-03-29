package api_fiber

import (
	"context"
	"crypto/tls"
	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"
	"gorm.io/gorm"
	"net"

	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/tool"
	"github.com/samber/lo"
)

type Engine struct {
	*fiber.App
	zone interfaces.Uzone
	cfg  api.Config
}

func (e *Engine) SetAddress(address string) {
	e.cfg.Address = address
}

func (e *Engine) SetListener(ln net.Listener) {
	e.cfg.Listener = ln
}

func (e *Engine) SetTLSConfig(cfg *tls.Config) {
	e.cfg.TLSConfig = cfg
}

func (e *Engine) SetRecover(recover bool) {
	e.cfg.Recover = recover
}

func (e *Engine) SetUZone(u interfaces.Uzone) {
	e.zone = u
}

func (e *Engine) GetUZone() (interfaces.Uzone, api.Config) {
	return e.zone, e.cfg
}

func (e *Engine) UseLogger() *log.UzoneLogger {
	return e.zone.UseLogger()
}
func (e *Engine) UseDB() *gorm.DB {
	return e.zone.UseDB()
}
func (e *Engine) UseCache() cache.Cache {
	return e.zone.UseCache()
}
func (e *Engine) UseES() *es.Client {
	return e.zone.UseES()
}
func (e *Engine) UseMQ() *mq.Client {
	return e.zone.UseMQ()
}

func (e *Engine) Group(path string, handlers ...api.Handler) api.ApiGroup {
	hs := newHandlers(e, false, handlers...)
	group := e.App.Group(path, hs...)
	return NewGroup(e, group)
}

func (e *Engine) GET(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Get(path, hs[0], hs[1:]...)
}

func (e *Engine) POST(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Post(path, hs[0], hs[1:]...)
}

func (e *Engine) PUT(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Put(path, hs[0], hs[1:]...)
}

func (e *Engine) DELETE(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Delete(path, hs[0], hs[1:]...)
}

func (e *Engine) HEAD(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Head(path, hs[0], hs[1:]...)
}

func (e *Engine) PATCH(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Patch(path, hs[0], hs[1:]...)
}

func (e *Engine) OPTIONS(path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Options(path, hs[0], hs[1:]...)
}

func (e *Engine) Handle(method, path string, handlers ...api.Handler) {
	hs := newHandlers(e, true, handlers...)
	e.App.Add([]string{method}, path, hs[0], hs[1:]...)
}

func (e *Engine) Use(handlers ...api.Handler) {
	for _, item := range newHandlers(e, true, handlers...) {
		e.App.Use(item)
	}
}

func (e *Engine) Run(ctx context.Context) error {
	cfg := fiber.ListenConfig{
		DisableStartupMessage: true,
		EnablePrintRoutes:     false,
		TLSConfigFunc: func(tlsConfig *tls.Config) {
			tlsConfig = e.cfg.TLSConfig
		},
		BeforeServeFunc: func(app *fiber.App) error {
			var rs api.Routes = lo.Map(app.GetRoutes(true), func(item fiber.Route, idx int) api.Route {
				return api.Route{
					Method:      item.Method,
					Path:        item.Path,
					HandlerName: tool.GetFunctionName(item.Handlers[len(item.Handlers)-1]),
				}
			})

			rs.Print()

			return nil
		},
	}

	if e.cfg.Listener != nil {
		return e.App.Listener(e.cfg.Listener, cfg)
	}

	return e.App.Listen(e.cfg.Address, cfg)
}

func (e *Engine) Shutdown(ctx context.Context) error {
	return e.App.ShutdownWithContext(ctx)
}
