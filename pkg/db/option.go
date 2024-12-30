package db

type OptionFn func(*Client)

func WithURI(uri string) OptionFn {
	return func(c *Client) {
		c.uri = uri
	}
}

func WithAutoMigrate(models ...any) OptionFn {
	return func(c *Client) {
		c.models = models
	}
}
