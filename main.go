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

	// val := "%s%s__3ALFr3wzWrh4wBPN459MRncgqzxhgfj4tWx__3ALx6RUd7Br64x5cwWvUrcYif2T5XSBNffQ"
	// dataTransaction("155.138.239.104", &val, nil, nil)

	// lease("3AB9qsn8UADZUPq8Dnt64X7Bb3i11RoYUXD")

	// leaseCancel("FF8tKf2dr3BQPCLw41StvRdU2dNhfeWMyJstNQdzASA1")

	startMonitor()
}
