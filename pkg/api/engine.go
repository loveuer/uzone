package api

import (
	"context"
	"github.com/loveuer/uzone/pkg/interfaces"
	"net"
)

type Engine interface {
	ApiGroup
	SetUZone(u interfaces.Uzone, cfg Config)
	GetUZone() (interfaces.Uzone, Config)
	Run(address string) error
	RunListener(ln net.Listener) error
	Shutdown(ctx context.Context) error
}
