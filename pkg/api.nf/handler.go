package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/samber/lo"
)

type Handler struct {
	fn nf.HandlerFunc
}

func newHandlers(engine api.Engine, must bool, handlers ...api.Handler) []nf.HandlerFunc {
	if must && len(handlers) == 0 {
		log.New().Panic("at least one handler required")
	}

	hs := lo.Map(
		handlers,
		func(item api.Handler, _ int) nf.HandlerFunc {
			return func(c *nf.Ctx) error {
				return item(NewCtx(c, engine))
			}
		},
	)

	return hs
}
