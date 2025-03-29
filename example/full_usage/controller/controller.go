package controller

import (
	"net/http"
	"time"

	"github.com/loveuer/uzone/pkg/api"
	api_nf "github.com/loveuer/uzone/pkg/api.nf"
)

func New() api.Engine {
	app := api_nf.New()

	app.GET("/api/available", func(c api.Context) error {
		c.UseLogger().Info("hello world")
		return c.Status(http.StatusOK).JSON(map[string]any{"ok": true, "now": time.Now()})
	})

	app.POST("/api/kv/create", kvCreate)
	app.GET("/api/kv/get", kvGet)

	return app
}
