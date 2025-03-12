package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/samber/lo"
)

type Handler struct {
	fn nf.HandlerFunc
}

func NewHandlers(zone interfaces.Uzone, handlers ...api.Handler) []nf.HandlerFunc {
	if len(handlers) == 0 {
		log.New().Panic("at least one handler required")
	}

	hs := lo.Map(
		handlers,
		func(item api.Handler, _ int) nf.HandlerFunc {
			return func(c *nf.Ctx) error {
				return item(NewCtx(c, zone))
			}
		},
	)

	return hs
}
