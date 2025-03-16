package api

import (
	"crypto/tls"
	"net"
)

type Option func(e Engine)

func SetListenAddress(address string) Option {
	return func(e Engine) {
		e.SetAddress(address)
	}
}

func SetTLS(tlsConfig *tls.Config) Option {
	return func(e Engine) {
		e.SetTLSConfig(tlsConfig)
	}
}

func SetListener(ln net.Listener) Option {
	return func(e Engine) {
		e.SetListener(ln)
	}
}

func DisableRecover() Option {
	return func(e Engine) {
		e.SetRecover(false)
	}
}
