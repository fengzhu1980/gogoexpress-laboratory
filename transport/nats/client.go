package nats

import (
	"gogoexpress-laboratory/transport"

	gonats "github.com/nats-io/go-nats"
)

type Connection struct {
	Addrs []string
	conn  *gonats.Conn
}

type Client struct {
	sub     string
	payload []byte
	timeout int
}

type Server struct {
	reply   string
	handler transport.MsgHandler
}
