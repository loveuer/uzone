package api_fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/api"
)

type Group struct {
	group  fiber.Router
	engine api.Engine
}

func NewGroup(engine api.Engine, group fiber.Router) *Group {
	return &Group{
		group:  group,
		engine: engine,
	}
}

func (g *Group) Group(path string, handlers ...api.Handler) api.ApiGroup {
	hs := newHandlers(g.engine, false, handlers...)
	r := g.group.Group(path, hs...)
	return &Group{group: r}
}

func (g *Group) GET(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Get(path, hs[0], hs[1:]...)
}

func (g *Group) POST(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Post(path, hs[0], hs[1:]...)
}

func (g *Group) PUT(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Put(path, hs[0], hs[1:]...)
}

func (g *Group) DELETE(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Delete(path, hs[0], hs[1:]...)
}

func (g *Group) HEAD(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Head(path, hs[0], hs[1:]...)
}

func (g *Group) PATCH(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Patch(path, hs[0], hs[1:]...)
}

func (g *Group) OPTIONS(path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Options(path, hs[0], hs[1:]...)
}

func (g *Group) Handle(method, path string, handlers ...api.Handler) {
	hs := newHandlers(g.engine, true, handlers...)
	g.group.Add([]string{method}, path, hs[0], hs[1:]...)
}

func (g *Group) Use(handlers ...api.Handler) {
	for _, item := range newHandlers(g.engine, true, handlers...) {
		g.group.Use(item)
	}
}
