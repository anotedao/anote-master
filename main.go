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

	// leaseCancel("2LkCGSJeN6Mt2xZ1X7XcdN9jvJG2hAujuibGJNhMyL3W")

	m.Run("127.0.0.1", Port)
}
