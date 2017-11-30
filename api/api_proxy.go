// +build ignore

package main

import (
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
)

const serviceName = "gogo.api.gateway"
const serviceID = "01"
const natsServerAddress = "nats://dev.gogox.co.nz:4222"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func processMsg(w http.ResponseWriter, r *http.Request, _ httprouter.params) {
	body, _ := ioutil.ReadAll(r.Body)

}

func main() {

}
