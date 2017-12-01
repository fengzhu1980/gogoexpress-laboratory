package codec

import (
	"gogoexpress-laboratory/transport"
)

type Codec interface {
	Marshall(msg *transport.Message) []byte
	Unmarshall([]byte) *transport.Message
}
