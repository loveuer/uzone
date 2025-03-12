package api_nf

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/samber/lo"
)

type Engine struct {
	*nf.App
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

func (e *Engine) Run(ctx context.Context) error {
	var rs api.Routes = lo.Map(e.GetRoutes(), func(item nf.RouteInfo, idx int) api.Route {
		return api.Route{
			Method:      item.Method,
			Path:        item.Path,
			HandlerName: item.Handler,
		}
	})

	rs.Print()

	if e.cfg.Listener != nil {
		if e.cfg.TLSConfig != nil {
			return e.App.RunListenerTls(e.cfg.Listener, e.cfg.TLSConfig)
		}

		return e.App.RunListener(e.cfg.Listener)
	}

	if e.cfg.TLSConfig != nil {
		return e.App.RunTLS(e.cfg.Address, e.cfg.TLSConfig)
	}

	return e.App.Run(e.cfg.Address)
}

func (e *Engine) Shutdown(ctx context.Context) error {
	return e.App.Shutdown(ctx)
}
