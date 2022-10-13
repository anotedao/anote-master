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
	// dataTransaction("167.99.128.22", nil, nil, nil)
	// dataTransaction("3AQQcszuhC1QqEbRuycVw7sSZQaukexGk4K", nil, nil, nil)
	// leaseCancel("6Xs7fCCMEnRaZYcC4Yecb2RKLsSgSzybDFBVsDpnVM5V")
	// a := "3AEYtCViouVG1hgeJecCLo49Q7tKCkmNfGf"
	// b := "3ATqxdbGWSp1tKtPSkZ2pX52a1RWFjNrGZx"
	// lease(a)
	// dataTransaction("165.227.141.41", &a, nil, nil)
	// dataTransaction(a, &b, nil, nil)
	// dataTransaction("46.101.146.126", nil, nil, nil)
	// dataTransaction("3AFW2wXx8Ex7mg8yjiZCrBkWCPnR1Q29q5S", nil, nil, nil)

	m.Run("127.0.0.1", Port)
}
