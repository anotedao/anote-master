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

	wc = initWaves()

	// val := "%s%s__3AWxRSCN41zvVKvarcL9UNZ3SmyziRxmyyd__3ARSUUKjJKrEijUKcZ5qgvNjPDvZBcNpMse"
	// dataTransaction("45.76.242.175", &val, nil, nil)

	// lease("3AWxRSCN41zvVKvarcL9UNZ3SmyziRxmyyd")

	startMonitor()
}
