package uapi

import (
	"crypto/tls"
	"net"
)

type Config struct {
	Address   string
	Listener  net.Listener
	TLSConfig *tls.Config
	Recover   bool
}

type Handler func(c Context) error
