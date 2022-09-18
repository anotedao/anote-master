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
	// dataTransaction("3A9tArcHUwAuuvBTPhwKhMsmY6LFPpr7Y8P", nil, nil, nil)

	// leaseCancel("G3Uqy4ZF5UWaBNmJxzNK4ptLNM1rjekT39JnQ5FDkcC4")

	// for {
	// 	callDistributeReward("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM")
	// 	time.Sleep(time.Millisecond * 500)
	// }

	// callDistributeReward("3AEDbSc69ZZ2ok3BqgTEvzQX7MjnmhZHjBM")

	// d, _ := getData("3AAA8MEsF3R9Qaa9ZKTHpdsLvXeDNZua2xk")

	// log.Println(d)

	m.Run("127.0.0.1", Port)
}
