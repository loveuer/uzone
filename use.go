package uzone

import (
	"context"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"

	"github.com/loveuer/uzone/pkg/cache"
	"gorm.io/gorm"
)

func (u *uzone) Debug() bool {
	return u.debug
}

func (u *uzone) UseCtx() context.Context {
	return u.ctx
}

func (u *uzone) UseDB(opts ...db.SessionOpt) *gorm.DB {
	tx := u.db.Session()

	if u.Debug() {
		tx = tx.Debug()
	}

	return tx
}

func (u *uzone) UseCache() cache.Cache {
	return u.cache
}

func (u *uzone) UseES() *es.Client {
	return u.es
}

func (u *uzone) UseMQ() *mq.Client {
	return u.mq
}

func (u *uzone) UseLogger(ctxs ...context.Context) *log.UzoneLogger {
	logger := u.logger.New().(*log.UzoneLogger)

	logger.WithContext(u.UseCtx())
	if len(ctxs) > 0 {
		logger.WithContext(ctxs[0])
	}

	return logger
}
