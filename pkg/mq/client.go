package mq

import (
	"crypto/tls"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/url"
	"sync"
)

// Init - init mq client:
//   - @param.uri: "{scheme[amqp/amqps]}://{username}:{password}@{endpoint}/{virtual_host}"
//   - @param.certs: with amqps, certs[0]=client crt bytes, certs[0]=client key bytes

type Client struct {
	sync.Mutex
	uri     string
	tlsCfg  *tls.Config
	conn    *amqp.Connection
	ch      *amqp.Channel
	consume <-chan amqp.Delivery
	queue   *queueOption
}

func (c *Client) open() error {
	var (
		err error
	)

	c.Lock()
	defer c.Unlock()

	if c.tlsCfg != nil {
		c.conn, err = amqp.DialTLS(c.uri, c.tlsCfg)
	} else {
		c.conn, err = amqp.Dial(c.uri)
	}

	if err != nil {
		return err
	}

	if c.ch, err = c.conn.Channel(); err != nil {
		return err
	}

	if c.queue != nil && c.queue.name != "" {
		if _, err = c.ch.QueueDeclare(
			c.queue.name,
			c.queue.durable,
			c.queue.autoDelete,
			c.queue.exclusive,
			c.queue.noWait,
			c.queue.args,
		); err != nil {
			return fmt.Errorf("declare queue: %s, err: %w", c.queue.name, err)
		}
	}

	return nil
}

// New - init a new mq client
func New(opts ...OptionFn) (*Client, error) {
	var (
		err    error
		client = &Client{
			uri:     "amqp://localhost:5672",
			consume: make(<-chan amqp.Delivery),
		}
	)

	for _, fn := range opts {
		fn(client)
	}

	if _, err = url.Parse(client.uri); err != nil {
		return nil, fmt.Errorf("url parse uri err: %w", err)
	}

	if err = client.open(); err != nil {
		return nil, err
	}

	return client, nil
}
