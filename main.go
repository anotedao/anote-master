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

	// val := "%s%s__3AQ2iiPTBYRPdDDg4kAZDoDwy15NoNa5JHR__3AFi9Xs6gJzfzeAheAaprXNReNq7WWvs1Qg"
	// dataTransaction("188.166.164.144", &val, nil, nil)

	// lease("3AQ2iiPTBYRPdDDg4kAZDoDwy15NoNa5JHR")

	startMonitor()
}
