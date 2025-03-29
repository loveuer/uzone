package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/uapi"
	"github.com/samber/lo"
)

type Handler struct {
	fn nf.HandlerFunc
}

func newHandlers(engine uapi.Engine, must bool, handlers ...uapi.Handler) []nf.HandlerFunc {
	if must && len(handlers) == 0 {
		log.New().Panic("at least one handler required")
	}

	hs := lo.Map(
		handlers,
		func(item uapi.Handler, _ int) nf.HandlerFunc {
			return func(c *nf.Ctx) error {
				return item(NewCtx(c, engine))
			}
		},
	)

	return hs
}
