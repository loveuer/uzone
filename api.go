package uzone

import (
	"github.com/loveuer/uzone/pkg/api"
	"net/http"
)

func (u *uzone) API() api.Engine { return u.api }

func (u *uzone) ApiGroup(path string, handlers ...api.Handler) api.ApiGroup {
	return u.api.Group(path, handlers...)
}

func (u *uzone) GET(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodGet, path, handlers...)
}

func (u *uzone) POST(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodPost, path, handlers...)
}

func (u *uzone) PUT(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodPut, path, handlers...)
}

func (u *uzone) DELETE(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodDelete, path, handlers...)
}

func (u *uzone) PATCH(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodPatch, path, handlers...)
}

func (u *uzone) HEAD(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodHead, path, handlers...)
}

func (u *uzone) OPTIONS(path string, handlers ...api.Handler) {
	u.HandleAPI(http.MethodOptions, path, handlers...)
}

func (u *uzone) HandleAPI(method, path string, handlers ...api.Handler) {
	u.api.Handle(method, path, handlers...)
}
