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

	dataTransaction("3A9Rb3t91eHg1ypsmBiRth4Ld9ZytGwZe9p", nil, nil, nil)

	// leaseCancel("EB7NSuvEMDgff7Lyyk5G9pJeA2oiG6SHiGibiMqscn3P")

	m.Run("0.0.0.0", Port)
}
