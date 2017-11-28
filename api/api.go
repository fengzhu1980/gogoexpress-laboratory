package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/go-nats"
)

const serviceName = "gogo.api.gateway"
const serviceVersion = "0.0.1"
const serviceID = "01"
const serverAddress = "nats://dev.gogox.co.nz:4222"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var nc *nats.Conn

func processMsg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("[Request URI]%s", r.RequestURI)
	log.Printf("[Request Header]Accept: %s", r.Header.Get("Accept"))
	log.Printf("[Request Body]%s", json.Get(body, "data").ToString())

}

func main() {
	router := httprouter.New()
	router.POST("/", processMsg)

	nc, err := nats.Connect(serverAddress)
	if err != nil {
		log.Printf("[Connection Error]%s\n", err)
		os.Exit(1)
	}

	resSub := serviceName + "." + serviceID + ".response"

	log.Printf("[Server Address]%s\n", serverAddress)

	nc.Subscribe(resSub, func(msg *nats.Msg) {
		// Callback
	})

	log.Printf("[Subject Subscribed]%s\n", resSub)

	nc.Flush()

	log.Fatal(http.ListenAndServe(":8080", router))
}
