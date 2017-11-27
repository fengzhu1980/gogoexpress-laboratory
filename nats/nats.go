package nats

import (
	"os"

	"github.com/micro/go-micro"

	broker "github.com/micro/go-plugins/broker/nats"
	transport "github.com/micro/go-plugins/transport/nats"
)

// NewService returns a nats service compatible with go-micro.Service
func NewService(opts ...micro.Option) micro.Service {
	// our plugins
	b := broker.NewBroker()
	t := transport.NewTransport()

	// set default envs
	os.Setenv("MICRO_BROKER", b.String())
	os.Setenv("MICRO_TRANSPORT", t.String())

	// create options with priority for our opts
	options := []micro.Option{
		micro.Broker(b),
		micro.Transport(t),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return micro.NewService(options...)
}
