package uapi

import (
	"crypto/tls"
	"net"
)

type OptionFn func(e Engine)

func SetListenAddress(address string) OptionFn {
	return func(e Engine) {
		e.SetAddress(address)
	}
}

func SetTLS(tlsConfig *tls.Config) OptionFn {
	return func(e Engine) {
		e.SetTLSConfig(tlsConfig)
	}
}

func SetListener(ln net.Listener) OptionFn {
	return func(e Engine) {
		e.SetListener(ln)
	}
}

func DisableRecover() OptionFn {
	return func(e Engine) {
		e.SetRecover(false)
	}
}
