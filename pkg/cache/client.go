package cache

import (
	"context"
	"github.com/samber/lo"
	"time"
)

type Client struct {
	c      Cache
	uri    string
	prefix string
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	return c.c.Get(ctx, c.prefix+key)
}

func (c *Client) Gets(ctx context.Context, keys ...string) ([][]byte, error) {
	return c.c.Gets(ctx, lo.Map(keys, func(item string, _ int) string {
		return c.prefix + item
	})...)
}

func (c *Client) GetScan(ctx context.Context, key string) Scanner {
	return c.c.GetScan(ctx, c.prefix+key)
}

func (c *Client) GetEx(ctx context.Context, key string, duration time.Duration) ([]byte, error) {
	return c.c.GetEx(ctx, c.prefix+key, duration)
}

func (c *Client) GetExScan(ctx context.Context, key string, duration time.Duration) Scanner {
	return c.c.GetExScan(ctx, c.prefix+key, duration)
}

func (c *Client) Set(ctx context.Context, key string, value any) error {
	return c.Set(ctx, c.prefix+key, value)
}

func (c *Client) Sets(ctx context.Context, vm map[string]any) error {
	for k, v := range vm {
		vm[c.prefix+k] = v
		delete(vm, k)
	}
	return c.c.Sets(ctx, vm)
}

func (c *Client) SetEx(ctx context.Context, key string, value any, duration time.Duration) error {
	return c.c.SetEx(ctx, c.prefix+key, value, duration)
}

func (c *Client) Del(ctx context.Context, keys ...string) error {
	return c.c.Del(ctx, lo.Map(keys, func(item string, _ int) string {
		return c.prefix + item
	})...)
}

func (c *Client) Close() error {
	return c.c.Close()
}

var _ Cache = (*Client)(nil)

func New(opts ...OptionFn) (*Client, error) {
	var (
		err    error
		client = &Client{uri: "memory://", prefix: Prefix}
	)

	for _, fn := range opts {
		fn(client)
	}

	if client.c, err = _new(client.uri); err != nil {
		return nil, err
	}

	return client, nil
}
