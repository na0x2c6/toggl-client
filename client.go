package togglclient

import (
	"github.com/na0x2c6/toggl-client/client"
)

func New(email, password string, options ...client.Option) *client.Client {
	return client.NewClient(email, password, options...)
}

const (
	Version = "0.1.0"
)
