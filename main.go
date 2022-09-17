package main

import (
	"log"
	"time"

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

	// dataTransaction("3AXhRYtvb51tLuRDdALhFjf384vbdtCruDc", nil, nil, nil)

	// leaseCancel("3wyMQQwtcKwEjXMNc3ZSpoji3JnqqLHW8fp5FzZ2Mjjy")

	for {
		callDistributeReward("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM")
		time.Sleep(time.Millisecond * 500)
	}

	m.Run("127.0.0.1", Port)
}
