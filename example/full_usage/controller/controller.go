package controller

import (
	"github.com/loveuer/uzone/pkg/api.nf"
	"net/http"
	"time"
)

func New() *api_nf.App {
	app := api_nf.New()

	app.GET("/api/available", func(c *api_nf.Ctx) error {
		c.UseLogger().Info("hello world")
		return c.Status(http.StatusOK).JSON(map[string]any{"ok": true, "now": time.Now()})
	})

	app.POST("/api/kv/create", kvCreate)
	app.GET("/api/kv/get", kvGet)

	return app
}
