package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

func processMsg(msg *nats.Msg) {
	log.Printf("%s\n", string(msg.Data))
}

func main() {
	var natsServer = "nats://dev.gogox.co.nz:4222"

	nc, _ := nats.Connect(natsServer)

	nc.Subscribe("gogo.test.goroutine", func(msg *nats.Msg) {
		go processMsg(msg)
	})

	nc.Flush()

	runtime.Goexit()
}
