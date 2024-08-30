package main

import (
	"log"
	// _ "net/http/pprof"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	startMonitor()

	// leaseCancel("")
	// leaseCancel("")
	// leaseCancel("")
	// leaseCancel("")
	// leaseCancel("")

	// dataTransaction("3AH4mqeskgxu5zmS9bm7M5jefnTRiDSrJBm", nil, nil, nil)
}
