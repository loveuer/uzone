package controller

import (
	"github.com/loveuer/uzone/pkg/api"
	"net/http"
	"time"
)

func New() *api.App {
	app := api.New()

	app.GET("/api/available", func(c *api.Ctx) error {
		c.UseLogger().Info("hello world")
		return c.Status(http.StatusOK).JSON(map[string]any{"ok": true, "now": time.Now()})
	})

	app.POST("/api/kv/create", kvCreate)
	app.GET("/api/kv/get", kvGet)

	return app
}
