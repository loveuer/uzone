package api_fiber

import (
	"context"
	"testing"

	"github.com/loveuer/uzone/pkg/api"
)

func TestNew(t *testing.T) {
	app := New()
	app.GET("/api/hello", func(c api.Context) error {
		return c.SendString("world")
	})

	app.Run(context.TODO())
}
