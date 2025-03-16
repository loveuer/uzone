package api

import (
	"crypto/tls"
)

type Config struct {
	Address   string
	TLSConfig *tls.Config
}

type Handler func(c Context) error
