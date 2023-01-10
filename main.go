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

	// val := "%s%s__3AAqzzyFViLNvy9PrEdkMXA3nnWkzaHeE4h__3AWhzpyhXDDVBBrTjcq7dgfmWNWsQFguXPB"
	// dataTransaction("174.138.9.35", &val, nil, nil)

	startMonitor()

	// lease("3AJwp4kXyGxsjQuJAisE4jLszjywiwVcU7A")
	// leaseCancel("5fVp6qEXXJZdNww27Xxi8nK5hTQBaoWGc27o4JSmxnAD")

}
