package uzone

import (
	"net/http"

	"github.com/loveuer/uzone/pkg/api"
)

func (u *uzone) API() *api.App { return u.api.engine }

func (u *uzone) GET(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodGet, path, handlers...)
}

func (u *uzone) POST(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodPost, path, handlers...)
}

func (u *uzone) PUT(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodPut, path, handlers...)
}

func (u *uzone) DELETE(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodDelete, path, handlers...)
}

func (u *uzone) PATCH(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodPatch, path, handlers...)
}

func (u *uzone) HEAD(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodHead, path, handlers...)
}

func (u *uzone) OPTIONS(path string, handlers ...api.HandlerFunc) {
	u.HandleAPI(http.MethodOptions, path, handlers...)
}

func (u *uzone) HandleAPI(method, path string, handlers ...api.HandlerFunc) {
	u.api.engine.Handle(method, path, handlers...)
}
