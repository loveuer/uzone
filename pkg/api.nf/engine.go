package api_nf

import (
	"context"
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"net"
)

type Engine struct {
	*nf.App
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
	e.App.Get(path, hs...)
}

func (e *Engine) POST(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Post(path, hs...)
}

func (e *Engine) PUT(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Put(path, hs...)
}

func (e *Engine) DELETE(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Delete(path, hs...)
}

func (e *Engine) HEAD(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Head(path, hs...)
}

func (e *Engine) PATCH(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Patch(path, hs...)
}

func (e *Engine) OPTIONS(path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Options(path, hs...)
}

func (e *Engine) Handle(method, path string, handlers ...api.Handler) {
	hs := NewHandlers(e.zone, handlers...)
	e.App.Handle(method, path, hs...)
}

func (e *Engine) Use(handlers ...api.Handler) {
	for _, item := range NewHandlers(e.zone, handlers...) {
		e.App.Use(item)
	}
}

func (e *Engine) Run(address string) error {
	return e.App.Run(address)
}

func (e *Engine) RunListener(ln net.Listener) error {
	return e.App.RunListener(ln)
}

func (e *Engine) Shutdown(ctx context.Context) error {
	return e.App.Shutdown(ctx)
}
