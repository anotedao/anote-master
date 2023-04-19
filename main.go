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

	// val := "%s%s__3AJJjHbfw9MtY58HFi8oi1ajhGbWeKAgMKb__3AA5c4VhSjmE8C9eV1fezMdVE2yzujpBVV6"
	// dataTransaction("188.25.72.100", &val, nil, nil)

	// lease("3AJJjHbfw9MtY58HFi8oi1ajhGbWeKAgMKb")

	startMonitor()
}
