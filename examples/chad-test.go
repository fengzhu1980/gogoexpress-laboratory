package main

import (
  "os"
  "runtime"
  "fmt"
  "flag"
  "log"

  "github.com/nats-io/go-nats"
)

func usage() {
  fmt.Printf("Usage: chad-test [-s server] [-p port] <subject> \n")
  os.Exit(1)
}

func printMsg(m *nats.Msg, i int) {
  log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {
  var server = flag.String("s", "", "The nats server address")
  var port = flag.String("p", "4222", "The nats server port")

  flag.Usage = usage
  flag.Parse()

  args := flag.Args()
  if len(args) < 1 {
    usage()
  }

  if flag.NFlag() == 0 {
    usage()
  }

  if *server == "" {
    usage()
  }

  var url = "nats://" + *server + ":" + *port

  nc, err := nats.Connect(url)
  if err != nil {
    log.Fatalf("Cannot connect to server: %v\n", err)
  }

  subj, i := args[0], 0

  nc.Subscribe(subj, func(msg *nats.Msg) {
    i++
    printMsg(msg, i)
  })

  if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)

  runtime.Goexit()
}
