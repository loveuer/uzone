package api

import (
	"context"
	"crypto/tls"
	"github.com/loveuer/uzone/pkg/interfaces"
	"net"
)

type Engine interface {
	ApiGroup
	SetUZone(u interfaces.Uzone, cfg Config)
	GetUZone() (interfaces.Uzone, Config)
	SetAddress(address string)
	SetListener(ln net.Listener)
	SetTLSConfig(cfg *tls.Config)
	SetRecover(recover bool)
	Run(address string) error
	RunListener(ln net.Listener) error
	Shutdown(ctx context.Context) error
}
