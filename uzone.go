package uzone

import (
	"context"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/mq"
	"github.com/loveuer/uzone/pkg/uapi"
	"sync"

	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/log"
)

const Banner = `
  __  ______               
 / / / /_  / ___  ___  ___ 
/ /_/ / / /_/ _ \/ _ \/ -_)
\____/ /___/\___/_//_/\__/ 

`

type uzone struct {
	debug   bool
	ctx     context.Context
	logger  *sync.Pool
	db      *db.Client
	cache   *cache.Client
	es      *es.Client
	api     uapi.Engine
	mq      *mq.Client
	initFns struct {
		_sync  []func(interfaces.Uzone)
		_async []func(interfaces.Uzone)
	}
	taskCh []<-chan func(interfaces.Uzone) error
}

func (u *uzone) With(modules ...module) {
	for _, m := range modules {
		m(u)
	}
}

func New(configs ...Config) *uzone {
	config := Config{}

	if len(configs) > 0 {
		config = configs[0]
	}

	app := &uzone{
		logger: log.UzoneLoggerPool,
		initFns: struct {
			_sync  []func(interfaces.Uzone)
			_async []func(interfaces.Uzone)
		}{
			_sync:  make([]func(interfaces.Uzone), 0),
			_async: make([]func(interfaces.Uzone), 0),
		},
	}

	if config.Debug || property.Debug {
		log.SetLogLevel(log.LevelDebug)
		app.debug = true
	}

	return app
}
