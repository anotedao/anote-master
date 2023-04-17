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

	// val := "%s%s__3AEN5EEf6B3bQztxuvcJmjQ1X58ncXsvb1V__3AEWfAjHgqF8vHrii7CnZokBXd3rxJKFyoY"
	// dataTransaction("206.189.59.18", &val, nil, nil)

	// lease("3AEN5EEf6B3bQztxuvcJmjQ1X58ncXsvb1V")

	// startMonitor()
}
