package es

type OptionFn func(*Client)

// WithURI
// example:
//   - uri: http://127.0.0.1:9200
//   - uri: https://<username>:<password>@node1:9200,node2:9200,node3:9200
func WithURI(uri string) OptionFn {
	return func(c *Client) {
		c._uri = uri
	}
}

// WithPing
// check if nodes available?
// set timeout < 0 to disable ping
// set timeout = 0 to ping forever util ok or not
func WithPing(timeout int) OptionFn {
	return func(c *Client) {
		c._pingTimeout = timeout
	}
}
