package api

type ApiGroup interface {
	Group(path string, handlers ...Handler) ApiGroup
	GET(path string, handlers ...Handler)
	POST(path string, handlers ...Handler)
	PUT(path string, handlers ...Handler)
	DELETE(path string, handlers ...Handler)
	HEAD(path string, handlers ...Handler)
	PATCH(path string, handlers ...Handler)
	OPTIONS(path string, handlers ...Handler)
	Handle(method, path string, handlers ...Handler)
	Use(handlers ...Handler)
}
