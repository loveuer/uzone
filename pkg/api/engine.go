package api

import (
	"context"
	"crypto/tls"
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
}
