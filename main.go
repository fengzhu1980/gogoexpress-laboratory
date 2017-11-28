package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nats-io/go-nats"
)

const serviceName = "gogo.msg.repeater"
const serviceVersion = "0.0.1"
const serviceID = "01"
const serverAddress = "nats://dev.gogox.co.nz:4222"

func main() {
	nc, err := nats.Connect(serverAddress)
	if err != nil {
		fmt.Printf("[Connection Error]%s\n", err)
		os.Exit(1)
	}

	reqSub := serviceName + "." + serviceID + ".request"

	fmt.Printf("[Server Address]%s\n", serverAddress)

	nc.Subscribe(reqSub, func(msg *nats.Msg) {
		if msg.Reply != "" {
			nc.Publish(msg.Reply, msg.Data)
		}
	})

	fmt.Printf("[Subject Subscribed]%s\n", reqSub)

	nc.Flush()

	runtime.Goexit()
}
