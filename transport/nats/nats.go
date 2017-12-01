package nats

import (
	"gogoexpress-laboratory/transport"

	gonats "github.com/nats-io/go-nats"
)

type Connection struct {
	Addrs []string
	conn  *gonats.Conn
}

type Request struct {
	sub     string
	payload []byte
	timeout int
}

type Response struct {
	msg     *gonats.Msg
	handler transport.MsgHandler
}

type Client interface {
	Send(req *Request)
}

type Server interface {
	Handle(res *Response)
}

func Listen(c *Connection, sub string) {
	c.conn.Subscribe(sub, func(msg *gonats.Msg) {

	})
}
