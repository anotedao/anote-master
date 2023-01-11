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

	// val := "%s%s__3AVv7bW7oXx7GMLAyq9BFKgH6RpK7mHyqpL__3AQ7JCAQKWLjPnduYUuzGjHJuRWDXNW6UVx"
	// dataTransaction("146.190.55.13", &val, nil, nil)

	// lease("3AEhZXKHYmQHgRGDbqJZxMBKjbCFWwqDery")
	// leaseCancel("5fVp6qEXXJZdNww27Xxi8nK5hTQBaoWGc27o4JSmxnAD")

	startMonitor()
}
