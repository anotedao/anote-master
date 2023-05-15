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

	// val := "%s%s__3AN6QbJtvtAC7ZCfLYahTG4958sAJHGg9RQ__3AMoWvfYRvpiRPQzHmQ5rrHumGasSuuejPu"
	// dataTransaction("45.76.3.20", &val, nil, nil)

	// lease("3AN6QbJtvtAC7ZCfLYahTG4958sAJHGg9RQ")

	startMonitor()
}
