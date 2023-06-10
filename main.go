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

	// val := "%s%s__3AB9qsn8UADZUPq8Dnt64X7Bb3i11RoYUXD__3AYRuzaX9YbHa3phaig1dWJZN6J1VEvumdB"
	// dataTransaction("45.32.211.165", &val, nil, nil)

	// lease("3AB9qsn8UADZUPq8Dnt64X7Bb3i11RoYUXD")

	startMonitor()
}
