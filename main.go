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

	// dataTransaction("3ACyYVfFcyco4RS8WLbyRSGPHPeCCiUuSqP", nil, nil, nil)

	leaseCancel("EcT4eQemgmA1NVqZk78CoU7ysELa2ekw4r6ixAT8WkWW")

	m.Run("0.0.0.0", Port)
}
