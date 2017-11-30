package transport

type Message struct {
	Header map[string]string
	Body   []byte
}

type MsgHandler func(msg *Message)
