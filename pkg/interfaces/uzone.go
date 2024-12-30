package interfaces

import (
	"context"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"

	"github.com/loveuer/uzone/pkg/cache"
	"gorm.io/gorm"
)

type Uzone interface {
	Debug() bool
	UseCtx() context.Context
	UseDB(...db.SessionOpt) *gorm.DB
	UseCache() cache.Cache
	UseES() *es.Client
	UseLogger(ctxs ...context.Context) *log.UzoneLogger
	UseMQ() *mq.Client
}
