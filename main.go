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

	val := "%s%s__3ATYNXJA6xe3mBrXSQU4rKyDZSkT8MKUowU__3AVB7NmMDnDM16VJxNw9ZE2mt4SFz5wepsg"
	dataTransaction("206.189.59.18", &val, nil, nil)

	lease("3ATYNXJA6xe3mBrXSQU4rKyDZSkT8MKUowU")

	// startMonitor()
}
