package api_fiber

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"net"
)

type Engine struct {
	*fiber.App
	zone interfaces.Uzone
	cfg  api.Config
}

func (e *Engine) SetUZone(u interfaces.Uzone, cfg api.Config) {
	e.zone = u
	e.cfg = cfg
}

func (e *Engine) GetUZone() (interfaces.Uzone, api.Config) {
	return e.zone, e.cfg
}

func (e *Engine) Group(path string, handlers ...api.Handler) api.ApiGroup {
	hs := NewHandlers(e.zone, handlers...)
	group := e.App.Group(path, hs...)
	return NewGroup(e.zone, group)
}

func (e *Engine) GET(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Get(path, hs[0], hs[1:]...)
}

func (e *Engine) POST(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Post(path, hs[0], hs[1:]...)
}

func (e *Engine) PUT(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Put(path, hs[0], hs[1:]...)
}

func (e *Engine) DELETE(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Delete(path, hs[0], hs[1:]...)
}

func (e *Engine) HEAD(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Head(path, hs[0], hs[1:]...)
}

func (e *Engine) PATCH(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Patch(path, hs[0], hs[1:]...)
}

func (e *Engine) OPTIONS(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Options(path, hs[0], hs[1:]...)
}

func (e *Engine) Handle(method, path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Add([]string{method}, path, hs[0], hs[1:]...)
}

func (e *Engine) Use(handlers ...api.Handler) {
	for _, item := range NewHandlers(e.zone, handlers...) {
		e.App.Use(item)
	}
}

func (e *Engine) Run(address string) error {
	return e.App.Listen(address, fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}

func (e *Engine) RunListener(ln net.Listener) error {
	return e.App.Listener(ln)
}

func (e *Engine) Shutdown(ctx context.Context) error {
	return e.App.ShutdownWithContext(ctx)
}
