package uapi

import (
	"context"
	"crypto/tls"
	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"
	"gorm.io/gorm"
	"net"

	"github.com/loveuer/uzone/pkg/interfaces"
)

type Engine interface {
	ApiGroup
	SetUZone(u interfaces.Uzone)
	GetUZone() (interfaces.Uzone, Config)

	SetAddress(address string)
	SetListener(ln net.Listener)
	SetTLSConfig(cfg *tls.Config)
	SetRecover(recover bool)

	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error

	UseLogger() *log.UzoneLogger
	UseDB() *gorm.DB
	UseCache() cache.Cache
	UseES() *es.Client
	UseMQ() *mq.Client
}
