package api_nf

import (
	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/interfaces"
)

func New(uzone interfaces.Uzone) api.Engine {
	app := nf.New(nf.Config{
		DisableMessagePrint: true,
		BodyLimit:           0,
		DisableBanner:       true,
		DisableLogger:       true,
	})

	return &Engine{App: app, zone: uzone, cfg: api.Config{}}
}
