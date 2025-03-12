package api_fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/api"
)

func New() api.Engine {
	app := fiber.New(fiber.Config{
		ServerHeader:    "",
		BodyLimit:       0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		ReadBufferSize:  0,
		WriteBufferSize: 0,
		ErrorHandler:    nil,
		AppName:         "",
		JSONEncoder:     nil,
		JSONDecoder:     nil,
		StructValidator: nil,
	})

	return &Engine{App: app}
}
