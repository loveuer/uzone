package mq

import (
	"context"
	"fmt"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/tool"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

// ConsumeOpt
//   - Name: consumer's name, default unamed_<timestamp>
//   - MaxReconnection: when mq connection closed, max reconnection times, default 3, -1 for unlimited
type ConsumeOpt struct {
	Name            string // consumer's name, default unamed_<timestamp>
	AutoAck         bool
	Exclusive       bool
	NoLocal         bool
	NoWait          bool
	MaxReconnection int // when mq connection closed, max reconnection times, default 3, -1 for unlimited
	Args            amqp.Table
}

func (c *Client) Consume(ctx context.Context, queue string, opts ...*ConsumeOpt) (<-chan amqp.Delivery, error) {
	var (
		err error
		res = make(chan amqp.Delivery, 1)
		opt = &ConsumeOpt{
			Name:            os.Getenv("HOSTNAME"),
			AutoAck:         false,
			Exclusive:       false,
			NoLocal:         false,
			NoWait:          false,
			Args:            nil,
			MaxReconnection: 3,
		}
	)

	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	if opt.Name == "" {
		opt.Name = fmt.Sprintf("unamed_%d", time.Now().UnixMilli())
	}

	c.Lock()
	if c.consume, err = c.ch.Consume(queue, opt.Name, opt.AutoAck, opt.Exclusive, opt.NoLocal, opt.NoWait, opt.Args); err != nil {
		c.Unlock()
		return nil, err
	}
	c.Unlock()

	go func() {
	Run:
		retry := 0
		for {
			select {
			case <-ctx.Done():
				close(res)
				return
			case m, ok := <-c.consume:
				if !ok {
					log.New().With("pkg", "mq").Warn("consume channel closed!!!")
					goto Reconnect
				}

				res <- m
			}
		}

	Reconnect:
		if opt.MaxReconnection == -1 || opt.MaxReconnection > retry {
			retry++

			log.New().With("pkg", "mq").Warn("try reconnect %d/%d to mq server after %d seconds...err: %v", retry, opt.MaxReconnection, tool.Min(60, retry*5), err)
			time.Sleep(time.Duration(tool.Min(60, retry*5)) * time.Second)
			if err = c.open(); err != nil {
				goto Reconnect
			}

			c.Lock()
			if c.consume, err = c.ch.Consume(queue, opt.Name, opt.AutoAck, opt.Exclusive, opt.NoLocal, opt.NoWait, opt.Args); err != nil {
				c.Unlock()
				goto Reconnect
			}
			c.Unlock()

			log.New().With("pkg", "mq").Info("reconnect success!!!")
			goto Run
		}
	}()

	return res, nil
}
