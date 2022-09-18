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

	// node := "3AKCefhcrijSwwWM671ahhMrPVrE7Je3j4s"
	// dataTransaction("188.166.93.92", nil, nil, nil)
	// dataTransaction("3AQMExacaLKDXdWP8niVeyAe7Dw2vJdmNPy", nil, nil, nil)
	// leaseCancel("9Gz3mK5WBW38gFLV4ysBC121yoTFw1DckNMU25XYCfHg")

	m.Run("127.0.0.1", Port)
}
