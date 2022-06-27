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

	dataTransaction("fee", nil, nil, nil)

	// leaseCancel("GJ86wnvJis81Lv33y4gBXnDrtFdTX1rWTaNPo9VjTJPH")

	m.Run("0.0.0.0", Port)
}
