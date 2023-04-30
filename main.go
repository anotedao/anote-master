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

	// val := "%s%s__3AEXLWKfxaDh88s716qTzzj8bzZmWPyw2ED__3ARSUUKjJKrEijUKcZ5qgvNjPDvZBcNpMse"
	// dataTransaction("149.28.107.208", &val, nil, nil)

	// lease("3AEXLWKfxaDh88s716qTzzj8bzZmWPyw2ED")

	startMonitor()
}
