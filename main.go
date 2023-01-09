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

	// val := "%s%s__3APWcSs37en5WatNYqW3zf4UAbjriGd9nv6__3AShXVgRcRis82CwD7o9pz1Ac9vmRYMqELT"
	// dataTransaction("46.188.145.85", &val, nil, nil)

	startMonitor()
}
