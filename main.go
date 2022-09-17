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

	// node := "3A9Rb3t91eHg1ypsmBiRth4Ld9ZytGwZe9p"
	// dataTransaction("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM", &node, nil, nil)

	// leaseCancel("3wyMQQwtcKwEjXMNc3ZSpoji3JnqqLHW8fp5FzZ2Mjjy")

	// for {
	// 	callDistributeReward("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM")
	// 	time.Sleep(time.Millisecond * 500)
	// }

	// callDistributeReward("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM")

	// d, _ := getData("3AAA8MEsF3R9Qaa9ZKTHpdsLvXeDNZua2xk")

	// log.Println(d)

	m.Run("127.0.0.1", Port)
}
