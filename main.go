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

	// val := "%s%s__3AJzKvYgwSvs6EXYaNBZmtrvDfc9yo1qNGs__3ASLefwuE3dz9cW9bhP6ZC9N73pLqs2vPEH"
	// dataTransaction("64.227.39.44", &val, nil, nil)

	// lease("3AJzKvYgwSvs6EXYaNBZmtrvDfc9yo1qNGs")

	startMonitor()
}
