package cache

type OptionFn func(*Client)

func WithURI(uri string) OptionFn {
	return func(c *Client) {
		c.uri = uri
	}
}

func WithPrefix(prefix string) OptionFn {
	return func(c *Client) {
		c.prefix = prefix
	}
}
