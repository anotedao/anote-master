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

	// leaseCancel("CchGKRF61nz2Ct7b2MK299u6XEwJiG6urNzepqcLTcQC")

	m.Run("0.0.0.0", Port)
}
