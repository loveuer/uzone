package db

import (
	"context"
	"github.com/loveuer/uzone/pkg/tool"
	"gorm.io/gorm"
)

type Client struct {
	uri    string
	models []any
	tx     *gorm.DB
}

type SessionOpt struct {
	Ctx   context.Context
	Debug bool
}

func (c *Client) Session(opts ...SessionOpt) *gorm.DB {
	opt := SessionOpt{
		Ctx:   tool.Timeout(30),
		Debug: false,
	}

	if len(opts) > 0 {
		opt = opts[0]

		if opt.Ctx == nil {
			opt.Ctx = tool.Timeout(30)
		}
	}

	session := c.tx.Session(&gorm.Session{
		Context: opt.Ctx,
	})

	if opt.Debug {
		session = session.Debug()
	}

	return session
}
