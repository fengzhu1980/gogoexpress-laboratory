package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats"
)

var nc *nats.Conn

func main() {
	var natsServer = "nats://dev.gogox.co.nz:4222"

	nc, _ = nats.Connect(natsServer)

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%d", i)
		err := nc.Publish("gogo.test.goroutine", []byte(msg))
		if err != nil {
			log.Printf("[%s] Error!", err)
		} else {
			log.Printf("[%s] Sent!", msg)
		}
	}
	nc.Flush()
}
