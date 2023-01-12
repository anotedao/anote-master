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

	// lease("3AFAnkpLvjVJAs6EfMC8vTT7qTVhR5a73VZ")
	// leaseCancel("5fVp6qEXXJZdNww27Xxi8nK5hTQBaoWGc27o4JSmxnAD")

	startMonitor()
}
