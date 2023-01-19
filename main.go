package main

import (
	"log"

	"github.com/wavesplatform/gowaves/pkg/client"
	"gopkg.in/macaron.v1"
)

var conf *Config

var nodeAddress string

var m *macaron.Macaron

var wc *client.Client

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	wc = initWaves()

	// val := "%s%s__3AGAMvupfs74c48CjdPeAzRXvz3jsonuuWo__3ATvsMDMGEjx59WyVTT8CsXsHXAatZi2Sfq"
	// dataTransaction("193.149.180.83", &val, nil, nil)

	// lease("3AGAMvupfs74c48CjdPeAzRXvz3jsonuuWo")

	startMonitor()
}
