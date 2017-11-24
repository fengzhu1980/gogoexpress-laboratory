package main

import (
	"fmt"

	"gogoexpress-laboratory/nats"

	micro "github.com/micro/go-micro"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/codec/jsonrpc2"
)

func main() {
	client := client.NewClient(
		client.Codec("application/json", jsonrpc2.NewCodec),
		client.ContentType("application/json"),
	)

	server := server.NewServer(
		server.Codec("application/json", jsonrpc2.NewCodec),
	)

	service := nats.NewService(
		micro.Name("hello"),
		micro.Version("latest"),
		micro.Client(client),
		micro.Server(server),
	)

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
