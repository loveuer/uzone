package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
)

type Group struct {
	zone  interfaces.Uzone
	group *nf.RouterGroup
}

func NewGroup(zone interfaces.Uzone, group *nf.RouterGroup) *Group {
	return &Group{
		group: group,
		zone:  zone,
	}
}

func (g *Group) Group(path string, handlers ...api.Handler) api.ApiGroup {
	hs := newHandlers(g.zone, false, handlers...)
	r := g.group.Group(path, hs...)
	return &Group{group: r}
}

func (g *Group) GET(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Get(path, hs...)
}

func (g *Group) POST(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Post(path, hs...)
}

func (g *Group) PUT(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Put(path, hs...)
}

func (g *Group) DELETE(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Delete(path, hs...)
}

func (g *Group) HEAD(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Head(path, hs...)
}

func (g *Group) PATCH(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Patch(path, hs...)
}

func (g *Group) OPTIONS(path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Options(path, hs...)
}

func (g *Group) Handle(method, path string, handlers ...api.Handler) {
	hs := newHandlers(g.zone, true, handlers...)
	g.group.Handle(method, path, hs...)
}

func (g *Group) Use(handlers ...api.Handler) {
	for _, item := range newHandlers(g.zone, true, handlers...) {
		g.group.Use(item)
	}
}
