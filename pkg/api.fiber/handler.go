package api_fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/samber/lo"
)

type Handler struct {
	fn fiber.Handler
}

func newHandlers(zone interfaces.Uzone, must bool, handlers ...api.Handler) []fiber.Handler {
	if must && len(handlers) == 0 {
		log.New().Panic("at least one handler required")
	}

	hs := lo.Map(
		handlers,
		func(item api.Handler, _ int) fiber.Handler {
			return func(c fiber.Ctx) error {
				return item(NewCtx(c, zone))
			}
		},
	)

	return hs
}
