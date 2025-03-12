package api_fiber

import (
	"github.com/loveuer/uzone/pkg/api"
	"testing"
)

func TestNew(t *testing.T) {
	app := New()
	app.GET("/api/hello", func(c api.Context) error {
		return c.SendString("world")
	})

	app.Run(":9321")
}
