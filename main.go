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

	// dataTransaction("3AHtUmVN4cweXaN1U52Mvyv5yqrYLoFMsKp", nil, nil, nil)
}
