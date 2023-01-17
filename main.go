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

	// val := "%s%s__3AEhZXKHYmQHgRGDbqJZxMBKjbCFWwqDery__3AD6JBzYKQYBYhcLGd9Z1XXxnTBoSEnjPj8"
	// dataTransaction("167.99.135.70", &val, nil, nil)

	startMonitor()
}
