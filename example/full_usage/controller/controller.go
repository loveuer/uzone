package controller

import (
	"net/http"
	"time"

	api_nf "github.com/loveuer/uzone/pkg/api.nf"
	"github.com/loveuer/uzone/pkg/uapi"
)

func New() uapi.Engine {
	app := api_nf.New()

	app.GET("/api/available", func(c uapi.Context) error {
		c.UseLogger().Info("hello world")
		return c.Status(http.StatusOK).JSON(map[string]any{"ok": true, "now": time.Now()})
	})

	app.POST("/api/kv/create", kvCreate)
	app.GET("/api/kv/get", kvGet)

	return app
}
