package uzone

import (
	"github.com/loveuer/uzone/pkg/uapi"
	"net/http"
)

func (u *uzone) API() uapi.Engine { return u.api }

func (u *uzone) ApiGroup(path string, handlers ...uapi.Handler) uapi.ApiGroup {
	return u.api.Group(path, handlers...)
}

func (u *uzone) GET(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodGet, path, handlers...)
}

func (u *uzone) POST(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodPost, path, handlers...)
}

func (u *uzone) PUT(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodPut, path, handlers...)
}

func (u *uzone) DELETE(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodDelete, path, handlers...)
}

func (u *uzone) PATCH(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodPatch, path, handlers...)
}

func (u *uzone) HEAD(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodHead, path, handlers...)
}

func (u *uzone) OPTIONS(path string, handlers ...uapi.Handler) {
	u.HandleAPI(http.MethodOptions, path, handlers...)
}

func (u *uzone) HandleAPI(method, path string, handlers ...uapi.Handler) {
	u.api.Handle(method, path, handlers...)
}
