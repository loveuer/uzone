package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
)

func New() api.Engine {
	app := nf.New(nf.Config{
		DisableMessagePrint: true,
		BodyLimit:           0,
		DisableBanner:       true,
		DisableLogger:       true,
	})

	return &Engine{App: app, cfg: api.Config{}}
}
