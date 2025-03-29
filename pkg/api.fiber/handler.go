package api_fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/uapi"
	"github.com/samber/lo"
)

type Handler struct {
	fn fiber.Handler
}

func newHandlers(engine uapi.Engine, must bool, handlers ...uapi.Handler) []fiber.Handler {
	if must && len(handlers) == 0 {
		log.New().Panic("at least one handler required")
	}

	hs := lo.Map(
		handlers,
		func(item uapi.Handler, _ int) fiber.Handler {
			return func(c fiber.Ctx) error {
				return item(NewCtx(c, engine))
			}
		},
	)

	return hs
}
