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
	// dataTransaction("174.138.13.8", nil, nil, nil)
	// dataTransaction("3AD4PS3YWqXd3yQWi4xLBx23xSAFHRPBtpQ", nil, nil, nil)
	// leaseCancel("8SwaR4AV8CEftq8vMnYiygzzbhAS4wtWTSzC5wG9E9eh")

	m.Run("127.0.0.1", Port)
}
