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

	m = initMacaron()

	wc = initWaves()

	initMonitor()

	// dataTransaction("3AWhzpyhXDDVBBrTjcq7dgfmWNWsQFguXPB", nil, nil, nil)

	// leaseCancel("VAnckWXRcvkDiAQogJnthhFEz5jPdvNZE3U7oYMcqXt")

	m.Run("127.0.0.1", Port)
}
